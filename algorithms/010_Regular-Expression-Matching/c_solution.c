#include "../c_leetcode.h"

//leetcode submission begin ---------------------
void initStatus(int *statusArr, int length)
{
    int i = 0;
    while(i < length){
        statusArr[i] = -1;
        i++;
    }
}
void insertVal(int * statusArr, int length, int insertVal)
{
    int i = 0, relaceIndex = -1;
    if (insertVal >= length) return;
    while(i < length){
        if (relaceIndex == -1 && statusArr[i] == -1) relaceIndex = i;
        if (statusArr[i] == insertVal) return;
        i++;
    }
    printf("---\nset status[%d] = %d \n---\n", relaceIndex, insertVal);
    statusArr[relaceIndex] = insertVal;
}

void swapStatus(int * s1, int * s2)
{
    int * tmp;
    tmp = s1;
    s1 = s2;
    s2 = tmp;
}
void selfPrintStatus(int *s, int length)
{
    int i;
    for(i=0; i< length; i++)
    {
        printf("%d,", s[i]);
    }
    printf("\n--------------------\n");
}

bool isMatch(char* s, char* p) 
{
    int slen = strlen(s), plen = strlen(p); //length of string and pattern
    int * ruleIndex, * ruleRepeat;
    int * nowStatus, * preStatus;
    int i = 0, j = 0, countOfRule = 0;
    int valIsMatch = 0;
    int checkStatusTmp;

    ruleIndex = (int *)malloc(sizeof(int) * plen );
    ruleRepeat = (int *)malloc(sizeof(int) * plen);
    while( i < plen ){
        ruleIndex[countOfRule] = i;
        if (i+1 < plen && p[i+1] == '*'){
            ruleRepeat[countOfRule] = 1;
            i = i+2;
        }else{
            ruleRepeat[countOfRule] = 0;
            i = i+1;
        }
        countOfRule++;
    }
    nowStatus = (int *)malloc(sizeof(int)*(countOfRule + 1));
    preStatus = (int *)malloc(sizeof(int)*(countOfRule + 1));
    initStatus(nowStatus, countOfRule + 1);
    initStatus(preStatus, countOfRule + 1);
    preStatus[0] = 0;
    for(i = 0 ; i < slen; i++){
        selfPrintStatus(preStatus,countOfRule+1);
        for(j = 0; j < countOfRule + 1; j++){
            selfPrintStatus(nowStatus,countOfRule+1);
            if (preStatus[j] == -1) continue;
            printf("---\npresataus[j]: %d \n---\n",preStatus[j]  );
            if (s[i] == p[ruleIndex[preStatus[j]]] || p[ruleIndex[preStatus[j]]] == '.'){
                insertVal(nowStatus, countOfRule+1, preStatus[j]+1);
            }
            if (ruleRepeat[preStatus[j]])
                insertVal(nowStatus, countOfRule+1, preStatus[j]);
        }
        swapStatus(nowStatus, preStatus);
        initStatus(nowStatus, countOfRule+1);
    }
    for(i = 0; i < countOfRule + 1; i++){
        checkStatusTmp = preStatus[i]; 
        while( checkStatusTmp < countOfRule && ruleRepeat[checkStatusTmp+1]) checkStatusTmp++;
        if (checkStatusTmp == countOfRule){
            valIsMatch = 1;
            break;
        }
    }

    free(ruleIndex);
    free(ruleRepeat);
    free(nowStatus);
    free(preStatus);
    return valIsMatch;
}
//leetcode submission end ---------------------

int main(int argc, char* argv[])
{
    if (argc < 3 ) {
        printf("please int args\n");
    }else{
        if (isMatch(argv[1], argv[2])){
            printf("true\n");
        }else{
            printf("false\n");
        }
    }
}
