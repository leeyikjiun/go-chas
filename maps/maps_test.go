package maps

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMapIterationRandom(t *testing.T) {
    n := 10

    m := make(map[int]bool)
    for i := 0; i < n; i++ {
        m[i] = true
    }

    xs := make([]int, 0, n)
    for i := range m {
        xs = append(xs, i)
    }

    ys := make([]int, 0, n)
    for i := range m {
        ys = append(ys, i)
    }

    assert.NotEqual(t, xs, ys)
}