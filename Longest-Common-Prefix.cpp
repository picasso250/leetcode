#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
    string longestCommonPrefix(vector<string> &strs) {
        string cp;
        if (strs.size() == 0) {
            return cp;
        }
        auto si = strs.begin();
        cp = *si;
        for (si++; si < strs.end(); si++) {
            auto ps = (*si).begin();
            for (auto p = cp.begin(); p < cp.end(); p++) {
                if (ps == (*si).end()) {
                    cp.resize(p - cp.begin());
                    break;
                }
                if (*ps == *p) {
                    ps++;
                    continue;
                } else {
                    cp.resize(p - cp.begin());
                    break;
                }
            }
        }
        return cp;
    }
};

int main(int argc, char const *argv[])
{
    vector<string> strs{
        "hello",
        "hezzzzzz"
    };
    Solution s;
    cout << s.longestCommonPrefix(strs) << endl;
    return 0;
}
