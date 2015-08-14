#include <stdio.h>
#include <string.h>

//leetcode submission begin-----------------
int lengthOfLongestSubstring(char* s) {
    int alphabet[256] ;
    int charNum=0 , start=0 ;
    int i=0 ;
    int charIndex, betVal;
    for (i = 0 ; i<256;i++){
        alphabet[i] = -1;
    }
    i = 0;
    while (s[i] != '\0'){
        charIndex = (int)(s[i]);
        betVal = alphabet[charIndex];
        //printf("%d\t%d\t%d\n",i, charIndex, betVal);
        if (betVal >= 0 && betVal >= start){
            start = betVal + 1;
        }
        alphabet[charIndex] = i;
        if (i - start + 1 > charNum) charNum = i - start + 1;
        i++;
    }
    return charNum;
}
//leetcode submission end-----------------

int main()
{
    int nums = lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ abcd'");
    printf("long :%d\n", nums);
}
