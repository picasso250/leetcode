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
    bool isSymmetric(TreeNode *root) { // root=2
        if (root == NULL) { // no
            return true;
        }
        stack<TreeNode *> s; // from top to down, from out to inner
        auto left = root->left;   // left=3
        auto right = root->right; // right=3
        while (true) {
            // assert left.parent is not NULL
            // assert right.parent is not NULL
            // assert left and right are same level
            // cout << "-----" << endl;
            if (left == NULL && right == NULL) {
                if (!s.empty()) {
                    // cout << "pop" << endl;
                    right = s.top();
                    s.pop();
                    left = s.top();
                    s.pop();
                    continue;
                } else {
                    return true;
                }
            }
            if (left == NULL || right == NULL) {
                return false;
            }
            // cout << "compare " << right->val << " " << left->val << endl;
            if (right->val != left->val) {
                return false;
            }
            // cout << "go down " << endl;
            if (right->right || left->left) {
                // cout << "push" << endl;
                s.push(left->right);
                s.push(right->left);
                left = left->left;
                right = right->right;
                // cout << "will push?" << endl;
                // cout << left->right << " " << right->left <<endl;
            } else {
                // cout << "out end, go inner" << endl;
                left = left->right;
                right = right->left;
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
