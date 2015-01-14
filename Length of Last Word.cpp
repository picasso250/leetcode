#include <iostream>
#include <string>
#include <map>

using namespace std;

class Solution {
public:
    int lengthOfLastWord(const char *s) { // s="a"
        auto begin = s;
        auto end = s;
        bool in_space = false;
        for (; *s; s++) {     // s => 'a'
            // assert begin always point to the last word's first char
            if (*s == ' ') {  // false
                if (!in_space)
                {
                    end = s;
                }
                in_space = true; // in_space=true
            } else {
                if (in_space) {  // false
                    begin = s;     // begin => 
                }
                in_space = false; // in_space=false
            }
            // assert in_space is true if and only if char is ' '
        }
        if (!in_space) {
            end = s;
        }
        return end - begin;
    }
};

int main(int argc, char const *argv[])
{
    map<string, int> table{
        {" ", 0},
        {"a", 1}
    };
    Solution s;
    for (auto i: table) {
        auto len = s.lengthOfLastWord(i.first.c_str());
        cout << "\"" << i.first << "\" ==> " << len << endl;
        if (len != i.second) {
            cout << "FAIL!!! Expect " << i.second << endl;
        }
    }
    return 0;
}
