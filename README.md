# Algorithms Implemented in GO
## Intro
A personal learning record.  
All algorithms are Implemented in GO.

## Resources & Test Cases
[http://www.algorithmsilluminated.org/](http://www.algorithmsilluminated.org/)

## Implemented Algorithms
### Divide and Conquer
- Closest Pair of Points  
    `go run cmd/closest_pair/closest_pair.go`  
    shows  
    `5.0990195135927845`

### Sorting
- QuickSort (RandomPivot + In-place Partition)  
    `go run cmd/sort/quick_sort.go`  
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
    `go run cmd/random_selection/random_selection.go`  
    shows  
    ```bash
    5469
    4715
    ```
- Deterministic Selection  
    `go run cmd/deterministic_selection/deterministic_selection.go`  
    shows  
    ```bash
    5469
    4715
    ```

### Strong Connected Components
- Strong Connected Components  
    `go run cmd/scc/scc.go`  
    shows top 5 SCC size  
    ```bash
    [3 3 3 0 0]
    [3 3 2 0 0]
    [3 3 1 1 0]
    [7 1 0 0 0]
    [6 3 2 1 0]
    ```

### Heap
- Heap Operations  
    `go run cmd/test_heap/test_heap.go`  
    shows a sorted order  
    ```bash
    1
    3
    4
    5
    7
    10
    100
    ```

### Dijkstra Shortest Path
- Dijkstra Shortest Path
    `go run cmd/dijkstra_shortest_path/dijkstra_shortest_path.go`  
    shows the shortest-path distances from vertex 1 to every other vertex  
    ```bash
    vertices: 1, shortest path: 0
    vertices: 2, shortest path: 1
    vertices: 3, shortest path: 2
    vertices: 4, shortest path: 3
    vertices: 5, shortest path: 4
    vertices: 6, shortest path: 4
    vertices: 7, shortest path: 3
    vertices: 8, shortest path: 2
    ```
