#include <iostream>
#include <vector>
#include <map>

using namespace std;

class Solution {
public:
    int majorityElement(vector<int> &num) {
        map<int,int> info;
        for (int i : num) {
            auto cnt = info.find(i);
            if (cnt == info.end()) {
                info[i] = 1;
                continue;
            }
            info[i] += 1;
        }
        for (auto it : info) {
            int count = it.second;
            if (count > num.size() / 2) {
                return it.first;
            }
        }
        
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	vector<int> num{1, 2, 3, 3, 3};
	cout << s.majorityElement(num) << endl;
	return 0;
}
