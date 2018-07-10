package main

import "fmt"

type Node struct {
	key interface{}
	next *Node
}

type List struct{
	head *Node
}

func (L *List) insertKey(key interface{}){
	
	list := &Node{
		key:	key,
		next: 	nil,
	}

	if L.head == nil {
		L.head = list
	} else {
		temp := L.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = list
	}
}

func (l *List) printList(){
	list := l.head
	fmt.Printf("List is:")
	for (list != nil) {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("Implement linked list")
	link := List{}
	var n, j int
	fmt.Println("How many numbers?")
	fmt.Scanln(&n)
	for i:=0; i<n; i++ {
		fmt.Println("\nEnter a number ")
		fmt.Scanln(&j)
		link.insertKey(j)
		link.printList()
	}

	//temp := Node{}
	//fmt.Println(temp)
}