package explicit

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
)

func addService() {
	sources := []*Source{
		newSource(path.Join("api", appPath), "", ""),

		newSource(path.Join("build", appPath), "", ""),

		newSource(path.Join("deployments", appPath), "", ""),

		newSource(path.Join("cmd", appPath), cmdWireContent, "wire.go"),
		newSource(path.Join("cmd", appPath), cmdMainContent, "main.go"),

		newSource(path.Join("internal", appPath), appRootWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "presentation"), presentationWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "assembler"), presentationAssemblerWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "assembler"), presentationAssemblersContent, "assemblers.go"),
		newSource(path.Join("internal", appPath, "presentation", "console"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "controller"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "provider"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "subscriber"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "presentation", "runner"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", appPath, "application", "command"), applicationCommandCommandsContent, "bus.go"),
		newSource(path.Join("internal", appPath, "application", "command"), applicationCommandWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application", "query"), applicationQueryQueriesContent, "bus.go"),
		newSource(path.Join("internal", appPath, "application", "query"), applicationQueryWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application", "service"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", appPath, "application"), applicationWireContent, "wire.go"),

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

	if sample {
		sources = append(sources, newSource(path.Join("internal", appPath, "presentation", "runner"), presentationRunnerHelloContent, appBaseName+".go"))
		sources = append(sources, newSource(path.Join("internal", appPath, "presentation", "runner"), presentationRunnerRunnersContent, "runners.go"))
		sources = append(sources, newSource(path.Join("internal", appPath, "presentation", "runner"), presentationRunnerWireContent, "wire.go"))
	}

	if http {
		sources = append(sources, newSource(path.Join("api", appPath), apiHttpHelloContent, appBaseName+".go"))
		sources = append(sources, newSource(path.Join("internal", appPath, "presentation", "controller"), presentationControllerHelloContent, appBaseName+".go"))
		sources = append(sources, newSource(path.Join("internal", appPath, "presentation", "controller"), presentationControllerWireContent, "wire.go"))
	}

	if grpc {
		sources = append(sources, newSource(path.Join("api", appPath), apiGrpcServiceContent, appBaseName+".proto"))
	}

	for _, src := range sources {
		err := checkNotExist(src.DirPath)
		if err != nil {
			fmt.Println(src.FilePath, src.Name, err)
			os.Exit(1)
			return
		}
	}

	for _, src := range sources {
		err := checkNotExist(src.FilePath)
		if err != nil {
			fmt.Println(src.FilePath, src.Name, err)
			os.Exit(1)
			return
		}
	}

	for _, src := range sources {
		err := src.createSource()
		if err != nil {
			fmt.Println(src.FilePath, src.Name, err)
			os.Exit(1)
			return
		}
	}

	if err := exec.Command("make", "go_gen").Run(); err != nil {
		fmt.Println(err)
	}

	if err := exec.Command("make", "wire_gen").Run(); err != nil {
		fmt.Println(err)
	}

	if err := exec.Command("go", "get", "codeup.aliyun.com/qimao/leo/leo/...@master").Run(); err != nil {
		fmt.Println(err)
	}

	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		fmt.Println(err)
	}
}
