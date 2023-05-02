package main

import (
	"fmt"
	"math"
	"strings"
)

var asciimap = map[int64]string{0: "a",
	1:  "b",
	2:  "c",
	3:  "d",
	4:  "e",
	5:  "f",
	6:  "g",
	7:  "h",
	8:  "i",
	9:  "j",
	10: "k",
	11: "l",
	12: "m",
	13: "n",
	14: "o",
	15: "p",
	16: "q",
	17: "r",
	18: "s",
	19: "t",
	20: "u",
	21: "v",
	22: "w",
	23: "x",
	24: "y",
	25: "z",
	26: "A",
	27: "B",
	28: "C",
	29: "D",
	30: "E",
	31: "F",
	32: "G",
	33: "H",
	34: "I",
	35: "J",
	36: "K",
	37: "L",
	38: "M",
	39: "N",
	40: "O",
	41: "P",
	42: "Q",
	43: "R",
	44: "S",
	45: "T",
	46: "U",
	47: "V",
	48: "W",
	49: "X",
	50: "Y",
	51: "Z",
	52: "0",
	53: "1",
	54: "2",
	55: "3",
	56: "4",
	57: "5",
	58: "6",
	59: "7",
	60: "8",
	61: "9"}

func idToBase62(id int64) string {
	var coefs []int64
	num := id
	for num > 0 {
		remainder := num % 62
		coefs = append(coefs, remainder)
		num = num / 62
	}
	reverse(coefs)
	fmt.Println(coefs)
	var output string
	for _, coef := range coefs {
		output = output + asciimap[coef]
	}
	return output
}

func base62ToId(base62 string) int64 {
	charsList := strings.Split(base62, "")

	var id int64
	for i, ch := range charsList {
		foundKey, _ := findKey(asciimap, ch)
		id += int64(float64(foundKey) * math.Pow(float64(62), float64((len(charsList)-1)-i)))
	}

	return id
}

func findKey(m map[int64]string, value string) (key int64, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
