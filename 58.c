#include <stdio.h>
#include <string.h>
#include <stdlib.h>
int lengthOfLastWord(char *s)
{
    int ps = 0; // last position of space
    int i=0;
    for (i=0; s[i]; i++)
    {
        if (s[i] == ' ')
        {
            ps = i;
        }
    }
    if (ps==0 && s[ps] != ' ')
    {
        return i-ps;
    }
    return i - ps - 1;
}
int main(int argc, char const *argv[])
{

    char *s = "a";
    int ret = lengthOfLastWord(s);
    printf("%d\n", ret);
    return 0;
}
