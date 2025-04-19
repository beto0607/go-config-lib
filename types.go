package goconfiglib

type Settings struct {
	UseXDGConfigHome bool
}

type Configs struct {
	FilePath string
	Root     Section
}

type Section struct {
	Parent      *Section
	RawValue    string
	Name        string
	Subsections []Section
	Values      map[string]string
}
