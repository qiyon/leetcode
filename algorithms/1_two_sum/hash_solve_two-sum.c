#include <stdio.h>
#include <stdlib.h>


//leetcode submissions begin ----------------------------------

#define QY_HASH_TABLE_MAX 10000

struct qy_hash_node{
    int key;
    int val;
    struct qy_hash_node * next  ;
}

struct qy_hash_table{
    struct qy_hash_node * listNode[QY_HASH_TABLE_MAX]  ;
}
//statement functions
struct qy_hash_node * newNode(int key, int val);
struct qy_hash_table * newTable();
int hashIntKey(int key);
int tableSetVal(struct qy_hash_table * table, int key, int val);
int tableGetByKey(struct qy_hash_table * table, int key);
int destroyTable(struct qy_hash_table * table);

// new hashtable's node
struct qy_hash_node * newNode(int key, int val)
{
    struct qy_hash_node * node;
    node = (struct qy_hash_node *)molloc(sizeof(struct qy_hash_node));
    node->key = key;
    node->val = val;
    node->next = NULL;
    return node;
}

struct qy_hash_table * newTable()
{
    struct qy_hash_table * table;
    int i;
    table = (struct qy_hash_table *)molloc(sizeof(struct qy_hash_table));
    for(i = 0; i < QY_HASH_TABLE_MAX ; i++){
        table->listNode[i] = NULL;
    }
    return table;
}

int hashIntKey(int key)
{
    return key % QY_HASH_TABLE_MAX;
}

int tableSetVal(struct qy_hash_table * table, int key, int val)
{
    int hashKey ;
    struct qy_hash_node * lastNode = NULL;
    hashKey = hashIntKey(key);
    if (table->listNode[hashKey]){
        lastNode = table->listNode[hashKey];
        while(lastNode){
            if (lastNode->key == key){
                lastNode->val = val;
                return 0;
            }
            if (lastNode->next){
                lastNode = lastNode->next;
            }else{
                lastNode->next = newNode(key, val);
                return 0;
            }
        }
    }else{
        table->listNode[hashKey] = newNode(key, val);
    }
    return 0;
}

int tableGetByKey(struct qy_hash_table * table, int key)
{
    int hashKey;
    struct qy_hash_node * lastNode = NULL;
    hashKey = hashIntKey(key); 
    lastNode = table->listNode[hashKey];
    while(lastNode){
        if (lastNode->key == key) return lastNode->val;
        lastNode = lastNode->next;
    }
    return -1;
}

int destroyTable(struct qy_hash_table * table)
{
    //todo
}


int* setReturn(int index1, int index2)
{
    int * retIndex = (int *)malloc(2 * sizeof(int));
    *retIndex = index1 + 1;
    *(retIndex+1) = index2 + 1;
    return retIndex;
}

int* twoSum(int* nums, int numsSize, int target)
{
    //todo
    return setReturn(0, 0);
}
//leetcode submissions end --------------------------------


int main()
{
    int nums[] = {4,5,11,15};
    int * retAnwser ;
    retAnwser = twoSum(nums, 4, 9);
    printf("index1=%d index2=%d\n", *retAnwser,*(retAnwser+1));
}
