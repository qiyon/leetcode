/**
 * @param {number} n
 * @return {string[]}
 **/
var generateParenthesis = function(n) {
    var out = [], res = [], left_c = [];
    var tmp, len, lc;
    out.push('(');
    left_c.push(1);
    while (out.length > 0) {
        tmp = out.pop();
        len = tmp.length;
        lc = left_c.pop();
        if (len == 2 *n ) {
            res.push(tmp);
        }
        if (lc < n){
            out.push(tmp + '(');
            left_c.push(lc + 1);
        }
        if ((len - lc) < lc){
            out.push(tmp + ')');
            left_c.push(lc);
        }
    }
    return res;
};

console.log(generateParenthesis(3));
console.log(generateParenthesis(4));
