package sort


func argmin(target Sortable, start int, stop int) (i_min int) {
    i_min = start
    for i := start + 1; i < stop; i++ {
        if target.Less(i, i_min) {
            i_min = i
        }
    }
    return i_min
}


func Selection(target Sortable) {
    for i := 0; i < target.Len(); i++ {
        i_min := argmin(target, i, target.Len())
        target.Swap(i, i_min)
    }
}


func Insertion(target Sortable) {
    for i := 1; i < target.Len(); i++ {
        j := i - 1
        for (j >= 0) && (target.Less(j+1, j)) {
            target.Swap(j, j+1)
            j--
        }
    }
}
