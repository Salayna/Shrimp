# Shrimp CLI

_I'm a newbie in Golang, I'm still figuring out things so you'll probably have a lot to say about this, so feel free publish issues_

## Install

### Just building a binary file

To install it locally you can just run the following command `go build -o ./bin/shrimp` after cloning the project sources

### global install

You can Install it using the command ``npm install -g @salayna/shrimp_cli``
## Running the app

With the local installation you can run the cli with this command `./{path/to/your/binary/file}`
[comment]: <> And with the global installation you can just run `shrimp`.

### The arguments
_this part is deprecated, while the arguments are still there I'm not satisfied with them so by the release of V1 these will drastically change_
- **-n || --name**: is the name/directory of the project
- **-l || --language**: is the language of the project to be chosen from a specific set of configuration files, stored in the configs folder
- **-f || --file**: **Not Implement yet**is if you wanna created a project from a configuartion file (like for a one time project)
- **-v || --verbose** **Not Implement yet**
- **-v || --version** **Not Implement yet**

### Adding configuations

The configurations files will be stored in the configs folder at ``$HOME/.shrimp``. All of them must be written in JSON
here is an example of configuration files

```json
{
  "directories": {
    "src": {
      "baseDir": "internal",
      "subDir": [
        "cli"
      ]
    },
    "scripts": {
      "baseDir": "pkg"
    },
    "deployments": {
      "baseDir": "deployments"
    },
    "test": {
      "baseDir": "test"
    },
    "docs": {
      "baseDir": "docs"
    }
  },
  "commands": {
    "git": {
      "base": "git",
      "arguments": ["init"]
    },
    "golang init": {
      "base": "go",
      "arguments": ["mod", "init"]
    }
  },
  "files": {
    "LICENSE": {
      "name": "LICENSE",
      "directory": "./"
    },
    "Main": {
      "name": "main.go",
      "directory": "./"
    },
    "ReadMe": {
      "name": "README.md",
      "directory": "./"
    }
  }
}
```
