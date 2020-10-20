package main

import (
	"fmt"
)

type Input struct {
	Path string
	Name string
}

func (i *Input) getDeep() (counter int) {
	for _, c := range i.Path {
		if c == '/' {
			counter++
		}
	}
	return
}

func (i *Input) getPathName() string {
	return i.Path + i.Name
}

func main() {
	input := []Input{
		{"/a/b/c/", "e.txt"},
		{"/a/b/c/", "d.txt"},
		{"/a/b/c/d", "e.txt"},
	}

	result := handle(input)
	fmt.Println(result)

	counter, path := countFiles(input)
	fmt.Println(counter, path)
}

func handle(input []Input) (result []string) {
	maxCounter := -1
	for _, eachInput := range input {
		counter := eachInput.getDeep()

		if counter > maxCounter {
			maxCounter = counter
			result = []string{eachInput.getPathName()}
		} else if counter == maxCounter {
			result = append(result, eachInput.getPathName())
		}
	}
	return result
}

type result struct {
	path    string
	counter int
}

func countFiles(input []Input) (counter int, path string) {
	foldersMap := make(map[string]*result)
	for _, eachInput := range input {
		if eachFolder, ok := foldersMap[eachInput.Path]; ok {
			eachFolder.counter++
		} else {
			foldersMap[eachInput.Path] = &result{eachInput.Path, 1}
		}
	}

	maxCount := 0
	var maxPath *result
	for _, eachFolder := range foldersMap {
		if eachFolder.counter > maxCount {
			maxCount = eachFolder.counter
			maxPath = eachFolder
		}
	}

	return maxPath.counter, maxPath.path
}
