package main

import "fmt"

func groupTemps(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, t := range temps {
		key := int(t/10) * 10
		groups[key] = append(groups[key], t)
	}

	return groups
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(groupTemps(temps))
}
