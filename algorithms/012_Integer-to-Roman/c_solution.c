#include "../c_leetcode.h"

//leetcode submission begin ------------------
char* intToRoman(int num)
{
    char * romanStr = (char *)malloc(sizeof(char)*10);
    romanStr = "Example";
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
