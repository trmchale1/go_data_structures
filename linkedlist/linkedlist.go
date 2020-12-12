package linkedlist

import "fmt"

// create, destroy, add, remove, get largest, get smallest, print_list

func main () {
    linkedList := New()
    fmt.Println("Print basic line")
    linkedList.Append(1).Append(2).Append(3).Append(4).Append(5).PrintAll()
    fmt.Println("Delete 2")
    linkedList.DeleteWithValue(2).PrintAll()
}

type LinkedList struct {
    Head *Node
    Tail *Node
}

type Node struct {
    Next *Node
    Data interface{}
}

func New() *LinkedList {
    emptyNode := &Node{
        Next: nil,
        Data: nil,
    }

    return &LinkedList{
        Head: emptyNode,
        Tail: emptyNode,
    }
}

func (ll *LinkedList) IsEmpty() bool {
        if ll.Head == nil {
            return true
        }
        return false
}

func (ll *LinkedList) Append(d interface{}) *LinkedList {
    nextNode := &Node{
        Next: nil,
        Data: d,
    }
    if ll.Head.Data == nil {
        ll.Head = nextNode
    } else {
        ll.Tail.Next = nextNode
    }
    ll.Tail = nextNode
    return ll
}

func (ll *LinkedList) DeleteWithValue(v interface{}) *LinkedList {
    var node = ll.Head
    if node.Data == v {
        ll.Head = ll.Head.Next
        return ll
    }
    for {
        if v == node.Next.Data {
            if node.Next.Next != nil {
                node.Next = node.Next.Next
                break
                }
            node.Next = nil
            break
            }
        node = node.Next
        }
    return ll
}

func (ll *LinkedList) PrintAll() {
    var node = ll.Head
    for {
        fmt.Println(node.Data)
        if node.Next == nil {
            return
        }
        node = node.Next
    }
}
