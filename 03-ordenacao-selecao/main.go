package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 3, 6, 2, 10, 255, 1, 20, 54, 205, 79, 85}
	sortedArr := ordenacaoPorSelecao(arr)
	fmt.Println(sortedArr)
}

func buscaMenor(arr []int) int {
	menor := arr[0]
	menorIndice := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
			menorIndice = i
		}
	}
	return menorIndice
}

func ordenacaoPorSelecao(arr []int) []int {
	var novoArr []int
	for len(arr) > 0 {
		menor := buscaMenor(arr)
		novoArr = append(novoArr, arr[menor])
		arr = append(arr[:menor], arr[menor+1:]...)
	}
	return novoArr
}
