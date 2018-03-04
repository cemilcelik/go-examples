package main

import (
	"fmt"
	"strings"
)

func Index(source []string, target string) int {
	for k, v := range source {
		if v == target {
			return k
		}
	}
	return -1
}

func Exists(source []string, target string) bool {
	if Index(source, target) >= 0 {
		return true
	}
	return false
}

func Any(source []string, callback func(target string) bool) bool {
	for _, v := range source {
		if callback(v) {
			return true
		}
	}
	return false
}

func All(source []string, callback func(target string) bool) bool {
	for _, v := range source {
		if !callback(v) {
			return false
		}
	}
	return true
}

func Filter(source []string, callback func(target string) bool) []string {
	results := make([]string, 0)
	for _, v := range source {
		if callback(v) {
			results = append(results, v)
		}
	}
	return results
}

func Map(source []string, callback func(target string) string) []string {
	results := make([]string, len(source))
	for k, v := range source {
		results[k] = callback(v)
	}
	return results
}

func main() {
	fruits := []string{"peach", "orange", "banana", "kiwi", "apple"}
	fmt.Printf("fruits: %v, type: %T\n", fruits, fruits)

	// Index()
	fmt.Printf("index of 'kiwi': %d\n", Index(fruits, "kiwi"))

	// Exists()
	fmt.Printf("existance of 'kiwi': %v\n", Exists(fruits, "kiwi"))

	// Any()
	fmt.Printf("any starts with 'ki': %v\n", Any(fruits, func(target string) bool {
		return strings.HasPrefix(target, "ki")
	}))

	// All()
	fmt.Printf("all ends with 'wi': %v\n", All(fruits, func(target string) bool {
		return strings.HasSuffix(target, "wi")
	}))

	// Filter()
	fmt.Printf("filter contains with 'an': %v\n", Filter(fruits, func(target string) bool {
		return strings.Contains(target, "an")
	}))

	// Map()
	fmt.Printf("map uppercase: %v\n", Map(fruits, func(target string) string {
		return strings.ToUpper(target)
	}))
}
