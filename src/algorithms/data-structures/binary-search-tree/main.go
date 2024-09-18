package main

import (
	"fmt"
)

var count int // Inicializando a variável count como int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Função Insert
func (n *Node) Insert(k int) {
	if n.Key < k {
		// Move para a direita
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// Move para a esquerda
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Função Search
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	count++ // Incrementa o contador
	if n.Key < k {
		// Move para a direita
		return n.Right.Search(k)
	} else if n.Key > k {
		// Move para a esquerda
		return n.Left.Search(k)
	}

	return true // Retorna true se a chave foi encontrada
}

// Função para imprimir a árvore binária
func (n *Node) PrintTree(prefix string, isLeft bool) {
	if n != nil {
		// Imprimir o nó atual
		fmt.Println(prefix + "+-- " + fmt.Sprintf("%d", n.Key))

		// Formatar para o próximo nível da árvore
		newPrefix := prefix
		if isLeft {
			newPrefix += "|   "
		} else {
			newPrefix += "    "
		}

		// Recursivamente imprimir os nós da esquerda e da direita
		n.Left.PrintTree(newPrefix, true)
		n.Right.PrintTree(newPrefix, false)
	}
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(450)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(350)
	tree.Insert(500)

	// Imprimir a árvore formatada
	fmt.Println("Árvore binária:")
	tree.PrintTree("", false)

	fmt.Println("\nBusca por 50:", tree.Search(50)) // Deve retornar true
	fmt.Println("Busca por 450:", tree.Search(450)) // Deve retornar true
	fmt.Println(
		"Contagem de operações:",
		count,
	) // Mostra o número de operações de busca realizadas
}
