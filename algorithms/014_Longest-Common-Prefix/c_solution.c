#include "../c_leetcode.h"

//leetcode submission begin-----------
char* longestCommonPrefix(char** strs, int strsSize) 
{
    char * longPreStr;
    int i, j, slen;
    char tmpChar;

    if (strsSize == 0) {
        longPreStr = (char *)malloc(sizeof(char));
        longPreStr = "";
        return longPreStr;
    };
    slen = strlen(strs[0]);
    longPreStr = (char *)malloc(sizeof(char)*(slen+1));
    for(i = 0; i < slen; i++){
        tmpChar = '\0';
        for(j = 0; j < strsSize; j++){
            if (tmpChar == '\0') tmpChar = strs[j][i];
            if (tmpChar == '\0' || tmpChar != strs[j][i]){
                tmpChar = '\0';
                break;
            } 
        }
        if (tmpChar == '\0') break;
        longPreStr[i] = tmpChar;
    }
    longPreStr[i] = '\0';
    return longPreStr;
}
//leetcode submission end-----------



int main(int argc, char * argv[])
{
    //if (argc < 2){
    //    printf("need arg\n");
    //}else{
        printf("long prefix :%s \n", longestCommonPrefix(argv+1, argc - 1));
    //}
    return 0;
}
