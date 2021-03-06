package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Salayna/shrimp/internal/file_system"
	jsonparser "github.com/Salayna/shrimp/internal/parser"
)

//CheckConfigurationFolder Check is the configuration folder is set before using the CLI
func CheckConfigurationFolder() {
	homeDir, _ := os.UserHomeDir()
	_, err := fileSystem.CheckIfFolderExists(homeDir + "/.shrimp");
	if err != nil {
		fmt.Println("Creating config folder.....");
		err := os.Mkdir(homeDir + "/.shrimp", 0777)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}
//BuildArborescence will build the final Arborescence
func BuildArborescence() {
	homeDir, _ := os.UserHomeDir();
	var data string
	if remote {
		data = fileSystem.GetFileContentFromInternet(config)
	} else {
		data = fileSystem.OpenFile(homeDir+ "/.shrimp/" + language + ".json")
	}
	fileSystem.CreateDirectory(name)
	os.Chdir(name)
	newDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(newDir + " directory created....")

	var config jsonparser.Config
	jsonparser.ParseConfig([]byte(data), &config)

	fileSystem.BuildProjectDirectories(config.Directories)
	fileSystem.CreateProjectFiles(config.Files)
	fileSystem.LaunchCommands(newDir ,config.Commands)
}
