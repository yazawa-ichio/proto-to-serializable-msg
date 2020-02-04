package proto

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Generator interface {
	Generate(files []string) ([]*GenerateFile, error)
	GenerateAndOutput(files []string, outputRoot string) error
}

type GenerateFile struct {
	Name    string
	Content string
}

type GenerateTask struct {
	Files     []string
	Output    string
	Generator Generator
}

func (g *GenerateTask) Run() error {
	return output(g.Generator, g.Files, g.Output)
}

func GenerateAndOutput(lang string, files []string, outputRoot string) error {

	if f, err := os.Stat(outputRoot); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf(fmt.Sprintf("not found dir %s", outputRoot))
	}
	g, err := GetGenerator(lang)
	if err != nil {
		return err
	}
	out, err := g.Generate(files)
	if err != nil {
		return err
	}
	for _, o := range out {
		outputPath := filepath.Join(outputRoot, o.Name)
		if f, err := os.Stat(filepath.Dir(outputPath)); os.IsNotExist(err) || !f.IsDir() {
			if err := os.MkdirAll(filepath.Dir(outputPath), 0777); err != nil {
				return err
			}
		}
		ioutil.WriteFile(outputPath, []byte(o.Content), 0777)
	}
	return nil
}

func output(g Generator, files []string, outputRoot string) error {
	if f, err := os.Stat(outputRoot); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf(fmt.Sprintf("not found dir %s", outputRoot))
	}
	out, err := g.Generate(files)
	if err != nil {
		return err
	}
	for _, o := range out {
		outputPath := filepath.Join(outputRoot, o.Name)
		if f, err := os.Stat(filepath.Dir(outputPath)); os.IsNotExist(err) || !f.IsDir() {
			if err := os.MkdirAll(filepath.Dir(outputPath), 0777); err != nil {
				return err
			}
		}
		ioutil.WriteFile(outputPath, []byte(o.Content), 0777)
	}
	return nil
}

func GetGenerator(prm string) (Generator, error) {
	switch prm {
	case "js":
		return NewJSGenerator(), nil
	case "javascript":
		return NewJSGenerator(), nil
	case "cs":
		return NewCSGenerator(), nil
	case "csharp":
		return NewCSGenerator(), nil
	case "ts":
		return NewTSGenerator(), nil
	case "typescript":
		return NewTSGenerator(), nil
	case "go":
		return NewGoGenerator(), nil
	case "golang":
		return NewGoGenerator(), nil
	}
	return nil, errors.New("not found generator")
}
