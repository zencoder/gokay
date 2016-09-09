// gokay is a tool that parses the structs defined in a go source file and generates a Validate() function for each of
// those structs when necessary. See gkgen documentation for information on how Validate functions are built.
//
// Usage:
//	gokay {file_name} ({generator package} {generator contructor})
// Example:
//  gokay file.go gkcustom NewCustomGKGenerator
// It relies on goimports tool to resolve import path of custom generator
package main
