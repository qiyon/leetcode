/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function(lists) {
  var head = null, res = null, min_list_num;
  var k = lists.length;
  var all_null, i;
  while (true) {
    all_null = true;
    min_list_num = null;
    for (i = 0; i < k; i++) {
      if (lists[i]){
        all_null = false;
        if (min_list_num === null) {
          min_list_num = i;
        } else if (lists[min_list_num].val > lists[i].val) {
          min_list_num = i;
        }
      }
    }
    if (all_null) {
      break;
    }
    if (res === null) {
      res = lists[min_list_num];
      head = res;
    } else {
      res.next = lists[min_list_num];
      res = res.next;
    }
    lists[min_list_num] = lists[min_list_num].next;
  }
  return head;
};

//--------submission end----------

function ListNode(val) {
  this.val = val;
  this.next = null;
}

var list1 = new ListNode(1);
list1.next = new ListNode(5);
list1.next.next = new ListNode(10);
var list2 = new ListNode(2);
list2.next = new ListNode(6);
var list3 = new ListNode(8);
list3.next = new ListNode(15);

var head = mergeKLists([list1, list2, list3]);

console.log('------result-----');
while (head) {
  console.log(head.val);
  head = head.next;
}


var res = mergeKLists([new ListNode(1), new ListNode(2)]);
console.log(res);
