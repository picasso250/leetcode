#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

struct ListNode *makeList(int *vals, int size, int pos)
{
    struct ListNode *head = malloc(sizeof(struct ListNode) * size);
    struct ListNode *p = head;
    for (int i = 0; i < size; i++)
    {
        p->val = vals[i];
        p->next = head + i + 1;
    }
    if (pos == -1)
        p->next = NULL;
    else
        p->next = head + pos;
    return head;
}
bool hasCycle(struct ListNode *head)
{
    if (head == NULL)
        return false;
    struct ListNode *slow = head, *fast = head->next;
    for (; slow != NULL && fast != NULL && fast->next != NULL; slow = slow->next, fast = fast->next->next)
    {
        if (slow == head)
        {
            return true;
        }
    }
    return false;
}
int main(int argc, char const *argv[])
{

    int vals[] = {-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5};
    struct ListNode *head = makeList(vals, sizeof(vals) / sizeof(vals[0]), -1);
    bool result = hasCycle(head);
    printf("%d, ", result);
    return 0;
}
