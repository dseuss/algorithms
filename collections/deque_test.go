package collections


import (
    "github.com/stretchr/testify/assert"
    "testing"
    "strings"
)


func TestFifo(t *testing.T) {
    const input = "A B C"
    const output = "A B C"
    var dq Deque

    for _, word := range strings.Split(input, " ") {
        dq.AddFirst(word)
    }

    for _, word := range strings.Split(output, " ") {
        res, err := dq.RmLast()
        assert.Equal(t, word, res)
        assert.Nil(t, err)
    }

    assert.True(t, dq.IsEmpty())
    _, err := dq.RmLast()
    assert.NotNil(t, err)
}


func TestFifoReversed(t *testing.T) {
    const input = "A B C"
    const output = "A B C"
    var dq Deque

    for _, word := range strings.Split(input, " ") {
        dq.AddLast(word)
    }

    for _, word := range strings.Split(output, " ") {
        res, err := dq.RmFirst()
        assert.Equal(t, word, res)
        assert.Nil(t, err)
    }

    assert.True(t, dq.IsEmpty())
    _, err := dq.RmLast()
    assert.NotNil(t, err)
}


func TestFilo(t *testing.T) {
    const input = "A B C"
    const output = "C B A"
    var dq Deque

    for _, word := range strings.Split(input, " ") {
        dq.AddFirst(word)
    }

    for _, word := range strings.Split(output, " ") {
        res, err := dq.RmFirst()
        assert.Equal(t, word, res)
        assert.Nil(t, err)
    }

    assert.True(t, dq.IsEmpty())
    _, err := dq.RmLast()
    assert.NotNil(t, err)
}
