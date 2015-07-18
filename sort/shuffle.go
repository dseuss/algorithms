package sort


import (
    "math/rand"
)


type Shuffleable interface {
    Swap(i, j int)
    Len() int
}


func Shuffle(target Shuffleable) {
    for i := 1; i < target.Len(); i++ {
        j := rand.Intn(i + 1)
        target.Swap(i, j)
    }
}
