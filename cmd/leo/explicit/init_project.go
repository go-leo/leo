package explicit

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func initProject() {
	root := path.Base(moduleName)
	layout := []string{
		"api",
		"bin",
		"build",
		"cmd",
		"deployments",
		"doc",
		"internal",
		"pkg",
		"scripts",
		"tools",
	}
	for _, dir := range layout {
		err := os.MkdirAll(path.Join(root, dir), 0755)
		if err != nil {
			panic(err)
		}
	}

	command := exec.Command("cd", root, "&&", "go", "mod", "init", moduleName)
	err := command.Run()
	if err != nil {
		data, _ := command.Output()
		fmt.Println(string(data))
		panic(err)
	}
}
