package filesystem

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
