
/**
 *  * Note: The returned array must be malloced, assume caller calls free().
 *   */
int* setReturn(int index1, int index2)
{
    int * retIndex = (int *)malloc(2 * sizeof(int));
    *(retIndex+1) = index2+1;
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
