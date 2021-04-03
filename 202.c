#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
void print_array(int *vals, int size)
{
    for (int i = 0; i < size; i++)
    {
        printf("%d ", vals[i]);
    }
}

void reverse(char *nums, int begin, int end)
{
    for (int i = begin, j = end - 1; i < j; i++, j--)
    {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }
}
char *itoa(int n, char *buf, int radix)
{
    int i;
    for (i = 0; n != 0; i++)
    {
        int a = n % radix;
        buf[i] = (char)(a) + '0';
        n /= radix;
    }
    buf[i] = 0;
    reverse(buf, 0, i);
    return buf;
}
int ss(int n)
{
    if (n == 0)
        return 0;
    char buf[64];
    itoa(n, buf, 10);
    int s = 0;
    for (char *p = buf; *p != '\0'; p++)
    {
        int a = (*p) - '0';
        s += a * a;
    }
    return s;
}
int *stack = NULL;
int cap = 0;
int size = 0;
void push(int n)
{
    if (size == cap)
    {
        cap = (cap + 1) * 2;
        stack = realloc(stack, sizeof(int) * cap);
    }
    stack[size++] = n;
}
bool exists(int n)
{
    for (int i = 0; i < size; i++)
    {
        if (stack[i] == n)
        {
            return true;
        }
    }
    return false;
}
bool isHappy(int n)
{
    if (n == 0)
        return false;
    if (n == 1)
        return true;
    push(n);
    for (;;)
    {
        int s = ss(n);
        if (s == 1)
            return true;
        if (exists(s))
        {
            return false;
        }
        push(s);
        n = s;
    }
}
int main(int argc, char const *argv[])
{
    printf("%d\n", isHappy(13));
    return 0;
}
