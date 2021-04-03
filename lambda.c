#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

/*
x
x y
\x.body
*/
typedef struct _lambda
{
    int type;
    union {
        char name;
        struct
        {
            char param;
            struct _lambda *body;
        } func;
        struct
        {
            struct _lambda *func;
            struct _lambda *arg;
        } apply;
    };
} lambda;
#define NAME 1
#define FUNC 2
#define APPLY 3

lambda *make_name(char name)
{
    lambda *lmd = malloc(sizeof(lambda));
    lmd->type = NAME;
    lmd->name = name;
    return lmd;
}

lambda *make_func(char param,
                  lambda *body)
{
    lambda *lmd = malloc(sizeof(lambda));
    lmd->type = FUNC;
    lmd->func.param = param;
    lmd->func.body = body;
    return lmd;
}

lambda *make_apply(lambda *func,
                   lambda *arg)
{
    lambda *lmd = malloc(sizeof(lambda));
    lmd->type = APPLY;
    lmd->apply.func = func;
    lmd->apply.arg = arg;
    return lmd;
}

lambda *come(lambda *cur, lambda *with)
{
    if (with == NULL)
    {
        printf("empty ()\n");
        exit(-1);
    }
    if (cur == NULL)
    {
        return with; // a
    }
    return make_apply(cur, with);
    // switch (cur->type)
    // {
    // case NAME: // a b
    //     return make_apply(cur, with);
    //     break;
    // case FUNC: // \x.a
    //     return make_apply()
    //     cur->func.body = come(cur->func.body, with);
    //     break;
    // case APPLY: // a b c
    //     return make_apply(cur, with);
    //     break;
    // default:
    //     printf("impossible\n");
    //     exit(1);
    // }
}
bool in_body(char *str, int *stack, int sp)
{
    for (int i = 0; i < sp; i++)
    {
        if (str[stack[i] + 1] == '\\')
        {
            return true;
        }
    }
    return false;
}
lambda *parse(char *str, int size)
{

    // 合法性检查 λx.
    // 省略

    // 从左往右依次
    int state = 0;    // 状态 1 body中;
    int stack[200];   // 左括号的坐标
    int sp = 0;       // stack pointer

    lambda *lmd[200]; // lambda pointer stack
    int lp = 0;

    lambda *cur = NULL; // current or body

    int ss[200];        // stack of state
    int ssp = 0;        //

    // in body ?
    // (\x.
    // ^
    for (int i = 0; i < size; ++i)
    {
        char c = str[i];

        if (c == '\\')
        {
            ss[ssp++] = state;
            state = 1;
            if (i + 3 >= size)
            {
                printf("%d expect param and body\n");
                exit(1);
            }
            // i+1 parameter
            i++;
            printf("lambda %c\n", str[i]);
            char param = str[i];
            // i+2 .
            i++;
            if (str[i] != '.')
            {
                printf("expect .\n");
                exit(1);
            }
            lambda *_ = make_func(param, NULL);
            lmd[lp++] = come(cur, _);
            cur = NULL;
        }
        else if (c == '(')
        {
            stack[sp++] = i; // push
            if (cur)
            {
                lmd[lp++] = cur;
                cur = NULL;
            }
        }
        else if (c == ')')
        {
            if (sp == 0)
            {
                printf("unmatched ) %d\n", i);
                exit(1);
            }
            sp--; // pop
            // state = 0;
            // int lpos = lmd[lp - 1];
            if (str[stack[sp] + 1] == '\\')
            {
                printf("\\%c complete\n", str[stack[sp] + 2]);
                // end body
                if (lp == 0)
                {
                    printf("body need lambda\n");
                    exit(1);
                }
                lambda *f = lmd[--lp]; // pop
                if (f->type != FUNC)
                {
                    printf("must be func type\n");
                    exit(1);
                }
                f->func.body = cur;
                cur = f;
                state = ss[--ssp];
            }
            else
            {
                // a (b c)
                lambda *a = lp > 0 ? lmd[--lp] : NULL;
                cur = come(a, cur);
            }
        }
        else if (c == ' ')
        {
        }
        else
        {
            printf("must be alpha %c\n", str[i]);
            cur = come(cur, make_name(c));
        }
    }
    while (lp > 0)
    {
        if (cur == NULL)
        {
            printf("body empty\n");
            exit(1);
        }
        lmd[lp - 1]->func.body = cur;
        cur = lmd[--lp];
    }
    return cur;
}
void print_lambda(lambda *lmd)
{
    if (lmd == NULL)
    {
        printf("NULL");
        return;
    }
    switch (lmd->type)
    {
    case NAME:
        printf("%c", lmd->name);
        break;
    case FUNC:
        printf("(\\%c.", lmd->func.param);
        print_lambda(lmd->func.body);
        printf(")");
        break;
    case APPLY:
        printf("(");
        print_lambda(lmd->apply.func);
        printf(" ");
        print_lambda(lmd->apply.arg);
        printf(")");
        break;
    default:
        printf("impossible\n");
        exit(1);
    }
}
char *sub(char *str, int body_begin, int body_end, char param, char arg)
{
    int n = body_end - body_begin;
    char *res = malloc(n);
    strncpy(res, str + body_begin, n);
    int stack[200];
    int sp = 0;
    for (int i = body_begin; i < body_end; i++)
    {
        char c = str[i];
        if (c == '\\')
        {
        }
        else if (c == '(')
        {
        }
        else if (c == ')')
        {
        }
    }
}
int main(int argc, char const *argv[])
{
    // read a string
    // char str[2000];
    // scanf("%s", &str);
    char *str = "\\f.(\\u.u u)(\\x.f(x x))";
    // char *str = "(\\x.x)y";
    // char *str = "a";
    // char *str = "a b";
    // char *str = "a b c";
    // char *str = "a (b c)";
    // char *str = "\\x.y";
    // char *str = "(\\x.y)z";
    // char *str = "\\x.\\y.z";
    lambda *lmd = parse(str, strlen(str));
    print_lambda(lmd);
    return 0;
}
