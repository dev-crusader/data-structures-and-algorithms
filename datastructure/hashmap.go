package datastructure

import "fmt"

const Size = 10

type HashMap struct {
	ArrayMap [Size]*LinkListed
}

type LinkListed struct {
	head *ListNode
}

type ListNode struct {
	data string
	next *ListNode
}

// Insert function implements insertion of key to hashMap
func (h *HashMap) Insert(key string) {
	index := generateHash(key)
	h.ArrayMap[index].Insert(key)
}

func (h *HashMap) Delete(key string) {
	index := generateHash(key)
	h.ArrayMap[index].Delete(key)
}

func (h *HashMap) Search(key string) bool {
	searchIndex := generateHash(key)
	return h.ArrayMap[searchIndex].Search(key)
}

// Insert function insert the key into Linkedlist
func (l *LinkListed) Insert(key string) {
	if !l.Search(key) {
		currentNode := &ListNode{data: key}
		currentNode.next = l.head
		l.head = currentNode
	} else {
		fmt.Printf("Insert Key: %s already exists\n", key)
	}
}

func (l *LinkListed) Delete(key string) {
	currentNode := l.head

	if currentNode.data == key {
		l.head = currentNode.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.data == key {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
	fmt.Printf("Delete Key: %s does not exists\n", key)
}

func (l *LinkListed) Search(key string) bool {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.data == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func initMap() *HashMap {
	hm := &HashMap{}
	for k := range hm.ArrayMap {
		hm.ArrayMap[k] = &LinkListed{}
	}
	return hm
}

func generateHash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % len(key)
}

func InitHashMap() {
	m := initMap()
	fmt.Println(m)
	m.Insert("Jack")
	m.Insert("Pete")
	m.Insert("Ryan")
	m.Insert("Kyle")
	m.Insert("Stuart")
	m.Insert("John")
	m.Insert("Jim")
	m.Insert("Bard")
	m.Insert("Kate")
	fmt.Println(m.Search("Jack"))
	m.Delete("Kate")
	m.Insert("Jill")
}
