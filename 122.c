#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize <= 1)
        return 0;
    // 峰谷法
    int sum = 0;
    for (int i = 1; i < pricesSize; i++)
    {
        if (prices[i] > prices[i - 1])
            sum += prices[i] - prices[i - 1];
    }
    return sum;
}

int main(int argc, char const *argv[])
{
    int s[] = {7, 1, 5, 3, 6, 4};
    int ret = maxProfit(s, sizeof(s) / sizeof(s[0]));
    printf("%d\n", ret);
    return 0;
}
