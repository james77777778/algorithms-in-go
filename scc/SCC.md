# Kosaraju's Two-Pass Algorithm
Theorem:  
can compute in SCCs in $O(m+n)$ time

Algorithm: (given directed graph G)
1. Let $G^{rev} = G$ with all arcs reversed
2. run DFS-Loop on $G^{rev}$  
    goal: compute "magical ordering" of nodes  
    Let $f(v)$ = "finishing time" of each $v \in V$
3. run DFS-Loop on $G$  
    goal: discover the SCCs one-by-one  
    processing nodes in decreasing order of finishing times  
    (SCCs = nodes with the same "leader")
    
## DFS-Loop
```bash
DFS-Loop(graph G)
    global variable t = 0 (for finishing times in 1st pass)
    (number of nodes processed so far)
    global variable s = NULL (for leaders in 2nd pass)
    (current source vertex)
    Assume nodes labelled 1 to n
    for i=n down to 1:
        if i not yet explored
            s = i
            DFS(G, i)
```
```bash
DFS(graph G, node i)
    mark i as explored (for rest of DFS-Loop)
    set leader(i) = node s
    for each arc (i, j) in G:
        if j not yet explored:
            DFS(G, j)
    t++
    set f(i) = t (i's finishing time)
```

## Quiz
![](https://i.imgur.com/oog1nxa.png)

因為是從$G^{rev}$開始，所以會從node 9開始  
9->6->3->9  
(設f(3)=1)  
6->8->2->5->8  
(設f(5)=2, 設f(2)=3, 設f(8)=4, 設f(6)=5, 設f(9)=6)  
7->4->1->7  
(設f(1)=7, f(4)=8, f(7)=9)

## Example (2nd pass)
![](https://i.imgur.com/SfquRen.png)

將++magical ordering++取代原本的node name  
(1, 2, 3, 4, 5, 6, 7, 8, 9) =>  
(7, 3, 1, 8, 2, 5, 9, 4, 6)

![](https://i.imgur.com/jk3YtEa.png)

從finishing time最高的開始(9)  
9->7->8->9 (找到SCC, leader=9 (原本為7))  
8, 7都找過了  
6->1->5->6 (找到SCC, leader=6 (原本為9))  
5找過了  
4->2->3->4 (找到SCC, leader=4 (原本為8))  
3, 2, 1都找過了  
結束演算法

Running Time:  
2*DFS = $O(m+n)$
