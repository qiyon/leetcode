#include "../c_leetcode.h"

//leetcode submission begin ---------------------
void reFreshStatus(int * l, int len)
{
    len = len + 1;
    int i = 0;
    for( ; i < len ; i++){
        l[i] = -2;
        //printf("init val :%d \n----------\n", l[i]);
    }
}

void insertToStatus(int * l, int len, int statusIndex)
{
    int i = 0;
    int firstEmptySpace = -1;
    len = len + 1;
    for ( ; i < len ; i++){
        if (l[i] == -2 && firstEmptySpace == -1) firstEmptySpace = i;
        if (l[i] == statusIndex) return ;
    }
    l[firstEmptySpace] = statusIndex;
    printf("set val :%d \n----------\n", statusIndex);
}

bool isMatch(char* s, char* p) 
{
    int * pIndexBegin, * pCanRepeat;
    int slen = strlen(s) , plen = strlen(p);
    int ruleNum = 0;
    int i = 0, j = 0, ccc = 0;;
    int * nowRuleStatus , * preRuleStatus , * swapTmp;
    int outOfIsMatch = 0;
    int tmpPreStatus , tmpShouldCheckStatus  ;
    char tmpCheckChar;
    pIndexBegin = (int *)malloc(sizeof(int)*plen);
    pCanRepeat = (int *)malloc(sizeof(int)*plen);
    while(i <= plen){
        pIndexBegin[ruleNum] = i;
        if  (i + 1 < plen && p[i + 1] == '*'){
            i = i + 2;
            pCanRepeat[ruleNum] = 1;
        }else{
            i++;
            pCanRepeat[ruleNum] = 0;
        }
        ruleNum ;
    }
    nowRuleStatus = (int *)malloc(sizeof(int)* (ruleNum+1));
    preRuleStatus = (int *)malloc(sizeof(int)* (ruleNum+1));
    //init 
    reFreshStatus(nowRuleStatus, ruleNum);
    reFreshStatus(preRuleStatus, ruleNum);
    preRuleStatus[0] = -1;
    for(i = 0; i < slen ; i++){
        for(j = 0; j < ruleNum ; j++){
            tmpPreStatus = preRuleStatus[j];
            if (tmpPreStatus == -2) continue;
            tmpShouldCheckStatus = tmpPreStatus + 1;
            if (tmpShouldCheckStatus  < ruleNum){
                tmpCheckChar = p[pIndexBegin[tmpShouldCheckStatus]];
                if (tmpCheckChar == s[i] || tmpCheckChar == '.')
                    insertToStatus(nowRuleStatus, ruleNum, tmpShouldCheckStatus);
            } 
            if (pCanRepeat[tmpShouldCheckStatus]) 
                insertToStatus(nowRuleStatus, ruleNum, tmpPreStatus);
        }
        for (ccc = 0; ccc < ruleNum + 1; ccc++){
            printf("%d,", nowRuleStatus[ccc]);
        }
        printf("\n-----\n");
        swapTmp = preRuleStatus;
        preRuleStatus = nowRuleStatus;
        nowRuleStatus = swapTmp;
        reFreshStatus(nowRuleStatus, ruleNum);
    }
    for(i = 0; i < ruleNum ; i++){
        if (preRuleStatus[i] == -2) continue;
        for (j = preRuleStatus[i] + 1; j < ruleNum ; j++){
            if (!pCanRepeat[j]) break;
        }
        if (j == ruleNum) {
            outOfIsMatch = 1;
            break;
        }
    }
    free(pIndexBegin);
    free(pCanRepeat);
    free(nowRuleStatus);
    free(preRuleStatus);
    //get result & return
    return outOfIsMatch;
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
