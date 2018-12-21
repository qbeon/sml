package ast

// InputFile represents an input file
type InputFile struct {
	Name     str
	Path     str
	FullPath str
	Contents str
}

// Clone creates a clone of the input file
func (inf *InputFile) Clone() InputFile {
	name := make(str, len(inf.Name))
	path := make(str, len(inf.Path))
	fullPath := make(str, len(inf.FullPath))
	contents := make(str, len(inf.Contents))
	copy(name, inf.Name)
	copy(path, inf.Path)
	copy(fullPath, inf.FullPath)
	copy(contents, inf.Contents)
	return InputFile{
		Name:     name,
		Path:     path,
		FullPath: fullPath,
		Contents: contents,
	}
}
