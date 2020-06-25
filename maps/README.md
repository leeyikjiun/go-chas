# Map

### Map iteration is random
```
Input: [0 1 2 3 4 5 6 7 8 9]
First Output: [1 5 4 6 7 8 9 0 2 3]
Second Output: [1 5 8 9 0 2 3 4 6 7]
```
Source: https://blog.golang.org/maps
> Iteration order
>
> **When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next.** If you require a stable iteration order you must maintain a separate data structure that specifies that order. This example uses a separate sorted slice of keys to print a map[int]string in key order:
