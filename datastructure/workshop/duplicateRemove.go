package testproj

import (
	"container/list"
	"fmt"
)

func DuplicateRemove() {
	mylist := list.New()
	mylist.PushBack(InitEmployee("Jane", "Jones", 123))
	mylist.PushBack(InitEmployee("John", "Doe", 5678))
	mylist.PushBack(InitEmployee("Mike", "Wilson", 45))
	mylist.PushBack(InitEmployee("Mary", "Smith", 5555))
	mylist.PushBack(InitEmployee("John", "Doe", 5678))
	mylist.PushBack(InitEmployee("Bill", "End", 3948))
	mylist.PushBack(InitEmployee("Jane", "Jones", 123))
	printList(mylist)
	uniqueEmployeeId := make(map[int]struct{})

	for i := mylist.Front(); i != nil; {
		next := i.Next()
		val := i.Value.(*Employee)
		if _, ok := uniqueEmployeeId[val.id]; !ok {
			uniqueEmployeeId[val.id] = struct{}{}
		} else {
			mylist.Remove(i)
		}
		i = next
	}
	fmt.Println("After Removal:")
	printList(mylist)
}

func printList(mylist *list.List) {
	for i := mylist.Front(); i != nil; i = i.Next() {
		val := i.Value.(*Employee)
		fmt.Println(val.toString())
	}
}
