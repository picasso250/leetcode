#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int maxProfit(int *prices, int pricesSize)
{
    if (pricesSize == 0)
        return 0;
    // 动态规划 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
    int max = 0;
    int minPrice = prices[0];
    for (int i = 1; i < pricesSize; i++)
    {

        int maxToday = prices[i] - minPrice;
        if (maxToday > max)
            max = maxToday;

        // 前i-1天中的最小价格
        if (prices[i] < minPrice)
            minPrice = prices[i];
    }
    return max;
}

int main(int argc, char const *argv[])
{
    int s[] = {7, 1, 5, 3, 6, 4};
    int ret = maxProfit(s, sizeof(s) / sizeof(s[0]));
    printf("%d\n", ret);
    return 0;
}
