package file

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	. "github.com/rodrigodmd/go-common/base"
	yaml "gopkg.in/yaml.v2"
)

var log Logger

type FileType int

const (
	YAML FileType = iota + 0
	JSON
	OTHER
)

type FilePather interface {
	FillPath() *File
}

type File struct {
	ErrorHandler
	AbsPath  string
	FileType FileType
	Model    interface{}
}

// Save the model in the configured file
// Part of the Saving process can be overwiten:
// - validate    (generates path)
// - Marshal     (file marshal process)
// - HandleError (Handle error )
//
func SaveFile(fp FilePather) {
	f := fp.FillPath()
	log.Info("Writing file: " + f.AbsPath)

	byteFile, err := Marshal(f)
	f.HandleError()
	log.DebugObj(byteFile)

	err = ioutil.WriteFile(f.AbsPath, byteFile, 0644)
	f.HandleError(err)
}

// Save the model in the configured file
// Part of the Saving process can be overwiten:
// - validate    (generates path)
// - Unmarshal     (file Unmarshal process)
// - HandleError (Handle error )
//
func ReadFile(fp FilePather) interface{} {
	f := fp.FillPath()
	log.Info("Reading file: " + f.AbsPath)

	strFile, err := ioutil.ReadFile(f.AbsPath)
	f.HandleError(err)

	err = Unmarshal(strFile, f)
	log.DebugObj(f.Model)

	return f.Model
}

// Method to Marshal the model
// Can be marshalled to:
// - YAML
// - JSON
func Marshal(f *File) ([]byte, error) {
	var byteFile []byte
	var err error

	if f.FileType == JSON {
		byteFile, err = json.Marshal(f.Model)
	} else if f.FileType == YAML {
		byteFile, err = yaml.Marshal(f.Model)
	}

	f.HandleError()
	return byteFile, err
}

// Method to Unmarshal the file to model
// Can be unmarshalled to:
// - YAML
// - JSON
func Unmarshal(strFile []byte, f *File) error {
	var err error

	if f.FileType == JSON {
		err = json.Unmarshal(strFile, &f.Model)
	} else if f.FileType == YAML {
		err = yaml.Unmarshal(strFile, &f.Model)
	}

	f.HandleError(err)
	return err
}

// This method is executied before saving or writing
// Can be overwrited to modify the file path logic
func FullPath(path string) string {
	str, _ := filepath.Abs(path)
	return str
	//TODO: clean up file type seting type
	// if cf.BasePath == "" {
	//TODO: set a base path using an environment variable
	// }
}
