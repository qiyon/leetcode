#include "../c_leetcode.h"

//leetcode submission begin ------------------------------
int maxArea(int* height, int heightSize) 
{
    int i, j, maxNum = 0, minValOfIJ, tmpArea, tmpIVal, tmpJVal;
    i = 0;
    j = heightSize - 1;
    while(i < j){
        minValOfIJ = (height[i] > height[j]) ? height[j] : height[i];
        tmpArea = (j - i) * minValOfIJ;
        if (maxNum < tmpArea) maxNum = tmpArea;
        //printf("i,j:%d,%d  now maxNum:%d\n---\n",i,j,maxNum);
        tmpIVal = height[i];
        tmpJVal = height[j];
        if (tmpIVal < tmpJVal){
            while( i < j && height[i] <= tmpIVal) i++;
        }else{
            while( i < j && height[j] <= tmpJVal) j--;
        }
    }
    return maxNum;
}
//leetcode submission end ------------------------------

int main(int argc, char* argv)
{
    int height[] = {5,6,2,10,2,300,4,5,10,100, 99};
    printf("maxArea:%d\n",maxArea(height,11));
}
