#include <iostream>
#include <vector>
#include <array>
#include <string>

using namespace std;

class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    	long s=0;
    	int c=0;
        for (int i : nums1) {
        	s += i;
        	c++;
        }
        for (int i : nums2) {
        	s += i;
        	c++;
        }
        return s * 1.0 / c;
    }
};
int main(int argc, char const *argv[])
{
	Solution s;
	vector<int> nums1={};
	vector<int> nums2={2,3};
	cout<<s.findMedianSortedArrays(nums1,nums2)<<endl;
	vector<int> nums1_={4,5,6,8,9};
	vector<int> nums2_={};
	cout<<s.findMedianSortedArrays(nums1_,nums2_)<<endl;
	return 0;
}
