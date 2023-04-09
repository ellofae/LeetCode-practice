/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

#include <stdio.h>

int search(int* nums, int numsSize, int target){
    int buttom = 0;
    int upper = numsSize;
    int mid = numsSize;

    while(mid != 0) {
        mid = (upper - buttom) / 2;

        if (target > nums[buttom + mid]) {
            buttom = buttom + mid;
        } else if (target < nums[buttom + mid]) {
            upper = upper-mid;
        } else {
            return (buttom+mid);
        }
    }

    return -1;
}