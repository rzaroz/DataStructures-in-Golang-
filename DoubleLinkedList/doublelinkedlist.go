package doublelinklist

import (
	"fmt"
	"strconv"
)

type Node struct {
	data   int
	before *Node
	after  *Node
}

type DoubleLinkedList struct {
	root *Node
}

func (list *DoubleLinkedList) Append(data int) {
	newNode := Node{data: data, before: nil, after: nil}
	if list.root != nil {
		if list.root.after == nil {
			list.root.after = &newNode
			list.root.before = &newNode
			newNode.before = list.root
		} else {
			current := list.root.after
			for {
				if current.after == nil {
					current.after = &newNode
					newNode.before = current
					break
				} else {
					current = current.after
				}
			}
		}
	} else {
		list.root = &newNode
	}
}

func (list *DoubleLinkedList) PrintList() string {
	output := "["
	current := list.root
	for {
		output += strconv.Itoa(current.data)
		if current.after != nil {
			output += ", "
			current = current.after
		} else {
			output += "]"
			break
		}
	}

	return output
}

func (list *DoubleLinkedList) Drop(index int) (bool, error) {
	counter := 0
	current := list.root
	for {
		if counter == index {
			current.before.after = current.after
			current.after.before = current.before
			current = nil
			break
		} else {
			counter++
			current = current.after
		}
		if current == nil {
			return false, fmt.Errorf("there is no index like that")
		}
	}

	return true, nil
}

func (list *DoubleLinkedList) get(index int) (int, error) {
	counter := 0
	current := list.root
	for {
		if counter == index {
			return current.data, nil
		} else {
			counter++
			current = current.after
		}
		if current == nil {
			return 0, fmt.Errorf("there is no index like that")
		}
	}
}

func main() {
	Root := Node{1, nil, nil}
	list := DoubleLinkedList{root: &Root}
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)
	list.Append(9)
	list.Append(10)

	ListPrint := list.PrintList()
	fmt.Printf("list : %v \n", ListPrint)

	_, err := list.Drop(3)
	if err != nil {
		fmt.Println(err)
	}
	AfterDrop := list.PrintList()

	fmt.Printf("list after drop: %v \n", AfterDrop)

	number, err := list.get(4)
	fmt.Println(number)
}
