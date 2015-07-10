package sort


type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


func IsSorted(target Sortable) bool {
    for i := 0; i < target.Len() - 1; i++ {
        if target.Less(i + 1, i) {
            return false
        }
    }
    return true
}


type IntSlice []int


func (this *IntSlice) Len() int {
    return len(*this)
}

func (this *IntSlice) Less(i, j int) bool {
    return (*this)[i] < (*this)[j]
}

func (this *IntSlice) Swap(i, j int) {
    if i != j {
        (*this)[i], (*this)[j] = (*this)[j], (*this)[i]
    }
}
