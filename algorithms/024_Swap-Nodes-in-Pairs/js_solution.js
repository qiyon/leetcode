/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function(head) {
    if (!(head && head.next)) return head;
    var res = head.next;
    var pre = head;
    var end = (res.next && res.next.next) ? res.next.next : null;
    pre.next = res.next;
    res.next = pre;
    while (end) {
        pre.next.next = end.next;
        end.next = pre.next;
        pre.next = end;
        //mv pointer
        pre = end.next;
        if (pre.next && pre.next.next) {
            end = pre.next.next;
        } else {
            break;
        }
    }
    return res ? res : head;
};

//------dev code------
function ListNode(val) {
   this.val = val;
   this.next = null;
}

var vals = [1,2,3,4,5,6,7,8];
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
var out = swapPairs(list);
//log
while (out) {
  console.log(out.val);
  out = out.next;
}
