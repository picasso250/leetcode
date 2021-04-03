#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

struct TreeNode
{
    int val;
    struct TreeNode *left;
    struct TreeNode *right;
};

bool isSameTree(struct TreeNode *p, struct TreeNode *q)
{
    if (p == NULL && q == NULL)
        return true;
    if (p == NULL || q == NULL)
        return false;
    if (p->val != q->val)
        return false;
    return isSameTree(p->left, q->left) && isSameTree(p->right, q->right);
}

int main(int argc, char const *argv[])
{

    int nums1[] = {2, 0}, m = 1;
    int nums2[] = {1}, n = 1;
    merge(nums1, 2, m, nums2, n, n);
    for (int i = 0; i < m + n; i++)
    {
        printf("%d, ", nums1[i]);
    }
    return 0;
}
