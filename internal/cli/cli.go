package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"rsc.io/getopt"
)

var name string
var file string
var language string
var v bool

//Init initialize the cli flags
func Init() {
	flag.StringVar(&name, "name", "", "Set the name of the project")
	flag.StringVar(&language, "language", "", "get the language")
	flag.BoolVar(&v, "verbose", false, "show verbose output if there are multiple arguments")
	flag.BoolVar(&v, "version", false, "Show current cli version when only -v is called")
	flag.StringVar(&file, "file", "", "create a project from a yaml file")
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

//Version show the version of the current cli
func Version() {
	if v == true && len(os.Args) == 2 {
		fmt.Println("V 1.0.0")
	}
}

func promptName() string {
	namePrompt := promptui.Prompt{
		Label: "Project Name(or directory)",
	}
	promptedName, err := namePrompt.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return promptedName
}

func promptLanguage() string {
	langPrompt := promptui.Select{
		Label: "Language",
		Items: []string{"Javascript", "Typescript", "Golang", "C++"},
	}

	_, promptLanguage, langErr := langPrompt.Run()
	if langErr != nil {
		log.Fatal(langErr)
		os.Exit(1)
	}
	return promptLanguage
}

//Interactive handles Interactive mode
func Interactive() {

	if name == "" {
		name = promptName()
	}
	if language == "" {
		language = promptLanguage()
	}
	fmt.Println("Project name is", name)
	fmt.Println("Type", language)
}

//Cli is the cli function which will handle all the cli calls
func Cli() {
	Init()
	fmt.Println(file)
	if (((name) == "") || (language == "")) && (v == false) && (file == "") {
		Interactive()
	} else if file != "" {
		if (name != "") || (language != "") {
			log.Fatal("You can't pass other arguments with -f or --file")
		}
	}
	Version()
}
