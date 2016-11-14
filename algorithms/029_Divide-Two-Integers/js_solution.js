/**
 * @param {number} dividend
 * @param {number} divisor
 * @return {number}
 */
var divide = function(dividend, divisor) {
    var MAX_INT = 2147483647;
    var MIN_INT = -2147483648;
    var is_posi = (dividend >= 0 ) ? true : false;
    is_posi = (divisor >= 0) ? is_posi : !is_posi;
    if (divisor === 0) return is_posi ? MAX_INT : MIN_INT;
    if (divisor === 1) return dividend;
    var ans = 0;
    var total = Math.abs(dividend), part = Math.abs(divisor);
    var tmp, addon;
    while (part <= total) {
      addon = 1;
      tmp = part;
      while ((tmp<<1) < total && (tmp<<1) > (tmp)) {
        tmp = tmp<<1;
        addon = addon<<1;
      }
      total = total - tmp;
      ans = ans +addon;
    }
    return is_posi ? (ans > MAX_INT ? MAX_INT : ans) : -ans;
};


console.log(divide(-2147483648, 0));
console.log(divide(354123, 29));
console.log(divide(-1, 0));
