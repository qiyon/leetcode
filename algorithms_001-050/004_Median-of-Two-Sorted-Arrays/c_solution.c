#include <stdio.h>

//leetcode submission begin ----------------------
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    int toDelNum = (nums1Size + nums2Size + 1) / 2 - 1;
    int oneSingle = (nums1Size + nums2Size) % 2;
    int getMidArr[] = {0, 0};
    int arrIndex = 0;
    int aDelNum, bDelNum;
    if (nums1Size == 0 && nums2Size == 0) return 0;
    while (1){
        if (toDelNum == 0){
            if (nums1Size == 0) getMidArr[arrIndex] = nums2[0];
            if (nums2Size == 0) getMidArr[arrIndex] = nums1[0];
            if (nums1Size != 0 && nums2Size !=0 ){
                getMidArr[arrIndex] = minVal(nums1[0], nums2[0]);
            }
            if (!oneSingle){
                oneSingle = 1;
                arrIndex = 1;
                toDelNum = 1 ;  //del more one
            }else{
                break;  //end
            }
        }else{
            aDelNum = (toDelNum + 1) / 2;
            bDelNum = aDelNum;
            if (aDelNum > nums1Size) aDelNum = nums1Size ;
            if (bDelNum > nums2Size) bDelNum = nums2Size ;
            if (bDelNum != 0 && ( aDelNum == 0 || nums1[aDelNum - 1] > nums2[ bDelNum - 1]) ){
                nums2 = nums2 + bDelNum;
                nums2Size = nums2Size - bDelNum;
                toDelNum = toDelNum - bDelNum;
            }else{
                nums1 = nums1 + aDelNum;
                nums1Size = nums1Size - aDelNum;
                toDelNum = toDelNum - aDelNum;
            }
        }
    }
    if (arrIndex == 1) {
        return (getMidArr[0] + getMidArr[1]) / 2.0;
    }else{
        return getMidArr[0];
    }
}

int minVal(int a, int b ){
    return (a > b) ? b : a;
}

//leetcode submission end ----------------------
int main()
{
    int a[] = {};
    int b[] = {1,2};
    double median;
    median = findMedianSortedArrays(a, sizeof(a)/sizeof(int), b, sizeof(b)/sizeof(int) );
    printf("%f\n", median);

}
