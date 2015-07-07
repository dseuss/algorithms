package collections


import (
    "errors"
    "fmt"
)


type Queue struct {
    head int
    tail int
    elem []int
}


func (this *Queue) Len() int {
    if len(this.elem) < 1 {
        return 0
    } else {
        return (this.tail - this.head) % (len(this.elem) + 1)
    }
}


func (this *Queue) Asize() int {
    return len(this.elem)
}


func (this *Queue) Queue(val int) (err error) {
    if len(this.elem) < 1 {
        err = this.resize(1)
        if err != nil {
            return err
        }
    }

    if (this.Len() + 1) > len(this.elem) {
        err = this.resize(2 * len(this.elem))
        if err != nil {
            return err
        }
    }

    this.elem[this.tail % len(this.elem)] = val
    this.tail = (this.tail + 1) % (len(this.elem) + 1)
    return nil
}


func (this *Queue) Deque() (val int, err error) {
    val = this.elem[this.head]
    this.head = (this.head + 1) % len(this.elem)

    if 4 * this.Len() <= len(this.elem) {
        err = this.resize(len(this.elem) / 2)
        if err != nil {
            return val, err
        }
    }

    return val, nil
}


func (this *Queue) resize(newsize int) error {
    if clen := this.Len(); clen > newsize {
        errstr := fmt.Sprintf("Resizing Queue failed: new size too small: %v < %v",
                              clen, newsize)
        return errors.New(errstr)
    }

    tmp := make([]int, newsize)
    for i := 0; i < this.Len(); i++ {
        tmp[i] = this.elem[(this.head + i) % len(this.elem)]
    }
    this.elem = tmp
    this.tail = this.Len()
    this.head = 0

    return nil
}
