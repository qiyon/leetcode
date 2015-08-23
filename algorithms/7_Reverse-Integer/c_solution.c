#include <stdio.h>
#include <stdlib.h>

//leetcode submission begin--------------------
int reverse(int x) 
{
    long long isNeg = 0;
    long long retNum = 0;
    long long longx = x;
    //int preRetNum = 0;
    if (x>-10 && x<10) return x;
    if (longx < 0){
        longx  = -longx;
        isNeg = 1;
    }
    while(longx){
        retNum = (retNum * 10) + (longx % 10);
        //printf("nowNum:%ld now_x:%ld\n---------------\n", retNum, longx);
        if ( retNum > 2147483647 + isNeg ){
            return 0; 
        }
        longx = longx / 10;
    }
    return (isNeg) ? -retNum : retNum;
}
//leetcode submission end--------------------

int main(int argc, char * argv[])
{
    int x = atoi(argv[1]);
    printf("%d\n", reverse(x));
}
