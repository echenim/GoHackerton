package main

import "fmt"

func main() {
	// list := &SingleList{}
	// list.append(&Node{data: "10", next: nil})
	// list.append(&Node{data: "20", next: nil})
	// list.append(&Node{data: "30", next: nil})
	// list.append(&Node{data: "40", next: nil})
	// list.append(&Node{data: "50", next: nil})
	// list.append(&Node{data: "60", next: nil})
	// list.append(&Node{data: "70", next: nil})
	// list.append(&Node{data: "80", next: nil})
	// //node := list.head
	// //list.Display()
	// Display(list.head, 1)
	// list.Delete("50")
	// list.Display()

	arr := []string{"10", "20", "30", "40", "50", "60", "70", "80", "90", "100", "110", "120"}
	list := &SingleList{}
	list.head = Append(arr)
	//list.AppendFromArray(arr)

	// list.Display()
	// //Display(list.head)
	// fmt.Printf("\nnumber of node in the linked list is : %v \n", list.size())
	// fmt.Print("-------------------- \n")
	// //list.DeleteNodeFromLinkedlist("60")
	// list.Display()
	// fmt.Printf("\nnumber of node in the linked list is : %v \n", list.size())
	list.Insert(5, "arrow")
	fmt.Print("-------------------- \n")
	list.Display()
	fmt.Printf("\nnumber of node in the linked list is : %v \n", list.size())

}
