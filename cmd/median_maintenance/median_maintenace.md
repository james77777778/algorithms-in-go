### Application: Median Maintanence
I give you:
a sequence $x_1, ..., x_n$ of numbers, one-by-one

You tell me:
at each time step $i$, the median of $x_1, ..., x_n$

Constraint:
use $O(logi)$ time at each step $i$

Solution:
maintain heaps $H_{low}$: supports Extract-Max
maintain heaps $H_{high}$: supports Extract-Min

Key idea:
maintain invariant that $i/2$ smallest (largest) elements in $H_{low}$ ($H_{high}$)

假設有20個數字
最低的10個放在H_low，最大的10個放在H_high
當有新的數字進來，大於H_low的最大值則放在H_high，小於H_high的最小值則放在H_low，若兩者皆達成則隨意放，並成為新的中位數

要注意不平衡狀況，可以將多的H_high或H_low取出來並放到另一處，維持平衡狀態
