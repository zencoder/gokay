package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pborman/uuid"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Println("Usage : gokay <file_name> [{generator-package} {generator-contructor}] \n example: gokay file.go gkcustom NewCustomGKGenerator")
		return
	}
	log.Println("gokay started. file:", args[0])

	genPackage := "gkgen"
	genConstructor := "NewValidator"
	if len(args) >= 3 {
		genPackage = args[1]
		genConstructor = args[2]
	}

	fileName := args[0]

	fileName, _ = filepath.Abs(fileName)
	fileDir := filepath.Dir(fileName)

	tempName := uuid.NewRandom().String()

	tempDir := fmt.Sprintf("%s/%s", fileDir, tempName)
	err := os.Mkdir(tempDir, os.ModePerm)
	tempFile := fmt.Sprintf("%s/%s.go", tempDir, tempName)

	outFilePath := fmt.Sprintf("%s_validators.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))
	tempOut, err := os.Create(tempFile)
	defer tempOut.Close()

	outWriter := io.MultiWriter(tempOut, os.Stdout)

	fmt.Println(tempDir)
	//	fmt.Println(err.Error())

	fset := token.NewFileSet() // positions are relative to fset
	// Parse the file given in arguments
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error while parsing %v: %v\n", fileName, err)
	}
	_, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("IO Error while reading %v: %v\n", fileName, err)
	}

	fmt.Fprintf(outWriter, `package main
	func main() {
		out, err := os.Create("%s")
		defer out.Close()
		fmt.Fprint(out, "// Code in this file generated by gokay: github.com/zencoder/gokay\n")
		fmt.Fprint(out, "package %s\n")
		if err != nil {
			panic(err.Error())
		}
		v := %s.%s()
	`, outFilePath, f.Name.String(), genPackage, genConstructor)

	sortedObjectKeys := make([]string, 0, len(f.Scope.Objects))
	for k := range f.Scope.Objects {
		sortedObjectKeys = append(sortedObjectKeys, k)
	}
	sort.Strings(sortedObjectKeys)

	for _, k := range sortedObjectKeys {
		d := f.Scope.Objects[k]
		ts, ok := d.Decl.(*ast.TypeSpec)
		if !ok {
			continue
		}
		switch ts.Type.(type) {
		case *ast.StructType:
			fmt.Fprintf(outWriter, "if err := v.Generate(out, %s.%s{}); err != nil {\npanic(err.Error())\n}\n", f.Name.String(), ts.Name.String())
		}
	}

	fmt.Fprintf(outWriter, "}\n")

	tempOut.Close()

	tmpimportsCmd := exec.Command("goimports", "-w", tempOut.Name())
	tmpimportsCmd.Stdout = os.Stdout
	tmpimportsCmd.Stderr = os.Stderr
	err = tmpimportsCmd.Run()
	if err != nil {
		log.Fatalf("Failed running goimports on intermediate executable code: %v\n", err.Error())
	}

	generateCmd := exec.Command("go", "run", tempFile)
	generateCmd.Stderr = os.Stderr
	generateCmd.Stdout = os.Stdout
	err = generateCmd.Run()
	if err != nil {
		log.Fatalf("Failed executing intermediate executable to generate gokay validators failed: %v\n", err.Error())
	}

	importsCmd := exec.Command("goimports", "-w", outFilePath)
	importsCmd.Stdout = os.Stdout
	importsCmd.Stderr = os.Stderr
	err = importsCmd.Run()
	if err != nil {
		log.Fatalf("Failed running imports on gokay validators: %v\n", err.Error())
	}

	err = os.RemoveAll(tempDir)
	if err != nil {
		log.Printf("Warning: Deleting intermediate temp files %v failed (although gokay run appears to have succeeded): %v\n", tempDir, err.Error())
	}

	log.Println("gokay finished. file:", args[0])
}
