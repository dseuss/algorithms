package sort


type Sortable interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}


type IntSlice []int


func (this *IntSlice) Len() int {
    return len(*this)
}

func (this *IntSlice) Less(i, j int) bool {
    return (*this)[i] < (*this)[j]
}

func (this *IntSlice) Swap(i, j int) {
    (*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}
