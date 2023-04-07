#include <stdio.h>
#include <string.h>
#include <stdlib.h>

//leetcode submission begin -----------------
char* longestPalindrome(char* s) {
    int indexLong[1001];        
    int slen = strlen(s);
    int maxBegin = 0;
    int i = 0, j;
    int nowLength = 0;
    int loopCount = 0;
    int besideSame;
    char *retStr;
    if (slen <= 1) return s;
    while (1){
        loopCount++;
        for (i = 0; i < slen ; i++){
            if (loopCount == 1 ){
                besideSame = 0;
                while (i + 1 +besideSame < slen && s[i] == s[i + 1 +besideSame]) besideSame++;
                indexLong[i] = 1 + besideSame;
            }else{
                j = i + indexLong[i+1] + 1;
                if (i+1 < slen && j < slen && s[i] == s[j]){
                    indexLong[i] = indexLong[i+1] + 2;
                }
            }
            if (indexLong[i] > nowLength){
                nowLength = indexLong[i];
                maxBegin = i;
            } 
        }
        //printf("nowLength:%d  loopCount:%d\n----------\n", nowLength, loopCount);
        if (nowLength < loopCount*2 - 1) break;
    }
    retStr = (char *)malloc(sizeof(char) * (nowLength + 1));
    for(i = 0; i < nowLength; i++){
        retStr[i] = s[maxBegin+i];
    }
    retStr[i] = '\0';
    return retStr;
}

//leetcode submission end -----------------
int main()
{
    char * out;
    printf("%s\n",longestPalindrome("aaabaaaa"));
}
