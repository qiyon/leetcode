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
    while(l1 != NULL || l2 != NULL || cin!=0){
        firstVal = (l1 == NULL) ? 0 : l1->val;
        secondVal = (l2 == NULL) ? 0 : l2->val;
        firstVal = firstVal + secondVal + cin;
        firstVal = firstVal % 10;
        newNode = malloc(sizeof(struct ListNode));
        newNode->val = firstVal;
        if (sumHead == NULL){
            sumHead = newNode;
            sumEnd = newNode;
        }else{
            sumEnd->next = newNode;
            sumEnd = sumEnd->next;
        }
        // for next
        cin = firstVal / 10;
        if (l1 != NULL) l1 = l1->next;
        if (l2 != NULL) l2 = l2->next;
    }
    return sumHead;
}

int main(){
    struct ListNode *hehe;
    struct ListNode a;
    struct ListNode b;
    a.val = 1;
    a.next = NULL;
    b.val = 4;
    b.next=NULL;
    hehe = addTwoNumbers(&a, &b);
    printf("%d\n", hehe->val);
}
