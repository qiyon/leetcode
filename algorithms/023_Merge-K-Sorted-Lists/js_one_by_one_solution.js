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
    var len = lists.length;
    var res = null;
    var i;
    for (i = 0; i< len; i++) {
        res = mergeTwoLists(lists[i], res);
    }
    return res;

    function mergeTwoLists(a, b) {
        var res = null, mv_p = null;
        var min_is_a;
        while (a || b) {
            min_is_a = null;
            if (a && b) {
                min_is_a = (a.val > b.val) ? false : true;
            } else {
                min_is_a = a ? true : false;
            }
            if (mv_p) {
                mv_p.next = min_is_a ? a : b;
                mv_p = mv_p.next;
            } else {
                mv_p = min_is_a ? a : b;
                res = mv_p;
            }
            if (min_is_a) {
                a = a.next;
            } else {
                b = b.next;
            }
        }
        return res;
    }
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
