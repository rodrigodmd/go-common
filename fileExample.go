package main

import (
	"github.com/rodrigodmd/go-common/base/file"
)

type Example struct {
	Username string `yaml:"username"        json:"username"`
	Password string `yaml:"password"        json:"password"`
}

type MyFile struct {
	Model  interface{}
	MyPath string
}

func (mf MyFile) FillPath() *file.File {
	return &file.File{
		AbsPath: file.FullPath(mf.MyPath),
		Model:   mf.Model,
	}
}

func main() {
	trySaveFile()
	tryLoadAndSaveFile()

}

func trySaveFile() {
	config := MyFile{
		MyPath: "./data/config.yml",
		Model: Example{
			"Username",
			"Paaasss",
		},
	}
	file.SaveFile(config)
}

func tryLoadAndSaveFile() {
	config := MyFile{
		MyPath: "./data/configFrom.yml",
	}

	config.Model = file.ReadFile(config)
	config.MyPath = "./data/configTo.yml"
	file.SaveFile(config)

}
