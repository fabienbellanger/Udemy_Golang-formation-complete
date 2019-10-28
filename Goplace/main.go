// Udemy - Goplace project
//
// Result:
// == Summary ==
// Number of occurrences of Go: 10
// Number of lines: 7
// Lines: [ 1 - 8 - 15 - 17 - 19 - 23 - 28 ]
// == End of Summary ==

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	occ, lines, err := FindReplaceFile("wikigo.txt", "Go", "Python")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Println("\n== Summary ==")
	fmt.Printf("Number of occurrences of Go: %d\n", occ)
	fmt.Printf("            Number of lines: %d\n", len(lines))
	fmt.Printf("                      Lines: %v\n", lines)
	fmt.Println("== End of Summary ==")
}

// FindReplaceFile function
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	file, err := os.Open(src)
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		found, res, nb := ProcessLine(scanner.Text(), old+" ", new+" ")

		if found {
			occ += nb
			lines = append(lines, i)
		}
		fmt.Println(res)

		i++
	}

	return
}

// ProcessLine function
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)
	if found {
		occ = strings.Count(line, old)
		res = strings.ReplaceAll(line, old, new)
	} else {
		res = line
	}

	return
}
