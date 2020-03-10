#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

int singleNumber(int *nums, int numsSize)
{
    int r;
    for (int i = 0; i < numsSize; ++i)
    {
        if (i)
        {
            r ^= nums[i];
        }
        else
        {
            r = nums[i];
        }
    }
    return r;
}
int main(int argc, char const *argv[])
{

    int vals[] = {2,2,1};
    bool result = singleNumber(vals, sizeof(vals)/sizeof(vals[0]));
    printf("%d, ", result);
    return 0;
}
