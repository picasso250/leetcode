#include <iostream>
#include <vector>
#include <string>
#include <set>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<string> generateParenthesis(int n) {
        auto res = generateParenthesisSet(n);
        vector<string> r(res.begin(), res.end());
        return r;
    }
    set<string> generateParenthesisSet(int n) {
        set<string> res;
        if (n==0) {
            return res;
        }
        if (n==1) {
            res.insert("()");
            return res;
        }
        
        auto last = generateParenthesisSet(n-1);
        for (auto s : last) {
            auto out = s+"()";
            res.insert(out);
            res.insert("()"+s);
            auto in = "("+s+")";
            res.insert(in);
        }
        for (int i = 2; i <= n/2+1; ++i)
        {
            int b = n-i;
            if (b==0)
            {
                break;
            }
            // cout<<"combine "<<i<<" and "<<b<<" ==== "<<n<<endl;
            auto left= generateParenthesisSet(i);
            auto right = generateParenthesisSet(b);
            for (auto s1 : left) {
                for (auto s2 : right) {
                    res.insert(s1+s2);
                }
            }
        }
        return res;
    }
};
int main(int argc, char const *argv[])
{
    Solution s;
    for (auto str : s.generateParenthesis(3)) {
        cout<<str<<endl;
    }
    cout<<"====="<<endl;
    int i = 0;
    for (auto str : s.generateParenthesis(4)) {
        i++;
        cout<<i<<". "<<str<<endl;
    }
    cout<<"====="<<endl;
    i = 0;
    for (auto str : s.generateParenthesis(5)) {
        i++;
        cout<<i<<". "<<str<<endl;
    }
    cout<<"====="<<endl;
    return 0;
}
