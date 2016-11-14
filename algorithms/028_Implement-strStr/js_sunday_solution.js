var strStr = function (haystack, needle) {
    var end_cmp, sub_begin, i;
    if (needle === "") return 0;
    var t_len = haystack.length;
    var p_len = needle.length;
    if (t_len < p_len) return -1;
    sub_begin = 0;
    while (true) {
        for(i = 0; i < p_len; i++) {
            if (needle[i] != haystack[sub_begin + i]) {
                end_cmp = sub_begin + p_len;
                break;
            }
        }
        if (i == p_len) return sub_begin;
        while(end_cmp < t_len) {
            for (i = p_len -1; i >= 0; i--) {
                if (needle[i] == haystack[end_cmp]){
                    sub_begin = sub_begin + (p_len - i);
                    break;
                }
            }
            if (i == -1) {
                end_cmp++;
                continue;
            }else {
                break;
            }
        }
        if (end_cmp >= t_len) break;
    }
    return -1;
};


console.log(strStr('abcdaabcd','aab'));
console.log(strStr('',''));
console.log(strStr('abcdaabcd',''));
console.log(strStr('abcdaabcd','xxx'));
console.log(strStr('xxx','xxx'));
console.log(strStr('xxxxxx','xxx'));
console.log(strStr('abcxxx sss xxx','xxx'));
console.log(strStr('abcxxx sss xxx','s x'));
