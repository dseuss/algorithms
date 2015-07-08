package collections


import (
    "errors"
    "fmt"
)

const no_element_in_queue_error string = "No element in queue."


type DAQueue struct {
    head int
    tail int
    elem []int

}


func (this *DAQueue) Len() int {
    if len(this.elem) < 1 {
        return 0
    } else {
        return (this.tail - this.head) % (len(this.elem) + 1)
    }
}


func (this *DAQueue) Asize() int {
    return len(this.elem)
}


func (this *DAQueue) Queue(val int) (err error) {
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


func (this *DAQueue) Deque() (val int, err error) {
    if this.Len() <= 0 {
        return 0, errors.New(no_element_in_queue_error)
    }

    val = this.elem[this.head]
    this.elem[this.head] = 0    // Clean up for the GC
    this.head = (this.head + 1) % len(this.elem)

    if 4 * this.Len() <= len(this.elem) {
        err = this.resize(len(this.elem) / 2)
        if err != nil {
            return val, err
        }
    }

    return val, nil
}


func (this *DAQueue) resize(newsize int) error {
    if clen := this.Len(); clen > newsize {
        errstr := fmt.Sprintf("Resizing DAQueue failed: new size too small: %v < %v",
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
