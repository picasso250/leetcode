#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
    int titleToNumber(string s) {
        int col = 0;
        for (auto iter = s.begin(); iter != s.end(); ++iter) {
            char c = *iter;
            col = col * 26 + c - 'A' + 1;
        }
        return col;
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	string str("A");
	cout << s.titleToNumber(str) << endl;
	string aa("AA");
	cout << s.titleToNumber(aa) << endl;
	return 0;
}
