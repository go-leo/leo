package explicit

import (
	"bytes"
	"errors"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

func newSource(dirPath string, text string, name string) *Source {
	src := &Source{
		DirPath:  dirPath,
		Text:     text,
		Name:     name,
		FilePath: filepath.Join(dirPath, name),
		data: &SourceData{
			ModuleName:       module,
			AppPath:          appPath,
			AppPathDot:       appPathDot,
			AppBaseName:      appBaseName,
			AppUpperBaseName: appUpperBaseName,
			ServiceName:      serviceName,
			Package:          filepath.Base(dirPath),
			Sample:           sample,
			HTTP:             http,
			GRPC:             grpc,
			Schedule:         schedule,
			Stream:           stream,
		},
	}
	return src
}

type SourceData struct {
	ModuleName       string
	AppPath          string
	AppPathDot       string
	AppBaseName      string
	WireFilePath     string
	Package          string
	AppUpperBaseName string
	ServiceName      string
	Sample           bool
	HTTP             bool
	GRPC             bool
	Schedule         bool
	Stream           bool
}

type Source struct {
	DirPath  string
	Text     string
	Name     string
	FilePath string
	data     *SourceData
}

func (src *Source) createSource() error {
	if len(src.DirPath) > 0 {
		err := os.MkdirAll(src.DirPath, 0o777)
		if err != nil {
			return err
		}
	}

	if len(src.Name) == 0 && len(src.Text) == 0 {
		return nil
	}

	tmpl, err := template.New(src.FilePath).Parse(src.Text)
	if err != nil {
		return err
	}
	file, err := os.Create(src.FilePath)
	if err != nil {
		return err
	}
	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, src.data)
	if err != nil {
		return err
	}
	if path.Ext(src.FilePath) == ".go" {
		source, err := format.Source(buffer.Bytes())
		if err != nil {
			return err
		}
		_, err = file.Write(source)
		return err
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return file.Close()
}

func checkNotExist(name string) error {
	_, err := os.Stat(name)
	if err == nil {
		return errors.New(name + " exist")
	}
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}
