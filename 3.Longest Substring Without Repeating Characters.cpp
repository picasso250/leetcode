#include <iostream>
#include <vector>
#include <array>
#include <string>

using namespace std;

const int len=128;
class Solution {
public:
	void reset_map(int *map)
	{
		for (int i = 0; i < len; ++i)
		{
			map[i]=0;
		}
	}
	int lookup/*for*/(char c,/*in*/string s,/*begin*/int i)
	{
		for (int j=i-1;j>=0;j--)
        		{
        			if (s[j]==c){
        				return j+1;
        			}
        		}
	}
    int lengthOfLongestSubstring(string s) {
    	array<int,len> map;
    	map.fill(0);
    	int begin = 0;
    	int maxcandi;
    	int i=0;
    	vector<int> max;
        for (; i<s.size();i++)
        {
        	char c=s[i];
        	int offset = c;
        	if (map[offset]==0)
        	{
        		// cout<<c<<" Mark"<<endl;
	        	map[offset]=1;
        	} else {
        		int maxcandi = i-begin;
        		// cout<<c<<" Repeat "<<maxcandi<<endl;
        		max.push_back(maxcandi);
        		begin = this->lookup(c,s,i);
        		i = begin-1;
        		map.fill(0);
        	}
        }
        max.push_back(i-begin);
        // cout<<"max[";
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
};
int main(int argc, char const *argv[])
{
	Solution s;
	cout <<s.lengthOfLongestSubstring("abcabcbb")<<endl;
	cout <<s.lengthOfLongestSubstring("c")<<endl;
	cout <<s.lengthOfLongestSubstring("")<<endl;
	cout <<s.lengthOfLongestSubstring("aab")<<endl;
	return 0;
}
