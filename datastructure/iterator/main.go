package iterator

import "fmt"

func RunIterator() {
	item := NewConcreteCollection()
	item.Append(NewStudent("Tom", 19, "11"))
	item.Append(NewStudent("Jack", 21, "13"))
	item.Append(NewStudent("Jill", 18, "10"))
	item.Append(NewStudent("Jerry", 23, "14"))
	item.Append(NewStudent("Jim", 23, "25"))
	item.Append(NewStudent("Kate", 19, "11"))
	item.Append(NewStudent("Teddy", 24, "14"), NewStudent("Tim", 20, "11"))

	//create iterator
	it := item.CreateIterator()

	//Iterate using the iterator
	for it.HasNext() {
		data := it.Next()
		fmt.Println(data)
	}

	data := NewConcreteCollection()
	data.Append(48, 27, 49, 75, 57, 54, 57, 78, 7, 74, 91, 84, 62, 13, 41)

	//create iterator
	itr := data.CreateIterator()

	//Iterate using the iterator
	for itr.HasNext() {
		data := itr.Next()
		fmt.Println(data)
	}
}
