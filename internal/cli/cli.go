package cli

import (
	"github.com/Salayna/create-project-cli/internal/filesystemhelper"
	jsonparser "github.com/Salayna/create-project-cli/internal/parser"
)

var data = `
{
  "directories": {
    "src": {
      "baseDir": "src",
      "subDir": ["config", "controllers", "services", "models"]
    },
    "scripts": {
      "baseDir": "scrpits"
    }
  },
  "commands": {
    "npm-init": {
      "base": "",
      "arguments": ["init"]
    },
    "tsc-init": {
      "base": "tsc",
      "arguments": ["init"]
    }
  },
  "files": {
    "eslintr": {
      "name": ".eslintrc",
      "directory": "./"
		},
		"index": {
      "name": "index.js",
      "directory": "./src"
    }
  }
}

`

//Cli is the cli function which will handle all the cli calls
func Cli() {
	Init()
	CheckArguments()
	var conf jsonparser.Config
	jsonparser.ParseConfig([]byte(data), &conf)
	filesystemhelper.BuildProjectDirectories(conf.Directories)
	filesystemhelper.CreateProjectFiles(conf.Files)
}
