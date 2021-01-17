package jsonparser

type Config struct {
	Directories map[string]Directory `json:"directories"`
	Commands    map[string]Command   `json:"commands"`
	Files       map[string]File      `json:"files"`
}
type Directory struct {
	BaseDir string   `json:"baseDir"`
	Sub     []string `json:"subDir"`
}
type Command struct {
	Base      string   `json:"base"`
	Arguments []string `json:"arguments"`
}
type File struct {
	Name      string `json:"name"`
	Directory string `json:"directory"`
}
