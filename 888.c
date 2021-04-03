#include <stdio.h>

long long arraySum(int *A, int ASize)
{
    long long s = 0;
    for (int i = 0; i < ASize; i++)
    {
        s += A[i];
    }
    return s;
}
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *fairCandySwap(int *A, int ASize, int *B, int BSize, int *returnSize)
{
    long long asum = arraySum(A, ASize), bsum = arraySum(B, BSize);
    int d = bsum - asum;
    // printf("d %d\n",d);
    *returnSize = 2;
    int *ret = malloc(2 * sizeof(int));
    for (int i = 0; i < ASize; i++)
    {
        for (int j = 0; j < BSize; j++)
        {
            // judge
            if (B[j] - A[i] == d / 2)
            {
                // printf("ok %d %d\n",i,j);
                ret[0] = A[i];
                ret[1] = B[j];
                return ret;
            }
        }
    }
    return ret;
}
int main(int argc, char const *argv[])
{
    int A[] = {1, 2, 5}, B[] = {2, 4};
    int ret[] = {0, 0};
    fairCandySwap(A, sizeof(A) / sizeof(A[0]), B, sizeof(B) / sizeof(B[0]), ret);
    printf("[%d,%d]\n", ret[0], ret[1]);
}