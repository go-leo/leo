package execx

import (
	"os"
	"os/exec"
	"path/filepath"
)

func StartProcess() (*os.Process, error) {
	// 可执行文件的路径
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}
	// 使用EvalSymlinks获取真实路径
	realPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}
	dir, _ := os.Getwd()
	cmd := exec.Cmd{
		Path:   realPath,
		Args:   os.Args,
		Env:    os.Environ(),
		Dir:    dir,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Start()
	return cmd.Process, err
}
