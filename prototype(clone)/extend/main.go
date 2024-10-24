package main

import "fmt"

type INode interface {
	Clone() INode
	Print(string)
}

type File struct {
	Name string
}

func (f *File) Clone() INode {
	return &File{
		Name: f.Name,
	}
}

func (f *File) Print(s string) {
	fmt.Println(s + f.Name + "_clone")
}

type Folder struct {
	Childrens []INode
	Name      string
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name + "_clone")
	for _, v := range f.Childrens {
		v.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	tempFolder := make([]INode, len(f.Childrens))

	for i, v := range f.Childrens {
		tempFolder[i] = v.Clone()
	}
	return &Folder{
		Childrens: tempFolder,
		Name:      f.Name,
	}
}

func main() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}
	folder1 := &Folder{
		Childrens: []INode{file1, file2},
		Name:      "Folder1",
	}
	folder2 := &Folder{
		Childrens: []INode{file3, folder1},
		Name:      "Folder2",
	}

	cloneFolder := folder2.Clone()
	fmt.Println("Printing for clone folder2")
	cloneFolder.Print("  ")
}
