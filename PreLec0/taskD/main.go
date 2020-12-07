package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//TPP : D task Solution
func main() {
	var n, m int

	scannerInput := bufio.NewScanner(os.Stdin)

	scannerInput.Scan()
	n, err := strconv.Atoi(scannerInput.Text())
	if err != nil {
		log.Fatal(err)
	}

	scannerInput.Scan()
	m, err = strconv.Atoi(scannerInput.Text())
	if err != nil {
		log.Fatal(err)
	}

	listExist := make([]string, n)
	listWanted := make([]string, m)

	for i := 0; i < n; i++ {
		scannerInput.Scan()
		listExist[i] = scannerInput.Text()
	}

	for i := 0; i < m; i++ {
		scannerInput.Scan()
		listWanted[i] = scannerInput.Text()
	}

	for _, wantedFilm := range listWanted {
		found := false
		for _, existFilm := range listExist {
			if wantedFilm == existFilm {
				found = true
				break
			}
		}
		if found {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}
	}

}
