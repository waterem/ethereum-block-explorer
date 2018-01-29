package file

import (
	"io/ioutil"
	"os"
)

func init() {
}

func ReadLastBlock(filePath string) string {
	// create file if not exist
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		d := []byte("1")
		err := ioutil.WriteFile(filePath, d, 0644)
		if err != nil {
			panic(err)
		}
	}
	// read the file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}
