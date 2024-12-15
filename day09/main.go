package day09

import (
	"fmt"
	"riemer/utils"
	"sort"
)

type File struct {
	id, start, size int
}

type EmptySpace struct {
	start, size int
}

const emptySpace = -1

func Process() {
	lines := utils.ReadFile("day09/input.txt")
	fs := expandFilesystem(lines[0])
	files, emptyspaces := identifyFiles(lines[0])

	defragmentFileSystemPart1(fs)
	part1Total := calculateFilesystemHash(fs)

	fs = expandFilesystem(lines[0])
	defragmentFileSystemPart2(fs, files, emptyspaces)
	part2Total := calculateFilesystemHash(fs)

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

func defragmentFileSystemPart2(filesystem []int, files []File, emptySpaces []EmptySpace) []int {
	// Sort files DESC
	sort.Slice(files, func(i, j int) bool {
		return files[i].id > files[j].id
	})

	for _, file := range files {
		for i, space := range emptySpaces {
			if space.size >= file.size && space.start < file.start {
				copy(filesystem[space.start:space.start+file.size], filesystem[file.start:file.start+file.size])
				for j := file.start; j < file.start+file.size; j++ {
					filesystem[j] = emptySpace
				}
				// Update emptySpaces
				emptySpaces[i].size -= file.size
				emptySpaces[i].start += file.size
				break
			}
		}
	}

	return filesystem
}

// identifyFiles - parses the input string and returns it as arrays of File and EmptySpace info
func identifyFiles(input string) ([]File, []EmptySpace) {
	var files []File
	var emptySpaces []EmptySpace
	fileId, start := 0, 0

	for i := 0; i < len(input); i += 2 {
		size := int(input[i] - '0')
		files = append(files, File{id: fileId, start: start, size: size})
		fileId++
		start += size

		// Identify empty space
		if i+1 < len(input) {
			emptySize := int(input[i+1] - '0')
			if emptySize > 0 {
				emptySpaces = append(emptySpaces, EmptySpace{start, emptySize})
			}
			start += emptySize
		}
	}

	return files, emptySpaces
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
