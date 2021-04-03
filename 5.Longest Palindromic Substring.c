#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *longestPalindrome(char *s)
{
	// 找到的回文串的左端和右端
	int ri = 0, rj = 0;

	int sn=strlen(s);

	for (int i = 0; s[i]; i++)
	{
		for (int j = sn-1; j>=0; j--)
		{
			// 如要找的地方[i,j] 已经比我们所找到的还短，那么就没必要找了
			if (rj - ri + 1 >= j - i + 1)
			{
				continue;
			}

			// 从左右往中间找
			int p = 1; // 是否是回文串
			for (int a = i, b = j; a <= b; a++, b--)
			{
				if (s[a] != s[b])
				{
					p = 0;
					break;
				}
			}

			// 如果是回文
			if (p)
			{
				// 且长度大于我们之前所找到的最大长度
				if (j - i + 1 > rj - ri + 1)
				{
					// 则更新
					ri = i, rj = j;
				}
			}
		}
	}
	// 为其分配内存并返回
	int n = rj - ri + 1;
	char *ret = malloc(n + 1);
	strncpy(ret, s + ri, n);
	ret[n] = 0;
	return ret;
}

int main(int argc, char const *argv[])
{
	char *s = "1abcba2";
	char *ret = longestPalindrome(s);
	printf("%s\n", ret);
	return 0;
}
