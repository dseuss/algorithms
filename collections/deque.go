package collections


import "errors"


type elem struct {
    val string
    next, prev *elem
}


type Deque struct {
    first, last *elem
}


func (this *Deque) IsEmpty() bool {
    return (this.first == nil) || (this.last == nil)
}


func (this *Deque) AddFirst(val string) {
    new_elem := elem{val, this.first, nil}
    this.first = &new_elem

    if this.first.next != nil {
        this.first.next.prev = &new_elem
    }

    if this.IsEmpty() {
        this.last = &new_elem
    }
}


func (this *Deque) AddLast(val string) {
    new_elem := elem{val, nil, this.last}
    this.last = &new_elem

    if this.last.prev != nil {
        this.last.prev.next = &new_elem
    }

    if this.IsEmpty() {
        this.first = &new_elem
    }
}


func (this *Deque) RmFirst() (string, error) {
    if this.IsEmpty() {
        return "", errors.New("Deque is empty")
    }

    val := this.first.val
    this.first = this.first.next
    return val, nil
}


func (this *Deque) RmLast() (string, error) {
    if this.IsEmpty() {
        return "", errors.New("Deque is empty")
    }

    val := this.last.val
    this.last = this.last.prev
    return val, nil
}
