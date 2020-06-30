package maps

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestMapIterationRandom(t *testing.T) {
    n := 10

    // create a map with values from 1 to 10
    m := make(map[int]bool)
    for i := 0; i < n; i++ {
        m[i] = true
    }

    // iterate over map and assign values to xs
    xs := make([]int, 0, n)
    for i := range m {
        xs = append(xs, i)
    }

    // iterate over map again and assign values to ys
    ys := make([]int, 0, n)
    for i := range m {
        ys = append(ys, i)
    }

    // xs is not equal to ys with high probability
    assert.NotEqual(t, xs, ys)
}