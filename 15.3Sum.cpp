#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<vector<int> > threeSum(vector<int>& nums) {
    	sort(nums.begin(), nums.end());
        int size = nums.size();
        unordered_map<int,int> m;
        for (int i=0;i<size;i++) {
        	m[nums[i]] = i;
        }
        unordered_map<int,unordered_map<int,unordered_map<int,int> > > um;
        vector<vector<int>> ret;
        for (int i=0;i<size;i++) {
        	for (int j = i+1; j < size; ++j)
        	{
        		int c = 0 - nums[i] - nums[j];
        		auto it = m.find(c);
        		if (it!=m.end() && it->second != i && it->second != j)
        		{
        			int k = it->second;
        			vector<int> v;
        			v.push_back(nums[i]);
        			v.push_back(nums[j]);
        			v.push_back(nums[k]);
        			sort(v.begin(), v.end());
        			int ii = v[0];
        			int kk = v[2];
        			auto umm = um[ii][v[1]][kk]=1;
        			// auto ummm = &umm[v[1]];
        			// ummm[kk]=1;
        			for (auto iv : v) {
        				// cout<<iv<<",";
        			}
        			// cout<<endl;
        		}
        	}
        }
        // cout<<"===="<<endl;
        // cout<<"size "<<um.size()<<endl;
        for (auto i = um.begin(); i != um.end(); ++i)
        {
        	int ii=i->first;
        	// cout<<"["<<ii<<",";
        	auto umm = i->second;
	        for (auto j = umm.begin(); j != umm.end(); ++j) {
	        	int jj = j->first;
	        	// cout<<jj<<",";
	        	auto us = j->second;
		        for (auto k = us.begin(); k != us.end(); ++k) {
		        	auto v = new vector<int>;
		        	v->push_back(ii);
		        	v->push_back(jj);
		        	int kk = k->first;
		        	// cout<<ii<<","<<jj<<","<<kk<<endl;
		        	v->push_back(kk);
		        	ret.push_back(*v);
		        }
	        }
        }
        // cout<<endl;
        return ret;
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	vector<int> nums={-1,0,1,2,-1,-4};
	auto vec = s.threeSum(nums);
	for (auto v : vec)
	{
		for (auto i : v) {
			cout<<i<<",";
		}
		cout<<endl;
	}

	return 0;
}
