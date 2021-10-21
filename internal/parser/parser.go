package jsonparser

import (
	"encoding/json"
	"log"
	"os"
)

//ParseConfighaha
func ParseConfig(data []byte, c *Config) {
	err := json.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
