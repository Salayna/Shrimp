package fileSystem

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		fmt.Println(e)
		os.Exit(1)
	}
}

//CreateDirectory Will create a directory using the project name
func CreateDirectory(name string) {
	err := os.MkdirAll(name, 0777)
	check(err)
}

//CheckIfFolderExists
func CheckIfFolderExists(folder string) (string, error) {
	_, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println(err);
		return "", err
	}
	fmt.Println("Folder opened")
	return "Folder exist", nil
}

//CreateFile will create files
func CreateFile(file string) {
	d := []byte("")
	err := ioutil.WriteFile(file, d, 0666)
	check(err)
}

//ExecuteCommand will execute commands according to the corresponding yaml file
func ExecuteCommand(command string, arguments []string) {
	out, err := exec.Command(command, arguments...).Output()
	check(err)
	output := string(out[:])
	fmt.Println(output)
}

// OpenFile handles files
func OpenFile(name string) string {
	data, err := ioutil.ReadFile(name)
	check(err)
	return string(data)
}

//GetConfigs Parse and return all the configfiles with the corresponding extension
func GetConfigs(directory string, extension string) []string {
	homeDir, _ := os.UserHomeDir();
	var parsedFiles []string
	files, err := ioutil.ReadDir( homeDir + "/.shrimp")
	check(err)
	for _, f := range files {
		if filepath.Ext(f.Name()) == extension {
			res := strings.ReplaceAll(f.Name(), extension, "")
			parsedFiles = append(parsedFiles, res)
		}
	}
	fmt.Println(parsedFiles)
	return parsedFiles
}
