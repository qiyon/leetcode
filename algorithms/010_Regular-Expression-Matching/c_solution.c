#include "../c_leetcode.h"


void selfPrintStatus(int *s, int length)
{
    int i;
    for(i=0; i< length; i++)
    {
        printf("%d,", s[i]);
    }
    printf("\n--------------------\n");
}

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
    int i = 0;
    if (length <= insertVal) return;
    while(i < length){
        if (statusArr[i] == insertVal) return;
        if (statusArr[i] == -1 ){
            statusArr[i] = insertVal;
            return;
        }
        i++;
    }
}

bool isMatch(char* s, char* p) 
{
    int slen = strlen(s), plen = strlen(p); //length of string and pattern
    int * ruleIndex, * ruleRepeat;
    int * nowStatus, * preStatus, * swapStatusTmp;
    int i = 0, j = 0, countOfRule = 0, tmpMoveIndex;
    int valIsMatch = 0;
    int checkStatusTmp;

    if (plen == 0) {
        return (slen == 0) ? 1 : 0;
    }
    ruleIndex = (int *)malloc(sizeof(int) * plen);
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
    //printf("count of rule:%d\n---\n",countOfRule);
    nowStatus = (int *)malloc(sizeof(int)*(countOfRule + 1));
    preStatus = (int *)malloc(sizeof(int)*(countOfRule + 1));
    initStatus(nowStatus, countOfRule + 1);
    initStatus(preStatus, countOfRule + 1);
    preStatus[0] = 0;
    for(i = 0 ; i < slen; i++){
        //printf("--------one char ------------------------\n");
        //selfPrintStatus(preStatus,countOfRule+1);
        for(j = 0; j < countOfRule + 1; j++){
            if (preStatus[j] == -1) continue;
            //printf("---\npresataus[j]: %d ,s[i]=%c,ruleChar=%c \n---\n",preStatus[j] , s[i], p[ruleIndex[preStatus[j]]]  );
            tmpMoveIndex = preStatus[j];
            while(tmpMoveIndex < countOfRule && ruleRepeat[tmpMoveIndex]){
                if (s[i] == p[ruleIndex[tmpMoveIndex]] || p[ruleIndex[tmpMoveIndex]] == '.'){
                    insertVal(nowStatus, countOfRule+1, tmpMoveIndex);
                }
                tmpMoveIndex++;
            }
            if ( tmpMoveIndex < countOfRule && ( s[i] == p[ruleIndex[tmpMoveIndex]] || p[ruleIndex[tmpMoveIndex]] == '.')){
                insertVal(nowStatus, countOfRule+1, tmpMoveIndex + 1);
            }
            //selfPrintStatus(nowStatus,countOfRule+1);
            //printf("i:%d ,j:%d ,countOfRule:%d , tmpMoveIndex:%d \n--------------\n", i, j , countOfRule, tmpMoveIndex);
        }
        swapStatusTmp = preStatus;
        preStatus = nowStatus;
        nowStatus =swapStatusTmp;
        initStatus(nowStatus, countOfRule+1);
    }
    for(i = 0; i < countOfRule + 1; i++){
        checkStatusTmp = preStatus[i]; 
        if (checkStatusTmp == -1) continue;
        while( checkStatusTmp < countOfRule && ruleRepeat[checkStatusTmp]) checkStatusTmp++;
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
