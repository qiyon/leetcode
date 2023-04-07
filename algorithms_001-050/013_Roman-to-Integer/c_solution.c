#include "../c_leetcode.h"

//leetcode submission begin -------------------
int romanToInt(char* s) {
    char matchChar[] = {'I', 'V', 'X', 'L', 'C', 'D', 'M'};
    int  matchInt[] =  {1,    5,   10,  50,  100, 500, 1000};
    int  i =0, integerVal = 0, indexMatch, slen = strlen(s) ,preMatchVal, nowMatchVal;

    preMatchVal = 0;
    for (i = 0 ; i < slen ; i++ ){
        for (indexMatch = 0 ; indexMatch < 7; indexMatch ++ ){
            if (matchChar[indexMatch] == s[i]) break;
        }
        if (indexMatch == 7) return 0;
        nowMatchVal = matchInt[indexMatch];
        if (preMatchVal == 0) preMatchVal = nowMatchVal;
        if (preMatchVal >= nowMatchVal){
            integerVal += nowMatchVal;
        }else{
            integerVal += nowMatchVal - 2 * preMatchVal;
        }
        preMatchVal = nowMatchVal;
    }
    return integerVal;
}
//leetcode submission end -------------------

int main(int argc, char * argv[])
{
    if (argc < 2){
        printf("Need arg \n");
    }else{
        printf("Roman %s to integer is %d \n", argv[1], romanToInt(argv[1]));
    }
}
