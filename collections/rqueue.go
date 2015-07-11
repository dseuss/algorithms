package collections


import (
    "errors"
    "math/rand"
    "github.com/dseuss/algorithms/sort"
    "fmt"
)


type RQueue struct {
    DAQueue
}


func (this *RQueue) Deque() (val int, err error) {
    index, err := this.sampleIndex()
    if err != nil {
        return 0, err
    }

    val = this.elem[index]
    this.elem[index] = this.elem[this.head]
    this.head = (this.head + 1) % len(this.elem)

    if 4 * this.Len() <= len(this.elem) {
        err = this.resize(len(this.elem) / 2)
        if err != nil {
            return val, err
        }
    }

    return val, nil
}


func (this* RQueue) sampleIndex() (int, error) {
    current_len := this.Len()
    if current_len < 0 {
        return 0, errors.New(no_element_in_queue_error)
    }

    if current_len == 1 {
        return this.head, nil
    }

    i := rand.Intn(current_len - 1)
    return (i + this.head) % len(this.elem), nil
}


// TODO This is not adaptive against change
func (this* RQueue) Iter()  <-chan int {
    ch := make(chan int)
    var indices sort.IntSlice = make([]int, this.Len())
    for i := range indices {
        indices[i] = i
    }
    sort.Shuffle(&indices)

    go func() {
        fmt.Println(indices)
        for _, index := range indices {
            ch <- this.elem[index]
        }
        close(ch)
    } ();

    return ch
}
