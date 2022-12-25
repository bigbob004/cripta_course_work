package pkg

import (
	"math/rand"
	"time"
)

// функция отдаёт массив-поднмножество строк res, размера size, который был сформирован из массива-множества строк strs
func UniqueSubSet(strs []string, size int) []string {
	rand.Seed(time.Now().UnixNano())
	set := make(map[string]struct{}, size)

	for len(set) != size {
		index := GenerateRandomNumInRange(0, len(strs)-1)
		set[strs[index]] = struct{}{}
	}

	return ConvertToSlice(set)
}

func ConvertToSlice(set map[string]struct{}) []string {
	res := make([]string, 0, len(set))

	for key, _ := range set {
		res = append(res, key)
	}

	return res
}

func GenerateRandomNumInRange(a, b int) int {
	return a + rand.Intn(b-a+1) // a ≤ n ≤ b
}

func ShuffleSlice(slice []string) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func IsEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
