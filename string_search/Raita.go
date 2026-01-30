package str_search

import (
	"fmt"
	p "laba4/parsing"
	r "laba4/recs"
	"time"
)

func restEquals(y []byte, yFrom int, x []byte, xFrom int, length int) bool {
	for i := 0; i < length; i++ {
		if y[yFrom+i] != x[xFrom+i] {
			return false
		}
	}
	return true
}

func preBmBc(x []byte) []int {
	const ASIZE = 256
	m := len(x)
	result := make([]int, ASIZE)
	for i := 0; i < ASIZE; i++ {
		result[i] = m
	}
	for i := 0; i <= m-2; i++ {
		result[int(x[i])] = m - i - 1
	}
	return result
}

func RaitaAll(x, y []byte) []int {
	m := len(x)
	n := len(y)
	var matches []int

	if m == 0 {
		return matches
	}
	if m == 1 {
		for i := 0; i < n; i++ {
			if y[i] == x[0] {
				matches = append(matches, i)
			}
		}
		return matches
	}

	bmBc := preBmBc(x)
	firstCh := x[0]
	middleCh := x[m/2]
	lastCh := x[m-1]

	j := 0
	for j <= n-m {
		c := y[j+m-1]
		if lastCh == c && middleCh == y[j+m/2] && firstCh == y[j] {
			if m <= 2 || restEquals(y, j+1, x, 1, m-2) {
				matches = append(matches, j)
			}
		}
		shift := bmBc[int(c)]
		if shift < 1 {
			shift = 1
		}
		j += shift
	}
	return matches
}

func RaitaTimed(x []byte, y []byte, arr []r.Record, n int) {
	start := time.Now()

	var output []string
	output = append(output, "Строка   Data\n")

	for i := 0; i < n; i++ {
		rec := arr[i]

		fullname := rec.FullName.Name + " " + rec.FullName.SurName + " " + rec.FullName.Otchestvo

		fullMatches := RaitaAll(x, []byte(fullname))
		descrMatches := RaitaAll(y, []byte(rec.Descrp))

		if len(fullMatches) > 0 && len(descrMatches) > 0 {
			result := fmt.Sprintf(
				"%-8d %-30s   %-55s \n", rec.Number, fullname, rec.Descrp,
			)
			output = append(output, result)
		}
	}

	elapsed := time.Since(start)

	output = append(output, fmt.Sprintf("Времени затрачено: %s", elapsed.String()))

	p.FillFile(output, "output/Raita.txt", len(output))
}
