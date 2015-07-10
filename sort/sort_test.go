package sort


import (
    "github.com/stretchr/testify/assert"
    "testing"
)


func TestSwap(t *testing.T) {
    var arr IntSlice = []int{1, 2}
    arr.Swap(0, 1)
    assert.Equal(t, arr[0], 2)
    assert.Equal(t, arr[1], 1)
}
