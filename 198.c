#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

int rob(int *nums, int numsSize)
{
    /*
    执行用时 :0 ms, 在所有 C 提交中击败了100.00% 的用户
    内存消耗 :5.3 MB, 在所有 C 提交中击败了100.00%的用户
    */

    // max ( 1 + rob([2:]), rob([1:]))

    int max0 = 0, max1 = 0; // max1, max0, ...
    for (int i = numsSize - 1; i >= 0; i--)
    {
        if (i == numsSize - 1)
        {
            max1 = max0 = nums[i];
        }
        else if (i == numsSize - 2)
        {
            if (nums[i] > max0)
            {
                max1 = nums[i];
            }
        }
        else
        {
            int a = nums[i] + max0;
            max0 = max1;
            max1 = a > max1 ? a : max1;
        }
    }

    return max1;
}
int main(int argc, char const *argv[])
{

    int vals[] = {2, 7, 9, 3, 1};
    int result = rob(vals, sizeof(vals) / sizeof(vals[0]));
    printf("%d, ", result);
    return 0;
}
