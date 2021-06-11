# Algorithms Implemented in GO
## Intro
A personal learning record.  
All algorithms are Implemented in GO.

## Resources
[http://www.algorithmsilluminated.org/](http://www.algorithmsilluminated.org/)

## Implemented Algorithms
### Divide and Conquer
- Closest Pair of Points  
    `go run divideconquer\closestpair.go`  
    shows  
    `5.0990195135927845`

### Sorting
- QuickSort (RandomPivot + In-place Partition)  
    `go run sorting\quicksort.go`  
    shows  
    ```bash
    Testcase 1:
    25 comparisons (choose first as pivot)
    31 comparisons (choose last as pivot)
    24 comparisons (median of 3)
    22 comparisons (random pivot)

    Testcase 2:
    620 comparisons (choose first as pivot)
    573 comparisons (choose last as pivot)
    507 comparisons (median of 3)
    649 comparisons (random pivot)
    ```

### Random Selection
- Random Selection
    `go run randomselection/randomselection.go`  
    shows  
    ```bash
    5469
    4715
    ```
