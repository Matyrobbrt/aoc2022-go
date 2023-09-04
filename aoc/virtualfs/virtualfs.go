package virtualfs

import "strings"

type ElementType int8

const (
	FileType   ElementType = 0
	FolderType ElementType = 1
)

type FSElement interface {
	Size() int
	Name() string
	Type() ElementType
	Parent() *Folder
}

type File struct {
	name   string
	size   int
	parent *Folder
}

type Folder struct {
	name     string
	Children map[string]FSElement
	parent   *Folder
}

type FileSystem struct {
	Root *Folder
}

func CreateFS() *FileSystem {
	return &FileSystem{
		Root: &Folder{"/", make(map[string]FSElement), nil},
	}
}

func (fs *FileSystem) Resolve(name string) FSElement {
	spl := strings.Split(name, "/")
	pathLength := len(spl)
	if pathLength == 0 || (pathLength == 1 && spl[0] == "") {
		return fs.Root
	}

	var folder = fs.Root

	for i := 0; i < len(spl); i++ {
		if folder == nil {
			panic("No such directory: " + spl[i-1])
		}
		child := folder.Children[spl[i]]
		if i < pathLength-1 {
			if child == nil {
				panic("No such directory: " + spl[i])
			}
			if child.Type() != FolderType {
				panic("Not a folder: " + spl[i])
			}

			f, _ := child.(Folder)
			folder = &f
		} else {
			return child
		}
	}
	return folder
}

//goland:noinspection GoReceiverNames,GoMixedReceiverTypes
func (folder *Folder) GetOrCreateFolder(childrenName string) *Folder {
	resolved := folder.Children[childrenName]

	if resolved == nil {
		newFldr := &Folder{childrenName, make(map[string]FSElement), folder}
		folder.Children[childrenName] = newFldr
		return newFldr
	} else if resolved.Type() != FolderType {
		panic(childrenName + " is not a folder")
	}
	return resolved.(*Folder)
}

func (f *Folder) PutFile(name string, size int) {
	existing := f.Children[name]
	if existing == nil {
		f.Children[name] = &File{name: name, size: size, parent: f}
	} else {
		panic("File " + name + " already exists!")
	}
}

func (fs *FileSystem) GetOrCreateFolder(name string) *Folder {
	resolved := fs.Resolve(name)

	if resolved == nil {
		spl := strings.Split(name, "/")
		folderName := spl[len(spl)-1]
		parent := fs.GetOrCreateFolder(strings.Join(spl[0:len(spl)-1], "/"))
		return parent.GetOrCreateFolder(folderName)
	} else if resolved.Type() != FolderType {
		panic(name + " is not a folder")
	}
	return resolved.(*Folder)
}

func (f File) Size() int {
	return f.size
}

func (f File) Name() string {
	return f.name
}

func (f File) Type() ElementType {
	return FileType
}

func (f File) Parent() *Folder {
	return f.parent
}

func (f Folder) Name() string {
	return f.Name()
}

func (f Folder) Type() ElementType {
	return FolderType
}

func (f Folder) Parent() *Folder {
	return f.parent
}

func (f Folder) Size() int {
	size := 0
	for _, el := range f.Children {
		size += el.Size()
	}
	return size
}
