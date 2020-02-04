package proto

import (
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type runConfig struct {
	Input  string     `yaml:"input"`
	Inputs []string   `yaml:"inputs"`
	Lang   langConfig `yaml:"lang"`
}

type langConfig struct {
	CsConf         *csConfig `yaml:"cs"`
	CSharpConf     *csConfig `yaml:"csharp"`
	GoConf         *goConfig `yaml:"go"`
	GolangConf     *goConfig `yaml:"golang"`
	JsConf         *jsConfig `yaml:"js"`
	JavaScriptConf *jsConfig `yaml:"javascript"`
}

type csConfig struct {
	Output         string `yaml:"output"`
	SkipSerializer bool   `yaml:"skip_serializer"`
	Property       bool   `yaml:"property"`
	Serializable   bool   `yaml:"serializable"`
}

type goConfig struct {
	Output         string `yaml:"output"`
	Root           string `yaml:"root"`
	SkipSerializer bool   `yaml:"skip_serializer"`
}

type jsConfig struct {
	Output                  string `yaml:"output"`
	SkipSerializer          bool   `yaml:"skip_serializer"`
	UseTypeScript           bool   `yaml:"use_ts"`
	DisablePackageNameToDir bool   `yaml:"disable_package_to_dir"`
}

func ParseConfig(buf []byte, basePath string) ([]*GenerateTask, error) {
	var config runConfig
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return nil, err
	}
	files, err := config.getFiles(basePath)
	if err != nil {
		return nil, err
	}
	ret := make([]*GenerateTask, 0)
	lang := config.Lang
	if lang.CsConf != nil {
		ret = append(ret, lang.CsConf.createGenerateTask(files, basePath))
	}
	if lang.CSharpConf != nil {
		ret = append(ret, lang.CSharpConf.createGenerateTask(files, basePath))
	}
	if lang.GoConf != nil {
		ret = append(ret, lang.GoConf.createGenerateTask(files, basePath))
	}
	if lang.GolangConf != nil {
		ret = append(ret, lang.GolangConf.createGenerateTask(files, basePath))
	}
	if lang.JsConf != nil {
		ret = append(ret, lang.JsConf.createGenerateTask(files, basePath)...)
	}
	if lang.JavaScriptConf != nil {
		ret = append(ret, lang.JavaScriptConf.createGenerateTask(files, basePath)...)
	}
	return ret, nil
}

func (c *runConfig) getFiles(basePath string) ([]string, error) {
	list := make(map[string]string, 0)

	if c.Input != "" {
		files, err := FindProtoFiles(filepath.Join(basePath, c.Input))
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			list[f] = f
		}
	}
	if len(c.Inputs) > 0 {
		for _, input := range c.Inputs {
			files, err := FindProtoFiles(filepath.Join(basePath, input))
			if err != nil {
				return nil, err
			}
			for _, f := range files {
				list[f] = f
			}
		}
	}
	ret := make([]string, 0)
	for _, v := range list {
		ret = append(ret, v)
	}
	return ret, nil
}

func (c *csConfig) createGenerateTask(files []string, basePath string) *GenerateTask {
	g := NewCSGenerator()
	g.SkipSerializer = c.SkipSerializer
	g.Property = c.Property
	g.Serializable = c.Serializable
	return &GenerateTask{
		Output:    filepath.Join(basePath, c.Output),
		Files:     files,
		Generator: g,
	}
}

func (c *goConfig) createGenerateTask(files []string, basePath string) *GenerateTask {
	g := NewGoGenerator()
	g.SkipSerializer = c.SkipSerializer
	g.PackageRoot = c.Root
	return &GenerateTask{
		Output:    filepath.Join(basePath, c.Output),
		Files:     files,
		Generator: g,
	}
}

func (c *jsConfig) createGenerateTask(files []string, basePath string) []*GenerateTask {
	ret := make([]*GenerateTask, 0)
	{
		g := NewJSGenerator()
		g.SkipSerializer = c.SkipSerializer
		g.PackageNameToDirectory = !c.DisablePackageNameToDir
		ret = append(ret, &GenerateTask{
			Output:    filepath.Join(basePath, c.Output),
			Files:     files,
			Generator: g,
		})
	}
	if c.UseTypeScript {
		g := NewTSGenerator()
		g.SkipSerializer = c.SkipSerializer
		g.PackageNameToDirectory = !c.DisablePackageNameToDir
		ret = append(ret, &GenerateTask{
			Output:    filepath.Join(basePath, c.Output),
			Files:     files,
			Generator: g,
		})
	}
	return ret
}
