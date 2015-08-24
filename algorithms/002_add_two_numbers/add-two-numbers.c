#include <stdio.h>
#include <stdlib.h>
struct ListNode {
    int val;
    struct ListNode *next;
};

//leetcode submission begin ----------------------------------------------
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int firstVal, secondVal;
    int cin = 0;
    struct ListNode * sumHead=NULL ;
    struct ListNode * sumEnd=NULL ;
    struct ListNode * newNode ;
    while(l1 || l2 || cin!=0){
        firstVal = l1 ? l1->val : 0;
        secondVal = l2 ?  l2->val : 0;
        firstVal = firstVal + secondVal + cin;
        //cin = firstVal / 10;  //cin for next
        cin = (firstVal >9 ) ? 1 : 0;
        //firstVal = firstVal % 10;
        firstVal = cin ? firstVal - 10 : firstVal;
        newNode = (struct ListNode*)malloc(sizeof(struct ListNode));
        newNode->val = firstVal;
        newNode->next = NULL;
        if (sumHead == NULL){
            sumHead = newNode;
            sumEnd = newNode;
        }else{
            sumEnd->next = newNode;
            sumEnd = sumEnd->next;
        }
        // for next
        if (l1) l1 = l1->next;
        if (l2) l2 = l2->next;
    }
    return sumHead;
}
//leetcode submission end ----------------------------------------------

struct ListNode* newList(int arr[], int length)
{
    int i ;
    struct ListNode * sumHead=NULL ;
    struct ListNode * sumEnd=NULL ;
    struct ListNode * newNode ;
    for(i = 0; i < length; i++){
        newNode = (struct ListNode*)malloc(sizeof(struct ListNode));
        newNode->val = arr[i];
        newNode->next = NULL;
        if (sumHead == NULL){
            sumHead = newNode;
            sumEnd = newNode;
        }else{
            sumEnd->next = newNode;
            sumEnd = sumEnd->next;
        }
    }
    return sumHead;
}
int printList(struct ListNode * listHead)
{
    while(listHead){
        printf("%d,", listHead->val);
        listHead = listHead->next;
    }
}

int main(){
    int a[] = {1,1,0,4,7,2,7,3,0,1 };
    int b[] = {2,5,5,6,8,0,3,0,5,9 };
    struct ListNode * hehe;
    struct ListNode * hehea;
    struct ListNode * heheb;

    hehea = newList(a, 10);
    heheb = newList(b, 10);
    printList(hehea);
    printf("\n");
    printList(heheb);
    printf("\n");

    hehe = addTwoNumbers( hehea, heheb);
    printList(hehe);
    printf("\n");
}
