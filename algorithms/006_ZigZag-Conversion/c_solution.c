#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//leetcode submission bgein -----------------
char* convert(char* s, int numRows)
{
    int slen = strlen(s);        
    int i, j;
    int modNum = numRows*2 - 2;
    char *outStr;
    int indexOfOutStr = 0;
    int besideIndex;
    
    if (slen < 3 || slen <= numRows || numRows < 2) return s;
    outStr= (char *)malloc(sizeof(char)*(slen + 1));
    for(i = 0; i < numRows; i++){
        j = i;
        while (j < slen){
            outStr[indexOfOutStr] = s[j];
            indexOfOutStr++;
            besideIndex = j + (numRows - i - 1) * 2;
            if (i > 0 && i < numRows-1  && besideIndex < slen ){
                outStr[indexOfOutStr] = s[besideIndex];
                indexOfOutStr++;
            }
            j += modNum;
        }
    }
    outStr[indexOfOutStr] = '\0';
    return outStr;
}
//leetcode submission end -----------------

int main(int argc, char* argv[])
{
    int nums = atoi(argv[2]);
    printf("%s\n", convert(argv[1], nums));
}
