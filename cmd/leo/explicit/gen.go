package explicit

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var (
	//go:embed tmpl/cmd_wire.go.template
	cmdWireContent string

	//go:embed tmpl/cmd_main.go.template
	cmdMainContent string
)

//go:embed tmpl/app_root_wire.go.template
var appRootWireContent string

var (
	//go:embed tmpl/presentation_wire.go.template
	presentationWireContent string

	//go:embed tmpl/presentation_bus_wire.go.template
	presentationBusWireContent string

	//go:embed tmpl/presentation_bus_commands.go.template
	presentationBusCommandsContent string

	//go:embed tmpl/presentation_bus_queries.go.template
	presentationBusQueriesContent string

	//go:embed tmpl/presentation_assemblers.go.template
	presentationAssemblersContent string

	//go:embed tmpl/presentation_assembler_wire.go.template
	presentationAssemblerWireContent string
)

//go:embed tmpl/application_wire.go.template
var applicationContent string

//go:embed tmpl/domain_wire.go.template
var domainWireContent string

var (
	//go:embed tmpl/infrastructure_wire.go.template
	infrastructureContent string

	//go:embed tmpl/infrastructure_client_wire.go.template
	infrastructureClientWireContent string

	//go:embed tmpl/infrastructure_publisher_wire.go.template
	infrastructurePublisherWireContent string

	//go:embed tmpl/infrastructure_repository_wire.go.template
	infrastructureRepositoryWireContent string

	//go:embed tmpl/infrastructure_converters.go.template
	infrastructureConvertersContent string

	//go:embed tmpl/infrastructure_converter_wire.go.template
	infrastructureConverterWireContent string
)

//go:embed tmpl/doc.go.template
var docContent string

//go:embed tmpl/sample_wire.go.template
var sampleWireContent string

//go:embed tmpl/tools_wire.go.template
var toolsWireContext string

func gen() {
	//goModInit := exec.Command("go", "mod", "init", *moduleName)
	//if err := goModInit.Run(); err != nil {
	//	log.Fatalf("goModInit.Run() failed with %s\n", err)
	//}
	sources := []*Source{
		newSource(path.Join("api", appPath), "", ""),

		newSource(path.Join("build", appPath), "", ""),

		newSource(path.Join("deployments", appPath), "", ""),

		newSource(path.Join("cmd", appPath), cmdWireContent, "wire.go"),
		newSource(path.Join("cmd", appPath), cmdMainContent, "main.go"),

		newSource(path.Join("internal", appPath), appRootWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "config"), "", ""),

		newSource(path.Join("internal", appPath, "presentation"), presentationWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "bus"), presentationBusWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "bus"), presentationBusCommandsContent, "commands.go"),
		newSource(path.Join("internal", appPath, "presentation", "bus"), presentationBusQueriesContent, "queries.go"),
		newSource(path.Join("internal", appPath, "presentation", "assembler"), presentationAssemblerWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "assembler"), presentationAssemblersContent, "assemblers.go"),
		newSource(path.Join("internal", appPath, "presentation", "console"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "controller"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "provider"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "subscriber"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "application"), applicationContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application", "command"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application", "query"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application", "service"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "domain"), domainWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "domain", "model"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "domain", "service"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "infrastructure"), infrastructureContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "client"), infrastructureClientWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "client", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "client", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "publisher"), infrastructurePublisherWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "publisher", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "publisher", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "repository"), infrastructureRepositoryWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "repository", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "repository", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "converter"), infrastructureConverterWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "infrastructure", "converter"), infrastructureConvertersContent, "converters.go"),
	}

	for _, src := range sources {
		err := checkNotExist(src.DirPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}

	others := []*Source{
		newSource(path.Join("tools"), toolsWireContext, "wire.go"),
	}
	for _, src := range sources {
		err := checkNotExist(src.FilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}

	for _, src := range sources {
		err := createSource(src)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}

	for _, src := range others {
		err := createSource(src)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return
		}
	}

	//goModTidy := exec.Command("go", "mod", "tidy")
	//if err := goModTidy.Run(); err != nil {
	//	log.Fatalf("goModTidy.Run() failed with %s\n", err)
	//}
}

func newSource(dirPath string, text string, name string) *Source {
	src := &Source{
		DirPath:  dirPath,
		Text:     text,
		Name:     name,
		FilePath: filepath.Join(dirPath, name),
		data: &SourceData{
			ModuleName:  moduleName,
			AppPath:     appPath,
			AppBaseName: filepath.Base(appPath),
			Package:     filepath.Base(dirPath),
		},
	}
	return src
}

type SourceData struct {
	ModuleName   string
	AppPath      string
	AppBaseName  string
	WireFilePath string
	Package      string
}

type Source struct {
	DirPath  string
	Text     string
	Name     string
	FilePath string
	data     *SourceData
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

func createSource(src *Source) error {
	err := os.MkdirAll(src.DirPath, 0777)
	if err != nil {
		return err
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
	source, err := format.Source(buffer.Bytes())
	if err != nil {
		return err
	}
	_, err = file.Write(source)
	if err != nil {
		return err
	}
	return nil
}
