#include <iostream>
#include <vector>
#include <array>
#include <string>
#include <map>
#include <algorithm>

using namespace std;

class Solution {
public:
    int maxArea(vector<int>& height) {
        int max = 0;
        // > i or < i
        vector<array<int,2>> v;
        for (int i=0; i< height.size(); i++) {
            v.push_back(new array<int,2>(i,height[i]));
        }
        sort(v.begin(), v.end(), [](int *a,int *b){return a[1] - b[1];});
        for (int i=v.size()-1; i>=0;i++) {
            auto a=v[i];
            

        }
        return max;
    }
};

int main(int argc, char const *argv[])
{
    vector<int> v={1,2,3};
    Solution s;
    cout<<s.maxArea(v)<<endl;
    return 0;
}
