package main

import "fmt"

const NMAX int = 10

type array [NMAX]int

func main() {
	var inputArr, rankedArr, gitsArr array
	var jumPTertinggi, jumRank, jumGits int

	fmt.Print("Masukkan jumlah pemain skor tertinggi: ")
	fmt.Scan(&jumPTertinggi)

	fmt.Print("Masukkan skor-skor tertinggi para pemain: ")
	inputArray(&inputArr, jumPTertinggi)
	sortArray(&inputArr, jumPTertinggi)
	rankArray(inputArr, &rankedArr, jumPTertinggi, &jumRank)

	fmt.Print("Masukkan jumlah skor setiap permainan GITS: ")
	fmt.Scan(&jumGits)
	fmt.Print("Masukkan skor-skor GITS: ")
	inputArray(&gitsArr, jumGits)
	rankOutput(rankedArr, gitsArr, jumRank, jumGits)
}

func inputArray(arr *array, jumAnggota int) {
	var i int
	for i = 0; i < jumAnggota; i++ {
		fmt.Scan(&(*arr)[i])
	}
}

func sortArray(arr *array, jumAnggota int) {
	var i, j, minIdx, temp int

	for i = 0; i < jumAnggota-1; i++ {
		minIdx = i

		for j = i + 1; j < jumAnggota; j++ {
			if (*arr)[j] > (*arr)[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[minIdx]
			(*arr)[minIdx] = temp
		}
	}
}

func rankArray(arrIn array, arrOut *array, jumAnggota int, jumRank *int) {
	var i, j int
	var found bool
	*jumRank = 0

	for i = 0; i < jumAnggota; i++ {
		if *jumRank == 0 {
			(*arrOut)[*jumRank] = arrIn[i]
			*jumRank++
		} else {
			found = false
			j = 0
			for j < *jumRank && !found {
				if arrIn[i] == (*arrOut)[j] {
					found = true
				}
				j++
			}
			if !found {
				(*arrOut)[*jumRank] = arrIn[i]
				*jumRank++
			}
		}
	}
}

func rankOutput(rankedArr, gitsArr array, jumRank, jumGits int) {
	var i, j, rank int
	var found bool

	for i = 0; i < jumGits; i++ {
		j = 0
		found = false

		for j < jumRank && !found {
			if gitsArr[i] >= rankedArr[j] {
				found = true
			} else {
				j++
			}
		}

		if found {
			rank = j + 1
		} else {
			rank = jumRank + 1
		}

		fmt.Print(rank, " ")
	}
}
