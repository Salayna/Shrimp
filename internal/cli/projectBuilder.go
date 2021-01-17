package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/Salayna/create-project/internal/filesystemhelper"
	jsonparser "github.com/Salayna/create-project/internal/parser"
)

//BuildArborescence will build the final Arborescence
func BuildArborescence() {
	var data = filesystemhelper.OpenFile("./configs/" + language + ".json")
	filesystemhelper.CreateDirectory(name)
	os.Chdir(name)
	newDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(newDir + " directory created")

	var config jsonparser.Config
	jsonparser.ParseConfig([]byte(data), &config)

	filesystemhelper.BuildProjectDirectories(config.Directories)
	filesystemhelper.CreateProjectFiles(config.Files)
	filesystemhelper.LaunchCommands(config.Commands)
}
