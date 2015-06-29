#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
	bool isPalindrome(std::string::iterator begin, std::string::iterator end) {
		string s(begin, end);
		cout<<"check '"<<s<<"'"<<endl;
		while (begin < end) {
			if (*begin != *end) {
				return false;
			}
			begin++;
			end--;
		}
		return true;
	}
    void insert0(vector<vector<string>> v, std::string::iterator begin, std::string::iterator end) {
		for (auto i : v) {
			string s(begin, end);
			i.insert(i.begin(), s);
		}
    }
    vector<vector<string>> _partition(std::string::iterator begin, std::string::iterator end) {
    	vector<vector<string>> v;
    	if (end - begin <= 2) {
    		string s(begin, end);
    		vector<string> vs;
    		vs.push_back(s);
    		v.push_back(vs);
    		return v;
    	}
    	for (auto i = begin+1; i < end-1; i++) {
    		if (this->isPalindrome(begin, i)) {
		    	vector<vector<string>> vi;
    			vi = this->_partition(i, end);
    			this->insert0(vi, begin, i);
    			for (auto i : vi) {
    				v.push_back(i);
    			}
    		}
    	}
    	return v;
    }
public:
    vector<vector<string>> partition(string s) {
        return _partition(s.begin(), s.end());
    }

};

int main(int argc, char const *argv[])
{
	Solution s;
	string str("aab");
	for (auto i : s.partition(str)) {
		for (auto j : i) {
			cout << j <<"\t";
		}
		cout <<endl;
	}
	return 0;
}
