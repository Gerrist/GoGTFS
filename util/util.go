package util

import (
	"fmt"
	"strconv"
	"strings"
)

func IndexOf(element string, data []string) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

func IndexOfInt(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

func ParseInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

func ParseFloat(str string) float64 {
	i, _ := strconv.ParseFloat(str, 32)

	return i
}

func ParseString(i int) string {
	s := strconv.Itoa(i)

	return s
}

func ParseFloatString(i float64) string {
	s := fmt.Sprintf("%g", i)

	return s
}

func CSVRow(data []interface{}) string { // custom CSV row generator as the go CSV library sucks

	row := make([]string, 0)

	for _, value := range data{
		switch v := value.(type) {
		case float64:
			row = append(row, ParseFloatString(v))
		case int:
			row = append(row, ParseString(v))
		default:
		case string:
			if strings.Contains(v, ",") {
				row = append(row, "\"" + v + "\"")
			} else {
				row = append(row, v)
			}
		}
	}

	return strings.Join(row, ",") + "\n"
}