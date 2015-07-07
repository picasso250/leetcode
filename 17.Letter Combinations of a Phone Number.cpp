#include <iostream>
#include <vector>
#include <string>
#include <map>

using namespace std;

class Solution {
public:
    vector<string> letterCombinations(string digits) {
    	vector<string> res;
    	map<char, string> m;
    	m['2']="abc";
    	m['3']="def";
    	m['4']="ghi";
    	m['5']="jkl";
    	m['6']="mno";
    	m['7']="pqrs";
    	m['8']="tuv";
    	m['9']="wxyz";
        for (int i = 0; i < digits.size(); ++i)
        {
        	char c = digits[i];
        	res = addLayer(res, m[c]);
        }
        return res;
    }
    vector<string> addLayer(vector<string> v, string m) {
    	vector<string> res;
    	if (v.empty()) {
    		v.push_back("");
    	}
    	for (string s : v) {
		    for (char ch : m) {
		    	string ss = s+ch;
		    	res.push_back(ss);
	    	}
    	}
    	return res;
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	auto v = s.letterCombinations("23");
	for (auto str : v) {
		cout<<str<<endl;
	}
	return 0;
}
