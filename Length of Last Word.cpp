#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    int lengthOfLastWord(const char *s) { // s=" "
        auto begin = s;
        auto end = s;
        bool in_space = false;
        for (; *s; s++) {     // s => ' '
            // assert begin always point to the last word's first char
            if (*s == ' ') {  // true
                if (!in_space)
                {
                    end = s;
                }
                in_space = true; // in_space=true
            } else {
                if (in_space) {  // true
                    begin = s;     // begin => 'W'
                }
                in_space = false; // in_space=false
            }
            // assert in_space is true if and only if char is ' '
        }
        return end - begin;
    }
};
