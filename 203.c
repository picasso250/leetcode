#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "list.c"

static struct ListNode *rm(struct ListNode *prev, struct ListNode *nd)
{
    // nd!=NULL
    prev->next = nd->next;
    free(nd);
    return prev;
}
struct ListNode *removeElements(struct ListNode *head, int val)
{
    struct ListNode *p;
    struct ListNode *prev = head;
    for (p = head; p; p = p->next)
    {
        if (p->val == val)
        {
            if (prev == p)
            { // head
                head = p->next;
                free(p);
            }
            else
            {
                p = rm(prev, p);
            }
        }
        prev = p;
    }
    return head;
}

int main(int argc, char const *argv[])
{

    int a[] = {1, 2, 6, 3, 4, 5, 6};
    struct ListNode *lst = make_list(a, sizeof(a) / sizeof(a[0]));
    lst = removeElements(lst, 6);
    print_list(lst);
    return 0;
}
