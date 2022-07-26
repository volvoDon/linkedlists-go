package linkedlists

import (
	//"err"
	//"log"
	"fmt"

	
)


type Node struct {
	prev *Node
	next *Node
	key  string
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key string) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func PopT(l *List) Node {
	res := l.tail
	if l.tail.prev != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
		res.next = nil
		res.prev = nil
		} else {
			fmt.Println("This List only contains one node")
		}
	return *res		
} 	

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

//func main()  {
//list := List{}
//list.Insert("hello")

//popped := PopT(&list)
//fmt.Println(popped.key)

//}