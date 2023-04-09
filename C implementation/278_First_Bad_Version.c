/*
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check.
Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version.
You should minimize the number of calls to the API.
*/
#include <stdio.h>

#define NUMBER 5

int isBadVersion(int version) {
    if (version >= NUMBER) {
        return 1;
    } else {
        return 0;
    }
}

int firstBadVersion(int n) {
    int upper = n;
    int buttom = 0;
    int prev = n;

    int mid = n;
    while(mid != 0) {
      mid = (upper - buttom) / 2;
      if(!isBadVersion(buttom + mid)) {
        buttom = mid+buttom;
      } else {
        upper = buttom + mid;
        prev = buttom + mid;
      }
    }

    return prev;
}