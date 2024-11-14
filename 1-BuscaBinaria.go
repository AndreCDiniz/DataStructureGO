package main

import "fmt"

// Na Busca Binária a anotação para ela seria O(log n)
func main() {

	arr := []int{1, 3, 5, 7, 8, 9, 11, 15, 19, 20, 21, 25}
	alvo := 3

	resultado := buscaBinaria(arr, alvo)

	if resultado != -1 {
		fmt.Printf("Elemento %d encontardo no índice %d\n", alvo, resultado)
	} else {
		fmt.Printf("Elemento %d não encontrado \n", alvo)
	}

}

func buscaBinaria(arr []int, alvo int) int {

	esquerda := 0
	direita := len(arr) - 1

	for esquerda <= direita {
		meio := esquerda + (direita-esquerda)/2

		if arr[meio] == alvo {
			return meio
		}

		if arr[meio] < alvo {
			esquerda = meio + 1
		} else {
			direita = meio - 1
		}
	}
	return -1
}
