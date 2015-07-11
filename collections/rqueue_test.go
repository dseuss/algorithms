package collections


import (
    "github.com/stretchr/testify/assert"
    "testing"
)


func all(arr []bool) bool {
    for _, val := range arr {
        if !val {
            return false
        }
    }
    return true
}


func TestRQueueIterator(t *testing.T) {
    const entries = 10
    var found []bool = make([]bool, entries)

    var rq RQueue
    for i := 0; i < entries; i++ {
        rq.Queue(i)
    }

    for j := range rq.Iter() {
        assert.False(t, found[j])
        found[j] = true
    }

    assert.True(t, all(found))
    assert.Equal(t, entries, rq.Len())
}


func TestRQueueCompleteness(t *testing.T) {
    const entries = 10
    var found []bool = make([]bool, entries)

    var rq RQueue
    for i := 0; i < entries; i++ {
        rq.Queue(i)
    }

    for i := 0; i < entries; i++ {
        j, err := rq.Deque()
        assert.False(t, found[j])
        assert.Nil(t, err)
        found[j] = true
    }

    assert.True(t, all(found))
    assert.Equal(t, 0, rq.Len())
}
