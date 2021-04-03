struct ListNode
{
    int val;
    struct ListNode *next;
};
struct ListNode *make_list(int A[], int n)
{
    struct ListNode *head = NULL, *prev = NULL;
    for (int i = 0; i < n; ++i)
    {
        struct ListNode *node = malloc(sizeof(struct ListNode));
        node->val = A[i];
        node->next = NULL;
        if (head == NULL)
        {
            head = node;
        }
        if (prev)
        {
            prev->next = node;
        }
        if (prev == NULL)
        {
            prev = head;
        }
        else
        {
            prev = node;
        }
    }
    return head;
}
void print_list(struct ListNode *p)
{
    printf("> ");
    while (p)
    {
        printf("%d->", p->val);
        p = p->next;
    }
    printf("\n");
}