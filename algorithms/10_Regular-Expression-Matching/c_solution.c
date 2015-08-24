#include "../c_leetcode.h"

//leetcode submission begin ---------------------
bool isMatch(char* s, char* p) 
{    
    return 0;
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
