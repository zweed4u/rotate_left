#!/usr/bin/python3


# https://www.hackerrank.com/challenges/ctci-array-left-rotation/problem
def rotLeft(a, d):
    """
    a: List[int]
    d: int - number of times to rotate
        
        ie. 4 &  [1,2,3,4,5]  -> [2,3,4,5,1] -> [3,4,5,1,2] -> [4,5,1,2,3] -> [5,1,2,3,4] 
    """
    new_a = []
    length_of_array = len(a)
    # protect against wrap around - multiples of lengths longer than the array can be boiled down
    # ie instead of wrapping around an array twice (rotate 7) you can simplify it to once (rotate 2)
    # for an array of length 5
    if d > len(a):
        d %= len(a)
    for i in range(length_of_array):
        # determine the new index of the value to occupy the current index, protecting against
        # out of bounds where i is close to the last index in the array
        new_value_index = (i + d) % len(a)
        new_a.append(a[new_value_index])
    return new_a


if __name__ == "__main__":
    a = [1, 2, 3, 4, 5]
    d = 4
    print(f"Original array:\n{a}")
    result = rotLeft(a, d)
    print(f"Array transformed after rotating each element {d} to the left\n{result}")
