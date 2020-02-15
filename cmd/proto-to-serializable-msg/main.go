package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	proto "github.com/yazawa-ichio/proto-to-serializable-msg"
)

func main() {
	var config string
	var lang string
	var input string
	var output string
	var dryrun bool

	flag.StringVar(&config, "c", "", "generate config")
	flag.StringVar(&config, "config", "", "generate config")
	flag.StringVar(&lang, "l", "", "generate language")
	flag.StringVar(&lang, "lang", "", "generate language")
	flag.StringVar(&input, "i", "", "generate input path")
	flag.StringVar(&input, "input", "", "generate input path")
	flag.StringVar(&output, "o", "", "generate output path")
	flag.StringVar(&output, "output", "", "generate output path")
	flag.BoolVar(&dryrun, "d", false, "dryrun")
	flag.BoolVar(&dryrun, "dryrun", false, "dryrun")
	flag.Parse()

	if config != "" {
		if err := runConfig(config); err != nil {
			panic(err)
		}
	} else {
		if err := runCommand(lang, input, output, dryrun); err != nil {
			panic(err)
		}
	}
}

func runCommand(lang, input, output string, dryrun bool) error {
	g, err := proto.GetGenerator(lang)
	if err != nil {
		return err
	}
	files, err := proto.FindProtoFiles(input)
	if err != nil {
		return err
	}
	if dryrun {
		out, err := g.Generate(files)
		if err != nil {
			return err
		}
		for _, o := range out {
			fmt.Println("OutputPath:", filepath.Join(output, o.Name))
			fmt.Println(o.Content)
		}
		return nil
	}

	return g.GenerateAndOutput(files, output)
}

func runConfig(path string) error {
	dir := filepath.Dir(path)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	tasks, err := proto.ParseConfig(buf, dir)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if err := task.Run(); err != nil {
			return err
		}
	}
	return nil
}
