# Randomized Selection
## The Problem
Input:  
array $A$ with $n$ distinct numbers and a number $i \in\{1, 2, ..., n\}$

Output:  
$i^{th}$ order statistic  
(i.e. $i^{th}$ smallest element of input array)

Example:  
median  
($i=\frac{n+1}{2}$ for $n$ odd, $i=\frac{n}{2}$ for $n$ even)

![](https://i.imgur.com/ZCQVQPy.png)

## Randomized Selection
~~~
RSelect(array A, length n, order statistic i):
    (1) if n=1 return A[1]
    (2) choose pivot p from A uniformly at random
    (3) partition A around p
        let j=(new index of p)
    (4) if j=i return p
    (5) if j>i return RSelect(1st part of A, j-1, i)
    (6) if j<i return RSelect(2nd part of A, n-j, i-j)

~~~
only recur once not twice!!

# Deterministic Selection Algorithm (Optional)
## Guaranteeing a Good Pivot
Recall:  
"best" pivot = the median! (seems circular!)

Goal:  
find pivot guaranteed to be pretty good

Key Idea:  
use "median of medians"!

## A Deterministic ChoosePivot
~~~
ChoosePivot(A, n)
    (1) logically break A into n/5 groups of size 5 each
    (2) sort each group (e.g. using MergeSort)
    (3) copy n/5 medians (i.e. middle element of each sorted group)
    (4) recursively compute median of C (!)
    (5) return this as pivot
~~~

## The DSelect Algorithm
![](https://i.imgur.com/Xtj4Bel.png)

**(有錯誤$j<i$跟$j>i$應該反過來)**

groups of 5: 五個為一組
