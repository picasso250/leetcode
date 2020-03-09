#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *countAndSay(int n)
{
    int len = 1;
    char *s = malloc(len + 1);
    strcpy(s, "1");
    if (n == 1)
        return s;
    for (int i = 2; i <= n; i++)
    {
        len *= 2;
        char *s1 = malloc(len + 1);
        int cnt = 0;
        int si = 0;
        char *p;
        for (p = s; *p; p++)
        {
            if (p > s && *p != *(p - 1))
            {
                s1[si++] = '0' + cnt;
                s1[si++] = *(p - 1);
                cnt = 1;
            }
            else
            {
                cnt++;
            }
        }
        s1[si++] = '0' + cnt;
        s1[si++] = *(p - 1);
        free(s);
        s = s1;
        len = si;
        s[len] = 0;
        s = realloc(s, len + 1);
    }
    return s;
}

int main(int argc, char const *argv[])
{
    int n = 6;
    char *ret = countAndSay(n);
    printf("%s\n", ret);
    return 0;
}
