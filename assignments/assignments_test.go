package assignments

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestShortAssignmentsDifferentScopes(t *testing.T) {
    // declare a variable i and assign it to 0
    i := 0
    pi := &i

    i, j := 1, 2
    assert.Equal(t, 2, j)

    // still the same i and its value is overwritten to 1
    assert.Same(t, pi, &i)
    assert.Equal(t, 1, i)

    {
        i, j := 3, 4
        assert.Equal(t, 4, j)

        // a new variable i is declared in this scope and assigned to 3
        assert.NotSame(t, pi, &i)
        assert.Equal(t, 3, i)
    }

    // the variable i in this scope is still assigned to 1
    assert.Equal(t, 1, i)

    {
        var j int
        i, j = 5, 6
        assert.Equal(t, 6, j)

        // no new variable i is re-declared in this scope and its value is overwritten to 5
        assert.Same(t, pi, &i)
        assert.Equal(t, 5, i)
    }

    // the variable i is re-assigned to 5 in the assignment in the previous scope
    assert.Equal(t, 5, i)
}
