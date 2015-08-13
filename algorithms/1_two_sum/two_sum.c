#include <stdio.h>
#include <stdlib.h>


//leetcode submissions begin ----------------------------------
int* setReturn(int index1, int index2)
{
    int * retIndex = (int *)malloc(2 * sizeof(int));
    *retIndex = index1 + 1;
    *(retIndex+1) = index2 + 1;
    return retIndex;
}

int* twoSum(int* nums, int numsSize, int target)
{
    int bigOne, littleOne;
    int little;
    for( bigOne = 1; bigOne < numsSize; bigOne++){
        little = target - (* (nums + bigOne));
        for( littleOne = 0; littleOne < bigOne; littleOne++){
            if (little == (* (nums+littleOne))){
                return setReturn(littleOne, bigOne);
            }
        }
    }
    return setReturn(0, 0);
}
//leetcode submissions end --------------------------------


int main()
{
    int nums[] = {4,5,11,15};
    int * retAnwser ;
    retAnwser = twoSum(nums, 4, 9);
    printf("index1=%d index2=%d\n", *retAnwser,*(retAnwser+1));
}
