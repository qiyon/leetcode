/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */
var reverseKGroup = function(head, k) {
    var res, pre, first, end, rev_p, rev_tmp;
    var i;
    if ((k == 1) || !head) return head;
    first = head;
    while (true) {
        end = first;
        i = 1;
        while (end && end.next && i < k) {
            end = end.next;
            i++;
        }
        if (i != k) {
            break;
        }
        rev_p = first;
        while (true) {
            rev_tmp = first.next;
            first.next = rev_tmp.next;
            rev_tmp.next = rev_p;
            if (rev_tmp == end) {
                break;
            }
            rev_p = rev_tmp;
        }
        if (pre) {
            pre.next = end;
        } else {
            res = end;
        }
        pre = first;
        first = first.next;
    }
    return res ? res : head;
};


//------dev code------
function ListNode(val) {
   this.val = val;
   this.next = null;
}

var vals = [1,2,3,4,5,6,7,8,9,10,11];
var list = null, list_end;
vals.map(function(one_val){
    if (list) {
      list_end.next = new ListNode(one_val);
      list_end = list_end.next;
    } else {
      list = list_end = new ListNode(one_val);
    }
});
//run func
var out = reverseKGroup(list, 4);
//log
while (out) {
  console.log(out.val);
  out = out.next;
}
