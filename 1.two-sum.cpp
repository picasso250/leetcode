#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
		unordered_map<int,int> m;
        for (int n : nums) {
        	m[n] = 1;
		}
		for (int i = 0; i < nums.size(); ++i) {
			int n = nums[i];
        	int another = target - n;
        	auto it = m.find(another);
        	if (it != m.end())
        	{
        		for (int j = 0; j < nums.size(); ++j)
        		{
        			if (i != j && nums[j] == another)
        			{
		        		return {i+1, j+1};
        			}
        		}
        	}
		}

    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	vector<int> nums = {2, 7, 11, 15};
	int target = 9;
	auto v = s.twoSum(nums, target);
	cout<<v[0]<<" "<<v[1]<<endl;
	return 0;
}
