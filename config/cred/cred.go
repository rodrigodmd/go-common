package cred

import (
	"fismobile.com/mb-common/base/file"
	. "fismobile.com/mb-common/base/fileImpl"
)

var config ConfigFile

func load() {
	model := ConfigCredsModel{}

	config = ConfigFile{
		FileName: "creds.yml",
		Model:    model,
	}
	config.Model = file.ReadFile(config)

	// file.SaveFile(config)
}

// getCred

// saveCres
