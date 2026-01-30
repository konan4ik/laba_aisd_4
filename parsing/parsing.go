package parsing

import (
	"bufio"
	"fmt"
	r "laba4/recs"
	"os"
	"strconv"
	"strings"
)

func FillFile(arr []string, path string, n int) {
	f, _ := os.Create(path)
	defer f.Close()

	w := bufio.NewWriter(f)
	for k := 0; k < n; k++ {
		if _, err := w.WriteString(arr[k]); err != nil {
			panic(err)
		}
	}

	w.Flush()
}

func ParseFile(path string, n int) []r.Record {
	array := make([]r.Record, n)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return array
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan() && i < n; i++ {
		line := scanner.Text()
		if line == "" {
			return array
		}
		splittedLine := strings.Split(line, "\t")
		number, _ := strconv.Atoi(splittedLine[2])
		array[i] = r.Record{
			Date:     r.CreateDate(splittedLine[0]),
			FullName: r.CreateFullName(splittedLine[1]),
			Number:   number,
			Descrp:   splittedLine[3],
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка:", err)
	}

	return array
}
