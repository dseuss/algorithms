package collections


import (
    "github.com/stretchr/testify/assert"
    "testing"
)


func TestResizing(t *testing.T) {
    var qu Queue
    const n_max = 10

    for n := uint(0); n <= n_max; n++ {
        // FIXME This is dangerous since 1 << (n - 1) for n=0 depends on
        // overflow
        for j := 1 << (n - 1); j < (1 << n); j++ {
            err := qu.Queue(j)
            if err != nil {
                panic(err.Error())
            }
            assert.Equal(t, qu.Len(), j + 1)
        }
        // The allocated size should be the smallest n, such that j < 2^n
        assert.Equal(t, qu.Asize(), 1 << n)
    }

    for n := uint(n_max); n > 0; n-- {
        for j := 1 << n; j > 1 << (n - 1); j-- {
            qu.Deque()
        }
        assert.True(t, qu.Asize() < 4*(1 << n) )
    }

}


func TestCorrectness(t *testing.T) {
    var testarr = []int {1, 2, 3, 4}
    var qu Queue

    testQueue := func (arr []int) {
        for _, val := range arr {
            qu.Queue(val)
        }
    }

    testDeque := func (arr []int) {
        for _, word := range arr {
            res, err := qu.Deque()
            assert.Equal(t, word, res)
            assert.Nil(t, err)
        }
    }

    testQueue(testarr)
    testDeque(testarr[:3])
    assert.Equal(t, qu.Len(), 1)

    testQueue(testarr)

    val, err := qu.Deque()
    assert.Equal(t, val, 4)
    assert.Nil(t, err)

    testDeque(testarr[:3])
}
