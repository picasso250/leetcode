#include <vector>
#include <string>
#include <map>
#include <iostream>
using namespace std;

class Solution
{
public:
    vector<string> findRepeatedDnaSequences(string s)
    {
        map<string, int> m;
        for (int i = 0; i + 10 < s.size(); i++)
        {
            string ss = s.substr(i, 10);
            if (m.find(ss) != m.end())
                continue;
            if (s.find(ss, i + 1) != s.npos)
            {
                // cout << "find " << ss << " again in " << s.find(ss, i + 1) << endl;
                m[(ss)] = 1;
            }
        }
        vector<string> ret;
        for (auto i = m.begin(); i != m.end(); i++)
        {
            ret.push_back(i->first);
        }
        return ret;
    }
};
int main(int argc, char const *argv[])
{
    Solution s;
    auto a = s.findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT");
    for (auto i = a.begin(); i != a.end(); i++)
    {
        cout << *i << endl;
    }
}