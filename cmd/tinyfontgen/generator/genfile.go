package generator

import (
	"log"
	"os"
	"os/exec"
	"text/template"

	"tinygo.org/x/tinyfont"
)

const fontTemplateFile = "./templates/font.txt"

type FontTemplate struct {
	FontName string
	Font     tinyfont.Font
}

func TinyFontFile(filename string, fontName string, font tinyfont.Font) (*os.File, error) {
	tmpl, err := template.ParseFiles(fontTemplateFile)
	if err != nil {
		log.Printf("failed to parse template: %s: %s\n", fontTemplateFile, err.Error())
		return nil, err
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Printf("failed to create file with name %s: %s\n", filename, err.Error())
		return nil, err
	}

	err = tmpl.Execute(file, FontTemplate{
		FontName: fontName,
		Font:     font,
	})
	if err != nil {
		log.Printf("failed to execute template: %s: %s\n", fontTemplateFile, err.Error())
		return nil, err
	}

	err = executeGofmt(file)
	if err != nil {
		log.Printf("failed run command gofmt on file: %s: %s\n", filename, err.Error())
		return nil, err
	}

	return file, nil
}

func executeGofmt(file *os.File) error {
	cmd := exec.Command(`gofmt`, file.Name())

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
