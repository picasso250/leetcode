#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
    string addBinary(string a, string b) {
    	cout << a << " + " << b << " = ?" << endl;
        string sum;
        auto x = a.end();
        auto y = b.end();
        char carry = '0';
        char s;
        for (--x, --y; x >= a.begin() || y >= b.begin(); x--, y--) {
        	char aa, bb;
        	bool is_a_end = x >= a.begin();
        	bool is_b_end = y >= b.begin();
        	aa = is_a_end ? *x : '0';
        	bb = is_b_end ? *y : '0';
            cout << "^" << carry << " ";
            if (aa == bb) {
                s = carry;
                carry = aa;
            } else {
                s = this->notBit(carry);
            }
            sum.insert(sum.begin(), s);
            cout << aa << " + " << bb << " = " << s << " ^" << carry << endl;
        }
        if (carry == '1')
        {
            sum.insert(sum.begin(), carry);
        }
        return sum;
    }
    
    char notBit(char b) {
        return b == '0' ? '1' : '0';
    }
};

int main(int argc, char const *argv[])
{
	Solution s;
	string a("11");
	string b("1");

	cout << s.addBinary(a, b) << endl;
	return 0;
}
