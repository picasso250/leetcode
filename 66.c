#include <stdio.h>
#include <string.h>
#include <stdlib.h>
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *plusOne(int *digits, int digitsSize, int *returnSize)
{
    int *result = malloc((digitsSize) * sizeof(int));
    // result[0] = 0;
    memcpy(result, digits, digitsSize * sizeof(int));
    int j = 0; // 进位
    for (int i = digitsSize - 1; i >= 0; i--)
    {
        int r = digits[i] + 1;
        if (r >= 10)
        {
            result[i] = r - 10;
            j = 1;
        }
        else
        {
            result[i] = r;
            j = 0;
            break;
        }
    }
    *returnSize = digitsSize;
    if (j == 1)
    {
        (*returnSize)++;
        result = realloc(result, *returnSize * sizeof(int));
        memmove(result + 1, result, digitsSize * sizeof(int));
        result[0] = 1;
    }
    return result;
}
int main(int argc, char const *argv[])
{

    int a[] = {9, 9, 9};
    int size;
    int *ret = plusOne(a, 3, &size);
    for (int i = 0; i < size; i++)
    {
        printf("%d", ret[i]);
    }
    return 0;
}
