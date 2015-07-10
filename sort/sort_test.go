package sort


import (
    "github.com/stretchr/testify/assert"
    "testing"

    "math/rand"
    "fmt"
)

var sorting_functions []func(Sortable) = []func (Sortable) { Selection }
var sorting_functions_desc []string = []string { "Selection sort" }


func random_intslice(length int) (result IntSlice) {
    result = make([]int, length)
    for n := range result {
        result[n] = rand.Intn(length)
    }
    return result
}


func copy_intslice(source IntSlice)(result IntSlice) {
    result = make([]int, len(source))
    copy(result, source)
    return result
}

func TestSwap(t *testing.T) {
    var arr IntSlice = []int{1, 2}
    arr.Swap(0, 1)
    assert.Equal(t, 2, arr[0])
    assert.Equal(t, 1, arr[1])
}


func TestIsSorted(t *testing.T) {
    var sorted_arr IntSlice = []int{1, 2, 2, 3}
    var unsorted_arr IntSlice = []int{2, 1, 3, 1}

    assert.True(t, IsSorted(&sorted_arr))
    assert.False(t, IsSorted(&unsorted_arr))
}


func TestSorting(t *testing.T) {
    instance_lengths := []int {5, 5, 10}

    for instance_nr, length := range instance_lengths {
        for index, sort := range sorting_functions {
            instance := random_intslice(length)
            instance_orig := copy_intslice(instance)

            sort(&instance)
            if ! IsSorted(&instance) {
                t.Log(sorting_functions_desc[index], "failed on instance", instance_nr)
                t.Log("    Got ", fmt.Sprintf("%v", instance))
                t.Log("    From", fmt.Sprintf("%v", instance_orig))
                t.Fail()
            }
        }
    }
}


func TestArgmin(t *testing.T) {
    var testarr IntSlice = []int {45, 2, 5, 0}
    assert.Equal(t, 3, argmin(&testarr, 0, testarr.Len()))
}
