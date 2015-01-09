#include <iostream>

using namespace std;

class Solution {
public:
    int searchInsert(int A[], int n, int target) {
        for (int i = 0; i < n; i++) {
            int a = A[i];
            if (a == target) {
                return i;
            } else if (a < target) {
                continue;
            } else {
                return i;
            }
        }
    }
};

int main(int argc, char const *argv[])
{
    int A[] = {1,3,5,6};
    Solution s;
    int p = s.searchInsert(A, sizeof A, 5);
    cout << p << endl;
    int p2 = s.searchInsert(A, sizeof A, 2);
    cout << p2 << endl;
    int p7 = s.searchInsert(A, sizeof A, 7);
    cout << p7 << endl;
    int p0 = s.searchInsert(A, sizeof A, 0);
    cout << p0 << endl;
    return 0;
}
