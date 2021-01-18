# Create Project CLI

_I'm a newbie in Golang, I'm still figuring out things so you'll probably have a lot to say about this, so feel free publish issues_

## Install

First you'll need to download or clone the sources of the project

### local install

To install it locally you can just run the following command `go build -o ./bin/create-project`

### global install

To install it gobally just run a `go install of the project`

## Running the app

With the local installation you can run the cli with this command `./{path/to/your/binary/file}`
And with the global installation you can just run `create-project` with it's arguments

### The arguments

- **-n || --name**: is the name/directory of the project
- **-l || --language**: is the language of the project to be chosen from a specific set of configuration files, stored in the configs folder
- **-f || --file**: **Not Implement yet**is if you wanna created a project from a configuartion file (like for a one time project)
- **-v || --verbose** **Not Implement yet**
- **-v || --version** **Not Implement yet**

### Adding configuations

The configurations files will be stored in the configs folder in the project root. All of them must be written in JSON
here is an example of configuration files

```json
{
  "directories": {
    "src": {
      "baseDir": "base directory",
      "subDir": ["sub", "directories", "optionals"]
    },
    "scripts": {
      "baseDir": "scripts"
    }
  },
  "commands": {
    "command1": {
      "base": "base command",
      "arguments": ["commands", "arguments"]
    }
  },
  "files": {
    "file1": {
      "name": "file name",
      "directory": "file directory"
    }
  }
}
```
