package fileSystem

import (
	"io/ioutil"
	"net/http"
)
//GetFileContentFromInternet ret
func GetFileContentFromInternet(url string) string {
	res, err := http.Get(url);
	check(err);
	responseContent, readErr := ioutil.ReadAll(res.Body)
	check(readErr)
	closingErr := res.Body.Close()
	check(closingErr)
	return string(responseContent)
}

func createATempFileFromTextContent(data string) string {
	fileData := data
	return fileData
}