#include <iostream>
#include <map>

using namespace std;

  struct ListNode {
	  int val;
	  ListNode *next;
	  ListNode(int x) : val(x), next(NULL) {}
  };
class Solution {
public:
	ListNode *mergeTwoLists(ListNode *l1, ListNode *l2) {
		if (l1 == NULL)
		{
			return l2;
		}
		if (l2 == NULL)
		{
			return l1;
		}
		ListNode *head = NULL;
		// cout << "head is " << (l1->val < l2->val ? "1" : "2") << endl;
		ListNode *p = NULL;
		while (1) {
			// cout << "compare " << l1->val << " " << l2->val;
			if (l1->val < l2->val) {
				// cout << "\t +++ " << l1->val << endl;
				if (p != NULL)
				{
					p->next = l1;
				}
				p = l1; // p points to smaller
				if (head == NULL) {
					head = p;
				}
				l1 = l1->next; // smaller++
				if (l1 == NULL) {
					// cout << "reach end of 1" << endl;
					p->next = l2;
					break;
				}
			} else {
				// cout << "\t +++ " << l2->val << endl;
				if (p != NULL)
					p->next = l2;
				p = l2; // now p point to -7, but the quesiton is, who points to -7 ?
				if (head == NULL) {
					head = p;
				}
				l2 = l2->next;
				if (l2 == NULL) {
					// cout << "reach end of 2" << endl;
					p->next = l1;
					break;
				}
			}
		}
		return head;
	}
};

ListNode * make_list(int A[], int n)
{
	ListNode * head = NULL, *p = NULL;
	for (int i = 0; i < n; ++i)
	{
		auto node = new ListNode(A[i]);
		if (head == NULL)
		{
			head = node;
		}
		if (p == NULL)
		{
			p = head;
		} else {
			p->next = node;
			p = node;
		}
	}
	return head;
}
void print_list(ListNode * p)
{
	cout << "> ";
	while (p) {
		cout << p->val << " -> ";
		p = p->next;
	}
	cout << endl;
}
int main(int argc, char const *argv[])
{
	Solution s;
	int ia1[] = {-10,-10,-9,-4,1,6,6};
	auto l1 = make_list(ia1, 7);
	int ia2[] = {-7};
	auto l2 = make_list(ia2, 1);
	print_list(l1);
	print_list(l2);
	auto l = s.mergeTwoLists(l1, l2);
	print_list(l);
	return 0;
}
