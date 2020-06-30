# Assignment

### Short assignment does not overwrite variable from outer scope
```
i := 0
{
    // a new variable i is declared here
    i, j := 1, 2
}
// i is still 0 here
```

Be careful if you are doing this in a defer function of a named return value function
```
// incorrect
func f() (i int) {
    defer func() {
        // not the same i as the named return value
        i, j := 1, 2

        // named return value is still 0
    }()

    return 0
}
```
```
// correct
func f() (i int) {
    defer func() {
        var j int
        // overwrites the i in the named return value
        i, j = 1, 2

        // named return value is now 1
    }()

    return 0
}
```
