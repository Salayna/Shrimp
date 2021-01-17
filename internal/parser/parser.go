package jsonparser

import "encoding/json"

//ParseConfighaha
func ParseConfig(data []byte, c *Config) {
	json.Unmarshal(data, &c)
}
