#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    bool isMatch(const char *s, const char *p) {
        // cout << s << " ~ /" << p << "/" << endl;
        auto i = s; // back trace pos, will reach
        auto p_start = p;
        auto s_start = s;
        bool is_match;
        auto olds = s;
        int bi = 0;
        while (true) {
            // cout << "2 left " << s << " " << p << endl;
            // assert p != '*'
            // cout << *s << " ==== ";
            if (*(p + 1) == '*') {
                // .*
                // cout << "reg " << *p << "*" << endl;
                olds = s;
                // cout << "olds " << (olds - s_start) << endl;
                is_match = false;
                while (*s && this->_match_char(s, p)) {
                    // cout << "eat " << *s << endl;
                    s++;
                    is_match = true;
                    i = olds;
                }
                // cout << "back trace pos " << (i - s_start) << endl;
                p++;
            } else {
                // .
                // cout << "reg " << *p << endl;
                if (*s) {
                    if (!this->_match_char(s, p)) {
                        return false;
                    }
                    s++;
                    i = s;
                } else {
                    // find last .*
                    for (; p - bi >= p_start; ++bi)
                    {
                        if (*(p - bi) == '*')
                        {
                            bi++;
                            break;
                        }
                    }
                    cout << "\tmatch " << *p << "fail" << endl;
                    while (true) {
                        if ((p - bi) >= p_start && *(p - bi + 1) == '*') {
                            cout << "\tback trace " << *(p - bi) << *(p - bi + 1) << endl;
                            if (this->_match_char(s - 1, p - bi))
                            {
                                s--; // back trace string
                                cout << "back to str[" <<(s- s_start)<< "] = " << *s << endl;
                                break;
                            } else {
                                bi += 2; // back trace regex
                            }
                        } else {
                            return false;
                        }
                    }
                    continue;
                }
            }
            p++;

            if (*p == '\0')
            {
                // cout << "reg end" << endl;
                // it must end
                return *s == '\0';
            }
        }
    }
    bool _match_char(const char *cp, const char *rp)
    {
        return *rp == '.' || *cp == *rp;
    }
};

int main(int argc, char const *argv[])
{
    vector<vector<string>> table({
        // Input Output Expected
        {"aa", "a", "false"},
        {"aa", "aa", "true"},
        {"aaa", "aa", "false"},
        {"aa", "a*", "true"},
        {"aa", ".*", "true"},
        {"ab", ".*", "true"},
        {"aab", "c*a*b", "true"},
        {"aaa", "a*a", "true"},
        {"abcd", "d*", "false"},
        {"aaa", "ab*a*c*a", "true"},
        {"a", "ab*", "true"},
        {"a", "ab*a", "false"},
        {"ab", ".*..", "true"},
        {"a", ".*..a*", "false"},
    });
    // auto table = 
    Solution s;
    for (int i = 0; i < table.size(); ++i)
    {
        auto str = table[i][0].c_str();
        auto reg = table[i][1].c_str();
        auto isMatch = s.isMatch(str, reg);
        cout << "isMatch(\"" << str << "\",\"" << reg << "\") → " 
             << isMatch << endl;
        auto expect = table[i][2];
        if ((isMatch && expect == "true") || !isMatch && expect == "false")
        {
            /* code */
        } else {
            cout << "\tFAIL!!!" << endl;
        }
    }
    return 0;
}
