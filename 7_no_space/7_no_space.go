package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	name           string
	size           int
	isDirectory    bool
	parent         *File
	subDirectories map[string]*File
}

var cursor *File

func printTree(root *File, prefix string) {
	fmt.Print(prefix + root.name)
	if root.isDirectory {
		fmt.Println("/")
		for _, f := range root.subDirectories {
			printTree(f, prefix+"    ")
		}
	} else {
		fmt.Println()
	}
}

func dirSize(root *File) int {
	size := 0
	if !root.isDirectory {
		return root.size
	}
	for _, f := range root.subDirectories {
		if f.isDirectory {
			size += dirSize(f)
		} else {
			size += f.size
		}
	}
	root.size = size
	return size
}

func pt1(root *File) int {
	size := 0
	if !root.isDirectory {
		return 0
	} else {
		if root.size <= 100000 {
			size += root.size
		} else {
			size += 0
		}
	}
	for _, f := range root.subDirectories {
		size += pt1(f)
	}
	return size
}

func pt2(toDel *File, sizeToDel int) *File {
	// fmt.Println("Checking dir " + toDel.name + " of size " + strconv.Itoa(toDel.size))
	if toDel.size < sizeToDel {
		// If too small, don't bother checking subdirs
		return nil
	} else {
		newFile := toDel
		// Check if any subdirs are smaller than the current dir
		// and big enough to delete
		for _, f := range toDel.subDirectories {
			if !f.isDirectory {
				// skip files
				continue
			}
			sub := pt2(f, sizeToDel)
			if sub != nil && sub.size <= newFile.size {
				newFile = sub
			}
		}
		return newFile
	}
}

func buildTree(stream string) *File {
	root := &File{
		"/", 0, true, nil, make(map[string]*File),
	}
	cursor = root

	for _, line := range strings.Split(stream, "\n") {
		tok := strings.Split(line, " ")
		// fmt.Println(line)

		if strings.Contains(line, "$") {
			// Command
			if strings.Contains(line, "cd") {
				nd := tok[2]
				// fmt.Println("'" + nd + "'")
				if nd == "/" {
					// fmt.Println("ignoring root cd, already there")
					continue
				} else if nd == ".." {
					// fmt.Println("moving cursor up")
					cursor = (*cursor).parent
				} else {
					// fmt.Println("moving cursor down to " + nd)
					cursor = (*cursor).subDirectories[nd]
				}
			}
		} else if strings.Contains(line, "dir") {
			// New dir just dropped
			f := &File{
				tok[1], 0, true, cursor, make(map[string]*File),
			}
			// Add to subdirs
			(*cursor).subDirectories[tok[1]] = f
		} else {
			// hehe no error here
			size, _ := strconv.Atoi(tok[0])

			f := &File{
				tok[1], size, false, cursor, nil,
			}
			(*cursor).subDirectories[tok[1]] = f
		}
	}
	return root
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	stream := string(content)
	root := buildTree(stream)
	// Populate directory sizes
	dirSize(root)

	fmt.Println("Pt 1:")
	fmt.Println(pt1(root))
	// printTree(root, "")
	spaceNeeded := 30000000 - (70000000 - root.size)
	fmt.Println("Pt 2:")
	fmt.Println("Need to delete ", spaceNeeded, " for update")
	toDel := pt2(root, spaceNeeded)
	fmt.Println("Delete ", toDel.name, " of size ", toDel.size)
}
