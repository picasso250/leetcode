#include <iostream>

using namespace std;

class Solution {
public:
    int maxSubArray(int A[], int n) {
        return this->_maxSubArray(A, 0, n);
    }
    int _maxSubArray(int A[], int begin, int end) {
        int n = end - begin;
        if (n == 0) {
            return 0;
        }
        if (n == 1)
        {
            return A[begin];
        }
        if (n == 2)
        {
            return this->_max3(A[begin], A[begin + 1], A[begin] + A[begin + 1]);
        }
        int mid = (begin + end) / 2 + 1; // mid belong to right
        // cout << "mid " << mid << endl;
        int left = this->_maxSubArray(A, begin, mid);
        int right = this->_maxSubArray(A, mid, end);
        int left_mid = this->_fixMaxSum(A, mid - 1, begin - 1, -1);
        int right_mid = this->_fixMaxSum(A, mid, end, 1);
        int mid_sum = left_mid + right_mid;
        return this->_max3(left, right, mid_sum);
    }
    int _max3(int a, int b, int c)
    {
        int tmp;
        tmp = a > b ? a : b;
        return tmp > c ? tmp : c;
    }
    int _fixMaxSum(int A[], int fix, int end, int step)
    {
        int sum, max;
        max = sum = A[fix];
        for (int i = fix + step; i != end; i += step)
        {
            int elem = A[i];
            sum += elem;
            if (max < sum)
            {
                max = sum;
            }
        }
        return max;
    }
};

int main(int argc, char const *argv[])
{
    Solution s;
    int A [] = {-2,1,-3,4,-1,2,1,-5,4};
    cout << s.maxSubArray(A, 9) << endl;
    return 0;
}
