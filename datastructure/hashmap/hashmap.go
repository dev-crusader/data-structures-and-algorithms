package hashmap

import "fmt"

const Size = 10

type HashMap struct {
	buckets [Size]*LinkListed
}

type LinkListed struct {
	head *ListNode
}

type ListNode struct {
	key   string
	value any
	next  *ListNode
}

// Insert function implements insertion of key to hashMap
func (h *HashMap) Insert(key string, value any) {
	index := generateHash(key)
	h.buckets[index].Insert(key, value)
}

func (h *HashMap) Delete(key string) {
	index := generateHash(key)
	h.buckets[index].Delete(key)
}

func (h *HashMap) Search(key string) bool {
	searchIndex := generateHash(key)
	return h.buckets[searchIndex].Search(key)
}

// Insert function insert the key into Linkedlist
func (l *LinkListed) Insert(key string, value any) {
	if !l.Search(key) {
		currentNode := &ListNode{key: key, value: value}
		currentNode.next = l.head
		l.head = currentNode
	} else {
		fmt.Printf("Insert Key: %s already exists\n", key)
	}
}

func (l *LinkListed) Delete(key string) {
	currentNode := l.head

	if currentNode.key == key {
		l.head = currentNode.next
		return
	}

	for currentNode.next != nil {
		if currentNode.next.key == key {
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
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func initMap() *HashMap {
	hm := &HashMap{}
	for k := range hm.buckets {
		hm.buckets[k] = &LinkListed{}
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
	m.Insert("Jack", 32)
	m.Insert("Pete", 44)
	m.Insert("Ryan", 76)
	m.Insert("Kyle", 675)
	m.Insert("Stuart", 354)
	m.Insert("John", "Parkers")
	m.Insert("Jim", "Kimmel")
	m.Insert("Bard", "AI")
	m.Insert("Kate", struct {
		age      int
		location string
	}{26, "Dallas"})
	fmt.Println(m.Search("Jack"))
	m.Delete("Bard")
	m.Insert("Jill", "Cody")
}
