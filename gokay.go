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
	"golang.org/x/tools/imports"
)

// usage is a string used to provide a user with the application usage
const usage = `usage: gokay <file> [generator-package generator-contructor]
	generator-package        custom package
	generator-contructor     custom generator

examples:
	gokay file.go
	gokay file.go gkcustom NewCustomGKGenerator
`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, usage)
		return
	}

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

	tempDir := fmt.Sprintf("%s/tmp/%s", fileDir, tempName)
	tempFile := fmt.Sprintf("%s/%s.go", tempDir, tempName)

	// os.Mkdir errors when the dir already exists, which is fine.
	// Real errors will still be detected when we try use the actual tmpfile.
	os.Mkdir(fileDir+"/tmp", os.ModePerm)
	os.Mkdir(tempDir, os.ModePerm)

	outFilePath := fmt.Sprintf("%s_validators.go", strings.TrimSuffix(fileName, filepath.Ext(fileName)))
	tempOut, err := os.Create(tempFile)
	if err != nil {
		log.Fatalf("Error while opening %v: %v\n", tempFile, err)
	}
	defer tempOut.Close()

	outWriter := io.MultiWriter(tempOut)

	fset := token.NewFileSet() // positions are relative to fset

	// Parse the file given in arguments
	f, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Error while parsing %v: %v\n", fileName, err)
	}
	if _, err = ioutil.ReadFile(fileName); err != nil {
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

	sortedObjectKeys := make([]string, len(f.Scope.Objects))
	for k := range f.Scope.Objects {
		sortedObjectKeys = append(sortedObjectKeys, k)
	}
	sort.Strings(sortedObjectKeys)

	for _, k := range sortedObjectKeys {
		if k == "" {
			continue
		}
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

	if err := formatFile(tempOut.Name(), nil); err != nil {
		log.Fatalf("Failed formatting intermediate executable code: %s", err.Error())
	}

	generateCmd := exec.Command("go", "run", tempFile)
	generateCmd.Stderr = os.Stderr
	generateCmd.Stdout = os.Stdout
	if err := generateCmd.Run(); err != nil {
		log.Fatalf("Failed executing intermediate executable to generate gokay validators failed: %v\n", err.Error())
	}

	if err := formatFile(outFilePath, nil); err != nil {
		log.Fatalf("Failed running imports on gokay validators: %v\n", err.Error())
	}

	// remove the temp directory
	if err := os.RemoveAll(tempDir); err != nil {
		log.Printf("Warning: Deleting intermediate temp files %v failed (although gokay run appears to have succeeded): %v\n", tempDir, err.Error())
	}

	log.Println("gokay finished file:", args[0])
}

func formatFile(filename string, src []byte) error {
	mode := os.FileMode(0644)
	formatted, err := imports.Process(filename, src, nil)
	if err != nil {
		return fmt.Errorf("failed formatting: %w", err)
	}

	err = ioutil.WriteFile(filename, formatted, mode)
	if err != nil {
		return fmt.Errorf("failed writing file: %w", err)
	}

	return nil
}
