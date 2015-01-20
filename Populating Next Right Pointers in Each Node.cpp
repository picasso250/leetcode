/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        auto uf = root; // uf is first node of upper level
        while (uf) {
            // for every level
            // assert upper level is linked
            auto up = uf; // pointer of upper level
            auto _uf = NULL;
            for (;up->next; up = up->next) {
                if (_uf == NULL) {
                    _uf = up->left ? up->left : up->right;
                }
                if (up->left) {
                    up->left->next = up->right;
                }
                if (up->right) {
                    if (up->next) {
                        up->right->next = up->next->left;
                    }
                }
            }
            uf = _uf;
            // assert this level is linked
        }
        
    }
};
