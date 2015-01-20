#include <iostream>
#include <limits>

using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        if (x < 10) {
            return true;
        }
        int base = this->getPower10(x);
        // cout <<"base "<<base<<endl;
        int first, last;
        while (base > 0) {
            cout <<x << endl;
            first = x / base;
            last = x % 10;
            if (first == last) {
                x = x % base; // trim first
                x = x / 10; // trim last
                base /= 100;
            } else {
                return false;
            }
        }
        return true;
    }
    int getPower10 (int x) {
        int base = 10;
        // cout<<"x / base " <<x / base<<endl;
        while (10 <= (x / base)) {
            // cout <<"base -> "<< base<<endl;
            base *= 10;
        }
        return base;
    }
};

int main(int argc, char const *argv[])
{
    int x = -2147447412;
    Solution s;
    cout << x << " -> " << s.isPalindrome(x) <<endl;
    return 0;
}
