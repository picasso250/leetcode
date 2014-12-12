# Definition for a  binary tree node
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# class Solution:
#     # @param root, a tree node
#     # @return a list of integers
#     def preorderTraversal(self, root):
#         if root is None:
#             return []
#         lst = [root.val]
#         lst.extend(self.preorderTraversal(root.left))
#         lst.extend(self.preorderTraversal(root.right))
#         return lst

class Solution:
    # @param root, a tree node
    # @return a list of integers
    def preorderTraversal(self, root):
        if root is None:
            return []
        lst = []
        stack = []
        p = root
        while True:
            if p:
                lst.append(p.val)
                if p.left:
                    stack.append(p)
                    p = p.left
                    continue
                if p.right:
                    p = p.right
                    continue
                if len(stack) > 0:
                    p = stack.pop()
                    p = p.right
                else:
                    break
        return lst

root = TreeNode(1)
root.left = TreeNode(2)
print(Solution().preorderTraversal(root))
