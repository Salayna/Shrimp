package cli

import (
	"flag"
	"fmt"

	"rsc.io/getopt"
)

var name string
var file string
var language string
var verbose bool
var version bool

//Init initialize the cli flags
func Init() {
	flag.StringVar(&name, "name", "", "Set the name of the project")
	flag.StringVar(&language, "language", "", "get the language")
	flag.BoolVar(&verbose, "verbose", false, "show verbose output if there are multiple arguments")
	flag.BoolVar(&version, "version", false, "Show current cli version")
	flag.StringVar(&file, "file", "./config.yaml", "create a project from a yaml file")
	//flag.StringVar(&help, "help", flag.PrintDefaults(), "Display help")
	getopt.Aliases(
		"n", "name",
		"f", "file",
		"l", "language",
		"v", "verbose",
		"V", "version",
		//"h", "help",
	)

	getopt.Parse()
}

//Cli is the cli function which will handle all the cli calls
func Cli() {
	Init()
	fmt.Println("Name:", name)
}
