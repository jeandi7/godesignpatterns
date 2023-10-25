package creational

type File struct {
	name string
}

func (f *File) search(keyword string) string {
	return f.name + "(" + keyword + ")"
}

func (f *File) getName() string {
	return f.name
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) string {
	s := f.name + "(" + keyword + ")"
	for _, composite := range f.components {
		s = s + composite.search(keyword)
	}
	return s
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

type Component interface {
	search(string) string
}
