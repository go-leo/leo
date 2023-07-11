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
		newSource(path.Join("api", app), "", ""),

		newSource(path.Join("build", app), "", ""),

		newSource(path.Join("deployments", app), "", ""),

		newSource(path.Join("cmd", app), cmdWireContent, "wire.go"),
		newSource(path.Join("cmd", app), cmdMainContent, "main.go"),

		newSource(path.Join("internal", app), appRootWireContent, "wire.go"),

		newSource(path.Join("internal", app, "presentation"), presentationWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "assembler"), presentationAssemblerWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "assembler"), presentationAssemblersContent, "assemblers.go"),
		newSource(path.Join("internal", app, "presentation", "console"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "controller"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "provider"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "subscriber"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "presentation", "runner"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", app, "application", "command"), applicationCommandCommandsContent, "bus.go"),
		newSource(path.Join("internal", app, "application", "command"), applicationCommandWireContent, "wire.go"),
		newSource(path.Join("internal", app, "application", "query"), applicationQueryQueriesContent, "bus.go"),
		newSource(path.Join("internal", app, "application", "query"), applicationQueryWireContent, "wire.go"),
		newSource(path.Join("internal", app, "application", "service"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "application"), applicationWireContent, "wire.go"),

		newSource(path.Join("internal", app, "domain"), domainWireContent, "wire.go"),
		newSource(path.Join("internal", app, "domain", "model"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "domain", "service"), sampleWireContent, "wire.go"),

		newSource(path.Join("internal", app, "infrastructure"), infrastructureContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "client"), infrastructureClientWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "client", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", app, "infrastructure", "client", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "publisher"), infrastructurePublisherWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "publisher", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", app, "infrastructure", "publisher", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "repository"), infrastructureRepositoryWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "repository", "port"), docContent, "doc.go"),
		newSource(path.Join("internal", app, "infrastructure", "repository", "adapter"), sampleWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "converter"), infrastructureConverterWireContent, "wire.go"),
		newSource(path.Join("internal", app, "infrastructure", "converter"), infrastructureConvertersContent, "converters.go"),
	}

	if sample {
		newSource(path.Join("internal", app, "presentation", "runner"), presentationRunnerHelloContent, appBaseName+".go")
		newSource(path.Join("internal", app, "presentation", "runner"), presentationRunnerRunnersContent, "runners.go")
		newSource(path.Join("internal", app, "presentation", "runner"), presentationRunnerWireContent, "wire.go")
	}

	if http {
		sources = append(sources, newSource(path.Join("api", app), apiGrpcServiceContent, appBaseName+".proto"))
	}

	if grpc {
		sources = append(sources, newSource(path.Join("api", app), apiGrpcServiceContent, appBaseName+".proto"))
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

	if err := exec.Command("make", "wire_gen").Run(); err != nil {
		fmt.Println(err)
	}

	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		fmt.Println(err)
	}
}
