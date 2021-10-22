package cli

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Salayna/shrimp/internal/filesystemhelper"
	"github.com/manifoldco/promptui"
	"rsc.io/getopt"
)

var version = "development"
var name string
var config string
var remote bool
var language string
var v bool

//Init initialize the cli flags
func Init() {
	flag.StringVar(&name, "name", "", "Set the name of the project")
	flag.StringVar(&language, "language", "", "get the language")
	flag.StringVar(&config, "config", "", "create a project from a remote config file")
	//flag.StringVar(&help, "help", flag.PrintDefaults(), "Display help")
	getopt.Aliases(
		"n", "name",
		"c", "config",
		"l", "language",
		//"h", "help",
	)

	getopt.Parse()
}

//CheckArguments will check the arguments of the CLI
func CheckArguments() {
	if config != "" {
		remote = true
	}
	if (((name) == "") || (language == "")) && (config == "") {
		Interactive()
	} else if config != "" {
		if language != "" {
			log.Fatal("You can only pass a name argument(-n or --name) with -c or --config")
		}
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
	homeDir, _ := os.UserHomeDir()
	langPrompt := promptui.Select{
		Label: "Language",
		Items: filesystemhelper.GetConfigs(homeDir + "/.shrimp", ".json"),
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
	fmt.Println("Project created at ", name)
	fmt.Println("Type", language)
}
