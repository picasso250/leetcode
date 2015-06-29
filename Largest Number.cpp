#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
public:
    string largestNumber(vector<int> &num) {
        string s = this->_largestNumber(num);
        return this->ltrim0(s);
    }
    string _largestNumber(vector<int> &num) {
        string s;
        if (num.empty()) {
            return s;
        }
        string max;
        max = this->max(num);
        s = max + this->_largestNumber(num);
        // cout<<"largestNumber: "<<s<<endl;
        return s;
    }
    string ltrim0(string s)
    {
        string z("0");
        if (this->all0(s)) {
            return z;
        }
        return s;
    }
    bool all0(string s) {
        for (auto c : s) {
            if (c != '0') {
                return false;
            }
        }
        return true;
    }
    string max(vector<int> &num) {
        // cout<< "=== max === size="<<num.size()<<endl;
        string max;
        auto mp = num.begin();
        for (auto i = num.begin(); i < num.end(); i++) {
            auto s = this->to_string(*i);
            if (max.empty() || this->cmp(s, max) > 0) {
                max = s;
                mp = i;
            }
        }
        // cout << "max <<< '"<<max<<"'"<<endl;
        num.erase(mp);
        return max;
    }
    string to_string(int n) {
        string s;
        if (n == 0)
        {
            s += '0';
            return s;
        }
        while (n > 0) {
            char c = n % 10 + '0';
            s.insert(s.begin(), c);
            n /= 10;
        }
        return s;
    }
    int cmp(string a, string b) {
        // cout << "compare a='"<<a<<"', b='"<<b<<"'"<<endl;
        int i;
        for (i = 0; i < a.size() && i < b.size(); i++) {
            if (a[i] != b[i]) {
                return a[i] - b[i];
            }
        }
        if (i == a.size() && i == b.size()) {
            return 0;
        }
        int begin = a[i-1];
        // cout << "begin="<<(begin-'0')<<endl;
        int k = 0;
        if (a.size() < b.size()) {
            for (int j = i; j < b.size(); j++, k++) {
                // cout<<"compare "<<begin-'0'<<" "<<c[j]-'0'<<endl;
                if (k >= a.size()) {
                    return 0;
                }
                if (a[i] != b[j]) {
                    return a[i] - b[j];
                }
            }
        } else {
            for (int j = i; j < a.size(); j++, k++) {
                // cout<<"compare "<<begin-'0'<<" "<<a[j]-'0'<<endl;
                if (k >= b.size()) {
                    return 0;
                }
                if (b[i] != a[j]) {
                    return a[j] - b[i];
                }
            }
        }
        return 0;
    }

};

int main(int argc, char const *argv[])
{
    Solution s;
    vector<int> v{824,938,1399,5607,6973,5703,9609,4398,8247};
    cout << s.largestNumber(v)<<endl;
    return 0;
}
