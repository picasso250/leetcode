#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
void print_array(int *vals, int size)
{
    for (int i = 0; i < size; i++)
    {
        printf("%d ", vals[i]);
    }
}

void reverse(int *nums, int begin, int end)
{
    for (int i = begin, j = end - 1; i < j; i++, j--)
    {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }
}
void rotate(int *nums, int numsSize, int k)
{
    k = k % numsSize;
    if (k == 0)
        return;
    reverse(nums, 0, numsSize);
    reverse(nums, 0, k);
    reverse(nums, k, numsSize);
}
int main(int argc, char const *argv[])
{

    int vals[] = {1, 2, 3, 4, 5, 6};
    int k = 2;
    rotate(vals, sizeof(vals) / sizeof(vals[0]), k);

    print_array(vals, sizeof(vals) / sizeof(vals[0]));
    return 0;
}
