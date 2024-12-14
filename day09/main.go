package day09

import (
	"fmt"
	"riemer/utils"
	"sort"
)

type File struct {
	id    int
	start int
	size  int
}

const emptySpace = -1

func Process() {
	lines := utils.ReadFile("day09/input.txt")
	fs1 := expandFilesystem(lines[0])
	fs2 := fs1
	files := identifyFiles(lines[0])

	fs1 = defragmentFileSystemPart1(fs1)
	part1Total := calculateFilesystemHash(fs1)

	fs2 = defragmentFileSystemPart2(fs2, files)
	part2Total := calculateFilesystemHash(fs2)

	fmt.Println("Day 9 Results")
	fmt.Println("Part1", part1Total)
	fmt.Println("Part2", part2Total)
}

func defragmentFileSystemPart1(filesystem []int) []int {
	start, end := 0, len(filesystem)-1

	for end > start {
		// Find first empty space
		for start < len(filesystem) && filesystem[start] != emptySpace {
			start++
		}

		// Find last file entry
		for end >= 0 && filesystem[end] == emptySpace {
			end--
		}

		// Swap
		if start < end {
			filesystem[start], filesystem[end] = filesystem[end], emptySpace
			start++
			end--
		}
	}

	return filesystem
}

func defragmentFileSystemPart2(filesystem []int, files []File) []int {
	// Sort files DESC
	sort.Slice(files, func(i, j int) bool {
		return files[i].id > files[j].id
	})

	for _, file := range files {
		freeSpaceStart := findFreeSpace(filesystem, file.size)

		if freeSpaceStart != -1 && freeSpaceStart < file.start {
			copy(filesystem[freeSpaceStart:freeSpaceStart+file.size], filesystem[file.start:file.start+file.size])
			for i := file.start; i < file.start+file.size; i++ {
				filesystem[i] = emptySpace
			}
		}
	}

	return filesystem
}

// findFreeSpace - find the most left free space available for the given size
func findFreeSpace(filesystem []int, size int) int {
	count := 0
	for i, value := range filesystem {
		if value == emptySpace {
			count++
			if count == size {
				return i - size + 1
			}
		} else {
			count = 0
		}
	}
	return -1
}

// identifyFiles - parses the input string and returns it as an array of File info
func identifyFiles(input string) []File {
	var files []File
	fileId := 0
	start := 0

	for i := 0; i < len(input); i += 2 {
		size := int(input[i] - '0')
		files = append(files, File{id: fileId, start: start, size: size})
		fileId++
		start += size
		// Skip free space
		if i+1 < len(input) {
			start += int(input[i+1] - '0')
		}
	}

	return files
}

// expandFilesystem - parses the input string and expands it into the filesystem
func expandFilesystem(input string) []int {
	totalFilesystemSize := 0
	for _, char := range input {
		totalFilesystemSize += int(char - '0')
	}

	result := make([]int, totalFilesystemSize)
	fileId := 0
	isFileMode := true
	index := 0

	for _, char := range input {
		value := emptySpace
		if isFileMode {
			value = fileId
			fileId++
		}

		// Fill the result slice with the current value
		length := int(char - '0')
		for i := 0; i < length; i++ {
			result[index] = value
			index++
		}

		isFileMode = !isFileMode
	}

	return result
}

// calculateFilesystemHash - computes the hash of the filesystem
func calculateFilesystemHash(filesystem []int) int {
	sum := 0
	for i, value := range filesystem {
		if value != emptySpace {
			sum += i * value
		}
	}
	return sum
}
