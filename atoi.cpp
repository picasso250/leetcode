#include <iostream>

using namespace std;

int INT_MAX = 2147483647;
int INT_MIN = -2147483648;

class Solution {
public:
    int atoi(const char *str) {
        int n = 0;
        int sign = 1;
        while (*str == ' ') {
            str ++;
        }
        if (*str == '-') {
            sign = -1;
            str++;
        } else if (*str == '+') {
            sign = 1;
            str++;
        }
        
        while (*str) {
            // cout << "for " << *str << endl;
            if (!(*str >= '0' && *str <= '9'))
            {
                return n;
            }
            if (n > INT_MAX / 10)
            {
                return INT_MAX;
            }
            if (n < INT_MIN / 10)
            {
                return INT_MIN;
            }
            int end = (*str - '0');
            n = n * 10 + sign * end;
            // cout << "sign of user " << sign << ", of n = " << n << " " << this->_sign(n) << endl;
            if (n != 0 && this->_sign(n) != sign) {
                // cout << "overflow" << endl;
                return sign == 1 ? INT_MAX : INT_MIN;
            }
            str++;
        }
        return n;
    }
    int _sign(int n)
    {
        return n >= 0 ? 1 : -1;
    }
};

int main(int argc, char const *argv[])
{
    Solution s;
    cout << s.atoi("    10522545459") << endl;
    return 0;
}
