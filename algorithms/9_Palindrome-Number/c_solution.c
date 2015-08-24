#include "../c_leetcode.h"

//leetcode submission begin ----------
bool isPalindrome(int x) 
{
    int halfOne;
    int modOut;
    if (x < 0 ) return false;
    if (x < 10) return true;
    halfOne = 0;
    while(x){
        modOut = x % 10; 
        if (modOut == 0 && halfOne == 0) return false;
        if (halfOne == x) return true;
        halfOne = halfOne * 10 + modOut;
        if (halfOne == x) return true;
        x = x / 10;
        if (x < halfOne) return false;
    }
    return false;
}
//leetcode submission end ----------

int main(int argc, char* argv[])
{
    int num ;
    if (argc >=2 ){
        num = atoi(argv[1]);
        if (isPalindrome(num)){
            printf("TRUE\n");
        }else{
            printf("FALSE\n");
        }
    }else{
        printf("please input one arg\n");
    }
}
