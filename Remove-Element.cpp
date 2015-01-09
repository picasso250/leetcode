#include <iostream>
#include <map>

using namespace std;

class Solution {
public:
    int removeElement(int A[], int n, int elem) {
        int skip = 0;
        for (int i = 0; i < n; i++) {
            int a = A[i];
            // cout << "for A[" << i << "] = " << a << endl;
            if (skip > 0 && i - skip >= 0) {
                // cout << "move A[" << i << "] to A[" << i - skip << "]" << endl;
                A[i - skip] = A[i];
            }
            if (elem == a) {
                skip ++;
            }
        }
        return n - skip;
    }
};

void print_array(int A[], int n)
{
    for (int i = 0; i < n; ++i)
    {
        cout << " " << A[i];
    }
    cout << endl;
}
int main(int argc, char const *argv[])
{
    int A[] = {4,5};
    Solution s;
    int len = s.removeElement(A, 2, 5);
    cout << "len = " << len << endl;
    print_array(A, len);
    return 0;
}
