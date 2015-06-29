#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>

using namespace std;

  struct ListNode {
      int val;
      ListNode *next;
      ListNode(int x) : val(x), next(NULL) {}
  };
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int carry = 0;
        ListNode *root = NULL;
        ListNode *p = NULL;
        while(l1 || l2 || carry == 1) {
        	if (l1==NULL && l2==NULL && carry==0)
        	{
        		break;
        	}
        	int v1=0;
        	int v2=0;
        	if (l1)
        	{
        		v1 = l1->val;
        	}
        	if (l2)
        	{
        		v2=l2->val;
        	}
            int s = v1 + v2 + carry;
            if (s > 9) {
                carry = 1;
                s -= 10;
            } else {
            	carry = 0;
            }
            // cout<<"ADD "<<v1<<"+"<<v2<<"="<<s<<"^"<<carry<<endl;
            auto n = new ListNode(s);
            if (p == NULL) {
                p = n;
                root = n;
            } else {
                p->next = n;
                p = n;
            }
            if (l1)
            {
	            l1 = l1->next;
            }
            if (l2)
            {
	            l2 = l2->next;
            }
        }
        return root;
    }
};
ListNode * build_list(vector<int> v)
{
	ListNode *root = NULL;
	ListNode *p = NULL;
	for (int iv : v)
	{
		auto n = new ListNode(iv);
		if (p != NULL)
		{
			p->next = n;
		} else {
			root = n;
		}
		p = n;
	}
	return root;
}
void print_list(ListNode * list)
{
	while (list) {
		cout<<list->val<<" ";
		list = list->next;
	}
	cout<<endl;
}
int main(int argc, char const *argv[])
{
	vector<int> a={2,4,3};
	vector<int> b={5,6,4};
	auto al = build_list(a);
	auto bl = build_list(b);
	Solution s;
	auto r = s.addTwoNumbers(al, bl);
	print_list(r);
	vector<int> a1={5};
	vector<int> b1={5};
	al = build_list(a1);
	bl = build_list(b1);
	r = s.addTwoNumbers(al, bl);
	print_list(r);
	return 0;
}
