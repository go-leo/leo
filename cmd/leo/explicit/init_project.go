package explicit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func initProject() {
	sources := []*Source{
		newSource(path.Join("api", appPath), "", ""),
		newSource(path.Join("build", appPath), "", ""),
		newSource(path.Join("deployments", appPath), "", ""),
		newSource(path.Join("cmd", appPath), "", ""),
		newSource(path.Join("internal", appPath), "", ""),

		newSource(path.Join("pkg"), pkgWireContent, "wire.go"),
		newSource(path.Join("pkg", "actuatorx"), pkgActuatorxConfigContent, "config.go"),
		newSource(path.Join("pkg", "configx"), pkgConfigxConfigurationContent, "config.go"),
		newSource(path.Join("pkg", "configx"), pkgConfigxLoadContent, "load.go"),
		newSource(path.Join("pkg", "configx"), pkgConfigxWireContent, "wire.go"),
		newSource(path.Join("pkg", "ginx"), pkgGinxConfigContent, "config.go"),
		newSource(path.Join("pkg", "ginx"), pkgGinxMiddlewareContent, "middleware.go"),
		newSource(path.Join("pkg", "grpcx"), pkggRPCxClientContent, "client.go"),
		newSource(path.Join("pkg", "grpcx"), pkggRPCxServerContent, "server.go"),
		newSource(path.Join("pkg", "grpcx"), pkggRPCxWireContent, "wire.go"),

		newSource(path.Join("scripts", "shell"), scriptsShellFormatContent, "format.sh"),
		newSource(path.Join("scripts", "shell"), scriptsShellGenContent, "gen.sh"),
		newSource(path.Join("scripts", "shell"), scriptsShellLintContent, "lint.sh"),
		newSource(path.Join("scripts", "shell"), scriptsShellProtocContent, "protoc.sh"),
		newSource(path.Join("scripts", "shell"), scriptsShellToolsContent, "tools.sh"),
		newSource(path.Join("scripts", "shell"), scriptsShellWireContent, "wire.sh"),

		newSource(path.Join("tools", appPath), toolsWireContent, "tools.go"),

		newSource(path.Join(), _MakefileContent, "Makefile"),
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

	initCmd := exec.Command("go", "mod", "init", module)
	if output, err := initCmd.CombinedOutput(); err != nil {
		fmt.Printf("combined out:\n%s\n", string(output))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	tidyCmd := exec.Command("go", "mod", "tidy")
	if output, err := tidyCmd.CombinedOutput(); err != nil {
		fmt.Printf("combined out:\n%s\n", string(output))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	getCmd := exec.Command("go", "get", "codeup.aliyun.com/qimao/leo/leo/...@master")
	if output, err := getCmd.CombinedOutput(); err != nil {
		fmt.Printf("combined out:\n%s\n", string(output))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
