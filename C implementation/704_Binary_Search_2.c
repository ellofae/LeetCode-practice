/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

#include <stdio.h>

int search(int* nums, int numsSize, int target){
    int buttom = 0;
    int upper = numsSize-1;
    int mid;

    while(buttom <= upper) {
        mid = (upper + buttom) / 2;

        if (nums[mid] == target) {
            return mid;
        }
        
        if(nums[mid] > target) {
            upper = mid - 1;
        } else {
            buttom = mid + 1;
        }
    }

    return -1;
}
