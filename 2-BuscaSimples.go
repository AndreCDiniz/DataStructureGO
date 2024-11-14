package main

import "fmt"

// Na Busca simples seria O(n) por ser constante.
func main() {

	arr := []int{1, 3, 5, 7, 8, 9, 11, 15, 19, 20, 21, 25}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
