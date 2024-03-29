# Question

This problem was asked by Apple.

Suppose you have a multiplication table that is N by N. That is, a 2D array where the value at the i-th row and j-th column is (i + 1) _ (j + 1) (if 0-indexed) or i _ j (if 1-indexed).

Given integers N and X, write a function that returns the number of times X appears as a value in an N by N multiplication table.

For example, given N = 6 and X = 12, you should return 4, since the multiplication table looks like this:

| 1 | 2 | 3 | 4 | 5 | 6 |

| 2 | 4 | 6 | 8 | 10 | 12 |

| 3 | 6 | 9 | 12 | 15 | 18 |

| 4 | 8 | 12 | 16 | 20 | 24 |

| 5 | 10 | 15 | 20 | 25 | 30 |

| 6 | 12 | 18 | 24 | 30 | 36 |

And there are 4 12's in the table.

# Answer

Super naive way of doing this is to just build the table then iterate through it. I imagine that'd be a 2\*O(n^2) effort though

More clever would be to do some multiplication and early exit? I still think it might be O(n^2)

12:

1 2 3 4 5 6

|---------| 6 - inc i

--|-------| 12 - add 2 - inc i

----|-----| 18 - dec j

----|---|-- 15 - dec j

----|-|---- 12 - add 2 - inc i

------|---- i == j -> exit

36:

1 2 3 4 5 6

|---------| 6 - inc i

--|-------| 12 - inc i

----|-----| 18 - inc i

------|---| 24 - inc i

--------|-| 30 - inc i

----------| 36 - add 1 - exit

this is a more logrithmic approach. we get 2, x2 for both sides = appears 4 times

6:

1 2 3 4 5 6

|---------| 6 - inc i, add 2

--|-------| 12 - dec j

--|-----|-- 10 - dec j

--|---|---- 8 - dec j

--|-|------ 6 - inc i add 2

----|------ 9 - exit
