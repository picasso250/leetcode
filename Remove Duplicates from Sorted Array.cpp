#include <iostream>
#include <limits>

using namespace std;

class Solution {
public:
    int removeDuplicates(int A[], int n) { // A=[1,1,1] n=3
        if (n <= 1) {
            return n;
        }
        int a = 0;
        int i = 1;
        while (i < n) { // i = 2
            // cout << "=== while loop === i(" <<i<<") < n("<<n<<")"<< endl;
            if (A[a] == A[i]) { // A[0] == A[2]
                // n--; // n=1
            } else {
                if (i - a > 1) {
                    A[a+1] = A[i];
                }
                a++;
            }
            i++;
            // cout << ">>>> loop end i("<<i<<") < n("<<n<<")"<< endl;
        }
        return a+1;
    }
};
int main(int argc, char const *argv[])
{
    Solution s;
    int a[] = {1, 1, 1};
    cout << s.removeDuplicates(a, 3);
    return 0;
}
