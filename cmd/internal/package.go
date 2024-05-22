package internal

import "google.golang.org/protobuf/compiler/protogen"

type Package struct {
	absPath  string
	relPath  string
	fullName string
	name     string
}

func (p *Package) Name() string {
	return p.name
}

func (p *Package) FullName() string {
	return p.fullName
}

func (p *Package) RelPath() string {
	return p.relPath
}

func (p *Package) GoImportPath() protogen.GoImportPath {
	return protogen.GoImportPath(p.FullName())
}

func NewPackage(absPath string, relPath string, fullName string, name string) *Package {
	return &Package{absPath: absPath, relPath: relPath, fullName: fullName, name: name}
}
