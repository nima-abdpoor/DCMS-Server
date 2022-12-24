package util

import "strconv"

func MapStringArrayToHashArray(data []string, f func(string) uint32) []string {

	mapped := make([]string, len(data))

	for i, e := range data {
		mapped[i] = strconv.Itoa(int(f(e)))
	}

	return mapped
}
