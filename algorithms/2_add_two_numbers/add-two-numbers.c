#include <stdio.h>
#include <stdlib.h>
struct ListNode {
    int val;
    struct ListNode *next;
};

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
        cin = firstVal / 10;  //cin for next
        firstVal = firstVal % 10;
        newNode = (struct ListNode*)malloc(sizeof(struct ListNode));
        newNode->val = firstVal;
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

int main(){
    struct ListNode *hehe;
    struct ListNode a = {9, NULL};
    struct ListNode b = {4, NULL};
    struct ListNode c = {1, NULL};
    a.next = &c;
    hehe = addTwoNumbers(&a, &b);
    printf("%d,%d\n", hehe->val, hehe->next->val);
}
