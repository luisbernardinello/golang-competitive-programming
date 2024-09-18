package main

import "fmt"

const ArraySize = 7

// Estrutura da tabela hash
type HashTable struct {
	array [ArraySize]*bucket
}

// Estrutura do bucket
type bucket struct {
	head *bucketNode
}

// Estrutura do bucketNode (nó de lista encadeada)
type bucketNode struct {
	key  string
	next *bucketNode
}

// Inserir na tabela hash
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Buscar na tabela hash
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Deletar na tabela hash
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Inserir em um bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// Buscar em um bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Deletar em um bucket
func (b *bucket) delete(k string) {
	if b.head == nil {
		fmt.Println("Key not found:", k)
		return
	}

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
	fmt.Println("Key not found:", k)
}

// Função hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Inicialização da tabela hash
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"RANDY",
		"LEONARD",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("LEONARD")

	fmt.Println("TOKEN found?", hashTable.Search("TOKEN"))
}
