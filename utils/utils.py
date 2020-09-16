class TreeNode():
    def __init__(self,x):
        self.left = None
        self.right = None
        self.val = x

def list_to_tree(l):
    root = TreeNode(l[0])
    if len(l) == 1:
        return root
    root.left = list_to_tree(l[1:len(l)//2 + 1])
    root.right = list_to_tree(l[len(l)//2 + 1:])
    return root

def travel(root:TreeNode):
    print(root.val,end=",")
    travel(root.left)
    travel(root.right)

if __name__ == "__main__":
    root = list_to_tree([1,2,3,None,4])
    travel(root)
    