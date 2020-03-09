#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int strStr(char *haystack, char *needle)
{
    if (strlen(needle) == 0)
        return 0;
    //  Sunday 算法
    int idx = 0;

    int n = strlen(needle);

    // 偏移表
    int shift[128];
    for (int i = 0; i < 128; i++)
    {
        shift[i] = n + 1;
    }
    for (int i = 0; i < n; i++)
    {
        shift[needle[i]] = n - i;
    }

    // 匹配
    for (; idx + n <= strlen(haystack);)
    {
        if (strncmp(haystack + idx, needle, n) == 0)
        {
            return idx;
        }
        else
        {
            int c = haystack[idx + n];
            if (shift[c] == n + 1)
            {
                idx += n;
            }
            else
            {
                idx += shift[c];
            }
        }
    }
    return -1;
}

int main(int argc, char const *argv[])
{
    "mississippi"
    "pi";
    char *haystack = "mississippi", *needle = "pi";
    int ret = strStr(haystack, needle);
    printf("%d\n", ret);
    return 0;
}
