package generics

import (
	"fmt"
)

/*
	This implementation features a generic HashMap designed to store key-value pairs using Go's generics,
	allowing for a flexible and type-safe apporach to hashmap creation. The HashMap supports any comparable
	key type and stores values of any (interface) type, making it versatile for various use cases.
	It uses hash function to compute the index for each key, facilitating efficient data retrieval by mapping
	keys to their corresponding array indices. Values are stored into the linkedlist.

	Basic Operations:
	1. Insert
	2. Retrieve
	3. Delete
*/

// HashValue represents a key-value pair.
type HashValue[K comparable, V any] struct {
	key   K
	value V
}

// Node represents a node in the linked list.
type Node[K comparable, V any] struct {
	hashValue *HashValue[K, V]
	next      *Node[K, V]
}

// LinkedList represents a linked list for handling collisions.
type LinkedList[K comparable, V any] struct {
	head *Node[K, V]
}

// Insert adds a new node to the linked list.
func (ll *LinkedList[K, V]) Insert(hashValue *HashValue[K, V]) {
	if _, found := ll.Find(hashValue.key); found {
		fmt.Printf("Insert Key: %v already exists\n", hashValue.key)
		return
	}
	newNode := &Node[K, V]{hashValue: hashValue}
	newNode.next = ll.head
	ll.head = newNode
}

// Find searches for a key in the linked list and returns its value if found.
func (ll *LinkedList[K, V]) Find(key K) (V, bool) {
	current := ll.head
	var zeroValue V // zero value of type V
	for current != nil {
		if current.hashValue.key == key {
			return current.hashValue.value, true
		}
		current = current.next
	}
	return zeroValue, false
}

// Delete deletes the node from the linkedlist if present
func (ll *LinkedList[K, V]) Delete(key K) bool {
	current := ll.head

	if current.hashValue.key == key {
		ll.head = current.next
		return true
	}
	for current.next != nil {
		if current.next.hashValue.key == key {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// HashTable represents the hash table.
type HashTable[K comparable, V any] struct {
	buckets []*LinkedList[K, V]
	size    int
}

// NewHashTable creates a new HashTable with a specified size.
func NewHashTable[K comparable, V any](size int) *HashTable[K, V] {
	buckets := make([]*LinkedList[K, V], size)
	for i := range buckets {
		buckets[i] = &LinkedList[K, V]{}
	}
	return &HashTable[K, V]{buckets: buckets, size: size}
}

// Hash function to compute the index for a given key.
func (ht *HashTable[K, V]) hash(key K) int {
	return int(fmt.Sprintf("%v", key)[0]) % ht.size
}

// Insert inserts a new key-value pair into the hash table.
func (ht *HashTable[K, V]) Insert(key K, value V) {
	hashValue := &HashValue[K, V]{key: key, value: value}
	index := ht.hash(key)
	ht.buckets[index].Insert(hashValue)
}

// Retrieve retrieves a value by key from the hash table and returns it if found.
func (ht *HashTable[K, V]) Retrieve(key K) (V, bool) {
	index := ht.hash(key)
	return ht.buckets[index].Find(key)
}

// Delete deletes a value by key from the hash table if present.
func (ht *HashTable[K, V]) Delete(key K) bool {
	index := ht.hash(key)
	return ht.buckets[index].Delete(key)
}

type Student struct {
	Age     int
	Id      string
	Name    string
	Address string
}

func NewStudent(age int, Id, name, addr string) *Student {
	return &Student{Id: Id, Age: age, Name: name, Address: addr}
}

func (s *Student) String() string {
	return fmt.Sprintf("Student{ Id: %s, Name: %s, Age: %d, Address: %s }", s.Id, s.Name, s.Age, s.Address)
}

// GenericMap function to demonstrate the generic hash table.
func GenericMap() {
	// Create a hash table for string keys and string values.
	stringHashTable := NewHashTable[string, string](10)
	stringHashTable.Insert("a", "hello")
	stringHashTable.Insert("b", "world")
	stringHashTable.Insert("a", "world")

	// Create a hash table for string keys and int values.
	intHashTable := NewHashTable[string, int](10)
	intHashTable.Insert("c", 42)
	intHashTable.Insert("d", 21)

	// Looking up keys in the string hash table.
	if value, found := stringHashTable.Retrieve("a"); found {
		fmt.Printf("Value for key 'a' in string hash table: %v\n", value)
	} else {
		fmt.Println("Key 'a' not found in string hash table.")
	}

	if value, found := stringHashTable.Retrieve("b"); found {
		fmt.Printf("Value for key 'b' in string hash table: %v\n", value)
	} else {
		fmt.Println("Key 'b' not found in string hash table.")
	}

	// Looking up keys in the int hash table.
	if value, found := intHashTable.Retrieve("c"); found {
		fmt.Printf("Value for key 'c' in int hash table: %v\n", value)
	} else {
		fmt.Println("Key 'c' not found in int hash table.")
	}

	if value, found := intHashTable.Retrieve("d"); found {
		fmt.Printf("Value for key 'd' in int hash table: %v\n", value)
	} else {
		fmt.Println("Key 'd' not found in int hash table.")
	}

	// Create a hash table for string keys(Id) and Student Info as values.
	infoHashTable := NewHashTable[string, *Student](10)
	infoHashTable.Insert("123", NewStudent(21, "123", "Jim", "Dallas"))
	infoHashTable.Insert("132", NewStudent(24, "132", "Railey", "Houston"))
	infoHashTable.Insert("543", NewStudent(22, "543", "Matt", "Florida"))
	infoHashTable.Insert("172", NewStudent(24, "172", "Bailey", "Colorado"))
	infoHashTable.Insert("123", NewStudent(21, "123", "Jill", "Seattle"))
	infoHashTable.Insert("863", NewStudent(21, "863", "Corey", "New York"))

	// Looking up keys in the student hash table.
	if value, found := infoHashTable.Retrieve("123"); found {
		fmt.Printf("Value for key '123' in student hash table: %v\n", value)
	} else {
		fmt.Println("Key '123' not found in student hash table.")
	}

	if value, found := infoHashTable.Retrieve("543"); found {
		fmt.Printf("Value for key '543' in student hash table: %v\n", value)
	} else {
		fmt.Println("Key '543' not found in student hash table.")
	}
	if value, found := infoHashTable.Retrieve("423"); found {
		fmt.Printf("Value for key '423' in student hash table: %v\n", value)
	} else {
		fmt.Println("Key '423' not found in student hash table.")
	}

	// Deleting the student info from the hash table
	fmt.Printf("Deleted: %t\n", infoHashTable.Delete("132"))
	fmt.Printf("Deleted: %t\n", infoHashTable.Delete("100"))
}
