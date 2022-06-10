package main

import "fmt"

type inode interface {
	search(string)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type folder struct {
	name     string
	children []inode
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.children {
		composite.search(keyword)
	}
}

func (f *folder) add(c inode) {
	f.children = append(f.children, c)
}
