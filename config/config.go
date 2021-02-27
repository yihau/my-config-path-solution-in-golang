package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func LoadConfig(path string) error {
	_, file, _, _ := runtime.Caller(1)
	content, err := ioutil.ReadFile(filepath.Join(filepath.Dir(file), path))
	if err != nil {
		return err
	}
	fmt.Println(string(content)) // for test here
	return nil
}
