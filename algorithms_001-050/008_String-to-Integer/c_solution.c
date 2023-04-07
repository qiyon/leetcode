#include "../c_leetcode.h"

//leetcode submission bgein -----------------------------
int myAtoi(char* str) {
    int slen = strlen(str);
    int i = 0;
    int isNeg = 0;
    int outPutNum = 0;
    int notSignYet = 1;
    while( (str[i] == ' ' || str[i] == '-' || str[i] == '+') && notSignYet ){
        if (str[i] == '-') isNeg = 1;
        if (str[i] == '+' || str[i] == '-') notSignYet = 0;
        i++;
        //printf("%c\n---------\n", str[i]);
    }
    while( str[i] >= '0' && str[i]<='9' && i<slen) {
        if (isNeg){
            if (outPutNum > INT_MAX / 10 || (outPutNum == INT_MAX /10 && str[i] >= '8')) return INT_MIN;
        }else{
            if (outPutNum > INT_MAX / 10 || (outPutNum == INT_MAX /10 && str[i] >= '7')) return INT_MAX;
        }
        outPutNum = outPutNum * 10 + (str[i] - '0');
        i++;
    }
    return (isNeg) ? -outPutNum : outPutNum;
}
//leetcode submission end -----------------------------

int main(int argc, char* argv[])
{
    if (argc >=2 ){
        printf("myAtoi out:%d\n", myAtoi(argv[1]) );
    }else{
        printf("please input a string\n");
    }
}
