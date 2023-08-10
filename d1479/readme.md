# Question
Given a pivot x, and a list lst, partition the list into three parts.

The first part contains all elements in lst that are less than x
The second part contains all elements in lst that are equal to x
The third part contains all elements in lst that are larger than x
Ordering within a part can be arbitrary.

For example, given x = 10 and lst = [9, 12, 3, 5, 14, 10, 10], one partition may be [9, 3, 5, 10, 10, 12, 14].

# Answers
## Option 1: Naive
- Create 3 arrays, each for holding elements smaller than, elements equal to, and elements larger than the pivot
- merge arrays back into one array

## Option 2: In-Place sorting and sorting isn't important
- Swap + Pointers
- start = beginning of the array
- end = end of the array
- iterate through array
- if value > pivot - move it to the end
- if value < pivot - move it to the start
- if value == pivot - do nothing
Iterations:
1) 9 is < 10, move to the start of the array = arr = [9, 12, 3, 5, 14, 10, 10]
    start++
2) 12 is > 10, move to the back of array = [9, 3, 5, 14, 10, 10, 12]
    end--
3) 3 is < 10, move to the start of the array  = [3, 9, 5, 14, 10, 10, 12]
    start++
4) 5 is < 10, move it to the start of the array = [5, 3, 9, 14, 10, 10, 12]