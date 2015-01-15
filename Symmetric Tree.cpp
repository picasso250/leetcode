#include <iostream>
#include <stack>

using namespace std;

 struct TreeNode {
     int val;
     TreeNode *left;
     TreeNode *right;
     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
 
// class Solution {
// public:
//     bool isSymmetric(TreeNode *root) {
//         if (root == NULL) {
//             return true;
//         }
//         if (root->left && root->right) {
//             return this->isMutual(root->left, root->right);
//         }
//         return root->left == NULL && root->right == NULL;
//     }
//     bool isMutual(TreeNode *a, TreeNode *b) {
//         if (a == NULL || b == NULL)
//         {
//             return a == b;
//         }
//         cout << a->val << " == " << b->val << endl;
//         if (a->val == b->val)
//         {
//             // cout << "a->left == b->right\t" << this->isMutual(a->left, b->right) << endl;
//             // cout << "a->right == b->left\t" << this->isMutual(a->right, b->left) << endl;
//             return this->isMutual(a->left, b->right) && this->isMutual(a->right, b->left);
//         }
//         return false;
//     }
// };


class Solution {
public:
    bool isSymmetric(TreeNode *root) {
        if (root == NULL) {
            return true;
        }
        stack<TreeNode *> s; // from top to down, from out to inner
        auto left = root->left;
        auto right = root->right;
        while (true) {
            cout << "---" << endl;
            if ((left || right) && !(left && right)) {
                return false;
            }
            if (left == NULL && right == NULL) {
                if (!s.empty())
                {
                    cout << "pop" <<endl;
                    right = s.top();
                    s.pop();
                    left = s.top();
                    s.pop();
                    continue;
                }
                return true;
            }
            cout << "compare " << right->val << " " << left->val << endl;
            if (right->val != left->val) {
                return false;
            }
            if (right->right && left->left) {
                left = left->left;
                right = right->right;
                cout << "will push?" << endl;
                cout << left->right << " " << right->left <<endl;
                if (left->right || right->left) {
                    cout << "push" << endl;
                    s.push(left->right);
                    s.push(right->left);
                }
            } else {
                if (right->right == NULL && left->left == NULL) {
                    cout << "out end, go inner" << endl;
                    left = left->right;
                    right = right->left;
                } else {
                    return false;
                }
            }
        }
    }
};

int main(int argc, char const *argv[])
{
    TreeNode a(2);
    TreeNode b(3);
    TreeNode c(3);
    TreeNode d(4);
    TreeNode e(5);
    TreeNode f(4);
    a.left = &b;
    a.right = &c;
    b.left = &d;
    b.right = &e;
    c.right = &f;
    Solution s;
    auto is_s = s.isSymmetric(&a);
    cout << is_s << endl;
    return 0;
}
