package internal

import (
	"google.golang.org/protobuf/compiler/protogen"
	"path"
)

type Package struct {
	name     string
	fullName string
	absPath  string
	relPath  string
}

func (p *Package) Name() string {
	return p.name
}

func (p *Package) FullName() string {
	return p.fullName
}

func (p *Package) AbsPath() string {
	return p.absPath
}

func (p *Package) RelPath() string {
	return p.relPath
}

func (p *Package) GoImportPath() protogen.GoImportPath {
	return protogen.GoImportPath(p.FullName())
}

func NewPackage(absPath string, relPath string, fullName string) *Package {
	return &Package{absPath: absPath, relPath: relPath, fullName: fullName, name: path.Base(fullName)}
}
