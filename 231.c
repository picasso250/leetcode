#include <stdbool.h>
#include <stdio.h>

bool isPowerOfTwo(int n)
{
    if (n == 0)
        return false;
    if (n == 1)
    {
        return true;
    }
    else
    {
        if (n & 1 == 0)
        {
            // printf("n %d\n", n);
            return isPowerOfTwo(n >> 1);
        }
        else
        {
            return false;
        }
    }
}
int main()
{
    printf("1\t%d\n",isPowerOfTwo(1));
    printf("16\t%d\n",isPowerOfTwo(16));
    printf("218\t%d\n", isPowerOfTwo(218));
}