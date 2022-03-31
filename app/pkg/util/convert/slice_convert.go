package convert

import "fmt"

func ToStringSlice(items []int) []string {
	var elems = make([]string, len(items))
	for index, value := range items {
		elems[index] = fmt.Sprintf("%d", value)
	}
	return elems
}
