package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/saiias/goml/array"
)

func Svmlight(p string) (array.Array, []array.Array) {

	label := make(array.Array, 0)
	matrix := make([]array.Array, 0)

	fp, err := os.Open(p)
	defer fp.Close()

	if err != nil {
		fmt.Println(err)
	}
	sc := bufio.NewScanner(fp)
	linum := 0
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(line, "#") {
			continue
		}

		vec := strings.Split(line, " ")
		if vec[0] == "+1" {
			label[linum] = 1.0
		} else {
			label[linum] = -1.0
		}

		fv := make(array.Array, 100)
		for i, feat := range vec[1:] {

			t := strings.Split(feat, ":")
			if len(t) == 1 {
				continue
			} else if len(t) != 2 {
				fmt.Printf("invalid format: line %v index %v\n", linum, i)
				os.Exit(1)
			}

			fn, err := strconv.Atoi(t[0])
			if err != nil {
				fmt.Printf("invalid format: line %v index %v\n", linum, i)
				os.Exit(1)
			}

			value, err := strconv.ParseFloat(t[1], 64)
			if err != nil {
				fmt.Printf("invalid format: line %v index %v\n", linum, i)
				os.Exit(1)
			}

			fv[fn] = value
		}
		matrix = append(matrix, fv)
		linum += 1
	}
	return label, matrix

}
