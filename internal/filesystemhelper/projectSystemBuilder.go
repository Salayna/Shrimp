package filesystemhelper

import (
	"fmt"

	jsonparser "github.com/Salayna/create-project/internal/parser"
)

//BuildProjectDirectories will build the directories, it takes a map as an argument
func BuildProjectDirectories(directories map[string]jsonparser.Directory) {
	for _, element := range directories {
		CreateDirectory("./" + element.BaseDir)
		for i := 0; i < len(element.Sub); i++ {
			fmt.Println("./" + element.BaseDir + "/" + element.Sub[i])
			CreateDirectory("./" + element.BaseDir + "/" + element.Sub[i])
		}
	}
}

//CreateProjectFiles will create the project files specified in the json
func CreateProjectFiles(files map[string]jsonparser.File) {
	for _, element := range files {
		CreateFile(element.Directory + "/" + element.Name)
	}
}

//LaunchCommands will launch the specified commands
func LaunchCommands(commands map[string]jsonparser.Command) {
	for _, element := range commands {
		ExecuteCommand(element.Base, element.Arguments)
	}
}
