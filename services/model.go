package services

import (
	"github.com/culturadevops/jgt/jio"
)

func createFolderIsNotExist(folderName string) {
	if !jio.IsFileExist(folderName) {
		jio.CreateDirAll(folderName)
	}
}

func Model(filename string) {
	createFolderIsNotExist("models")
	MapForReplace := make(map[string]string)
	MapForReplace["%MODELNAME%"] = filename
	newName := "models/" + filename + ".go"
	templateName := "C:/Users/Soporte/.config/anferdev/model.stub"
	jio.NewFileforTemplate(newName, templateName, MapForReplace)
}
