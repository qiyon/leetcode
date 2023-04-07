/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function (digits) {
    var out = [];
    var num_char_map = {
        '2': ['a', 'b', 'c'],
        '3': ['d', 'e', 'f'],
        '4': ['g','h','i'],
        '5': ['j','k','l'],
        '6': ['m','n','o'],
        '7': ['p','q','r','s'],
        '8': ['t','u','v'],
        '9': ['w','x','y','z'],
    };
    var pre_str, new_str;
    var loop_char_j;
    var i, ond_digit;
    var pre_out_length;
    for (i = 0; i < digits.length; i++) {
        one_digit = digits[i];
        if (num_char_map[one_digit]) {
            pre_out_length = out.length;
            do {
                pre_str = out.shift();
                for (loop_char_j in num_char_map[one_digit]) {
                    new_str = pre_str ? ('' + pre_str + num_char_map[one_digit][loop_char_j]) : ('' + num_char_map[one_digit][loop_char_j]);
                    out.push(new_str);
                }
                pre_out_length--;
            } while(pre_out_length > 0);
        }
    }
    return out;
};

//----leetcode submit end----
console.log(letterCombinations('23'));
console.log(letterCombinations('0213'));
console.log(letterCombinations('258'));
