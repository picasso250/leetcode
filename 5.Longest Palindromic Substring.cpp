#include <iostream>
#include <vector>
#include <array>
#include <string>
#include <map>

using namespace std;

class Solution {
public:
	int v_max(vector<int>max)
	{
		int imax = 0;
        for (int m : max) {
        	if (m > imax) {
        		imax = m;
        	}
        	// cout<<m<<",";
        }
        // cout<<"]"<<endl;
        return imax;
	}
	int checkLR(string s, int left, int right, int len, map<int,int[2]> &max)
	{
		// cout<<left<<","<<right<<endl;
		while (left >= 0 && right < s.size() && s[left]==s[right]) {
			left--;
			right++;
			len+=2;
		}
		max[len][0] = left;
		max[len][1] = right;
		// cout<<"=>"<<len<<endl;
		return len;
	}
    string longestPalindrome(string s) {
    	map<int,int[2]> max;
    	int left, right;
    	int len;
        for (int i = 0; i < s.size(); i++) {
        	right = i+1;
        	if (right<s.size() && s[i]==s[right]){
        		len = checkLR(s, i, right, 0, max);
        	}
        	right=i+2;
        	if (right<s.size() && s[i]==s[right]) {
        		len = checkLR(s, i, right, 1, max);
        	}
        }
        int imax = 0;
        int *p = NULL;
        // cout<<max.size()<<endl;
        for (auto it = max.begin(); it != max.end(); ++it) {
        	// cout<<it->first<<","<<endl;
        	if (it->first > imax) {
        		imax = it->first;
        		p = it->second;
        	}
        }
        if (p == NULL)
        {
        	// cout<<"p NULL"<<endl;
        	return s;
        }
        return s.substr(p[0]+1, p[1]-p[0]-1);
    }
};
int main(int argc, char const *argv[])
{
	Solution sn;
	string s("1abcba2");
	cout<<"\""<<sn.longestPalindrome(s)<<'"'<<endl;
	cout<<"\""<<sn.longestPalindrome("a")<<'"'<<endl;
	return 0;
}
