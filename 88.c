#include <stdio.h>
#include <string.h>
#include <stdlib.h>
void merge(int *nums1, int nums1Size, int m, int *nums2, int nums2Size, int n)
{
    int *nums1_ = malloc(m * sizeof(int));
    memcpy(nums1_, nums1, m * sizeof(int));
    int i, j, k;
    for (i = 0, j = 0, k = 0; i < m && j < n; k++)
    {
        if (nums1_[i] < nums2[j])
        {
            nums1[k] = nums1_[i++];
        }
        else
        {
            nums1[k] = nums2[j++];
        }
    }
    if (i == m)
    {
        // copy 2
        memcpy(nums1 + i + j, nums2 + j, sizeof(int) * (n - j));
    }
    else
    {
        memcpy(nums1 + i + j, nums1_ + i, sizeof(int) * (m - i));
    }
    free(nums1_);
}
int main(int argc, char const *argv[])
{

    int nums1[] = {2,0}, m = 1;
    int nums2[] = {1}, n = 1;
    merge(nums1,2,m,nums2,n,n);
    for (int i = 0; i < m+n; i++)
    {
        printf("%d, ", nums1[i]);
    }
    return 0;
}
