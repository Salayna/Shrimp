package cli

//Cli is the cli function which will handle all the cli calls
func Cli() {
	CheckConfigurationFolder()
	Init()
	CheckArguments()
	BuildArborescence()
}