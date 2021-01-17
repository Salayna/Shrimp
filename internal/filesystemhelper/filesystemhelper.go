package filesystemhelper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

//CreateDirectory Will create a directory using the project name
func CreateDirectory(name string) {
	err := os.Mkdir(name, 0600)
	check(err)
}

//CreateFile will create files
func CreateFile(file string) {
	d := []byte("")
	err := ioutil.WriteFile(file, d, 0600)
	check(err)
}

//ExecuteCommand will execute commands according to the corresponding yaml file
func ExecuteCommand(command string, arguments []string) {
	binary, lookErr := exec.LookPath(command)
	check(lookErr)
	env := os.Environ()
	args := []string{command}
	for i := 0; i < len(arguments); i++ {
		args = append(args, arguments[i])
	}
	fmt.Println(args)
	execErr := syscall.Exec(binary, args, env)
	check(execErr)
}

// OpenFile handles files
func OpenFile(name string) string {
	data, err := ioutil.ReadFile(name)
	check(err)
	return string(data)
}

//GetConfigs Parse and return all the configfiles with the corresponding extension
func GetConfigs(directory string, extension string) []string {
	var parsedFiles []string
	files, err := ioutil.ReadDir("./configs")
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
