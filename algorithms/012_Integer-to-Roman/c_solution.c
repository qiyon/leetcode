#include "../c_leetcode.h"

//leetcode submission begin ------------------
char* intToRoman(int num)
{
    char oneArr[]  = {'I','X','C','M'};
    char fiveArr[] = {'V','L','D'};
    int i = 0, j=0, romanIndex = 0, modValue;
    char tmpRoman[4][5];
    char * romanStr = (char *)malloc(sizeof(char)*15);

    for(i = 0;i<4;i++){
        for(j=0;j<5;j++){
            tmpRoman[i][j] = '\0';
        }
    }
    if (num == 0 || num > 3999) {
        romanStr = "";
        return romanStr;
    }
    i = 0;
    while (num) {
        modValue = num % 10;
        romanIndex = 0;
        if (modValue == 4 || modValue == 9){
            tmpRoman[i][romanIndex] = oneArr[i];
            romanIndex++;
            tmpRoman[i][romanIndex] = (modValue == 4) ? fiveArr[i] : oneArr[i+1] ;
            romanIndex++;
            modValue = 0;
        }
        if (modValue >= 5 && modValue < 9){
            tmpRoman[i][romanIndex] = fiveArr[i];
            romanIndex++;
            modValue = modValue - 5;
        }
        while (modValue) {
            tmpRoman[i][romanIndex] = oneArr[i];
            romanIndex++;
            modValue--;
        }
        num = num / 10;
        i++;
    }
    romanIndex = 0;
    for (i = 3 ; i>=0; i--){
        for(j = 0;j < 5; j++){
            if (tmpRoman[i][j] != '\0'){
                romanStr[romanIndex] = tmpRoman[i][j]; 
                romanIndex++;
            }
        }
    }
    romanStr[romanIndex] = '\0';
    return romanStr;
}
//leetcode submission end ------------------

int main(int argc, char* argv[] )
{
    int changeNum ;
    if (argc < 2){
        printf("please input arg of val\n");
    }else{
        changeNum = atoi( argv[1] );
        printf("Number %d to Roman is: %s \n", changeNum, intToRoman(changeNum)); 
    }
}
