package aoc

import (
	"aoc2022/aoc/virtualfs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day7() {
	file, _ := os.Open("input/aocday7.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fs := virtualfs.CreateFS()
	currentDir := fs.Root
	listing := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$ ") {
			cmd := strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(cmd, "cd ") {
				folderName := strings.TrimPrefix(cmd, "cd ")
				switch folderName {
				case "..":
					currentDir = currentDir.Parent()
				case "/":
					currentDir = fs.Root
				default:
					currentDir = (*currentDir).GetOrCreateFolder(folderName)
				}
				listing = false
			} else if cmd == "ls" {
				listing = true
			} else {
				listing = false
			}
		} else if listing {
			split := strings.Split(line, " ")
			if split[0] != "dir" {
				currentDir.PutFile(split[1], ToInt(split[0]))
			}
		}
	}

	sumOf100k := 0
	ConsiderSize(fs.Root, &sumOf100k)

	totalSize := fs.Root.Size()
	neededToDelete := 30000000 - (70000000 - totalSize)
	smallest := -1
	ForEachDir(fs.Root, func(folder *virtualfs.Folder) {
		size := folder.Size()
		if size >= neededToDelete && (smallest == -1 || size < smallest) {
			smallest = size
		}
	})

	fmt.Printf("Sum of folders with at most 100k bytes: %d\n", sumOf100k)
	fmt.Printf("Smallest folder that needs to be deleted: %d", smallest)
}

func ConsiderSize(folder *virtualfs.Folder, size *int) {
	for _, file := range folder.Children {
		if file.Type() == virtualfs.FolderType {
			fSize := file.Size()
			if fSize < 100000 {
				*size += fSize
			}
			ConsiderSize(file.(*virtualfs.Folder), size)
		}
	}
}

func ForEachDir(folder *virtualfs.Folder, fun func(folder *virtualfs.Folder)) {
	for _, file := range folder.Children {
		if file.Type() == virtualfs.FolderType {
			folder := file.(*virtualfs.Folder)
			fun(folder)
			ForEachDir(folder, fun)
		}
	}
}
