#
# @lc app=leetcode id=208 lang=python3
#
# [208] Implement Trie (Prefix Tree)
#

# @lc code=start
class TrieNode():
    def __init__(self,c = None,is_leaf=False):
        self.charactor = c
        self.children = {}
        self.is_leaf = False
class Trie:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        """
        current = self.root
        found = True
        for index,c in enumerate(word):
            if current.children.get(c):
                current = current.children.get(c)
                continue
            else:
                found = False
                break
        if found: 
            current.is_leaf = True
        else:
            for c in word[index:]:
                node = TrieNode(c=c)
                current.children[c] = node
                current = current.children[c]
        current.is_leaf = True


    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        current = self.root
        for char in word:
            if current.children.get(char):
                current = current.children.get(char)
            else:
                return False
        return True if current and current.is_leaf else False

    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        current = self.root
        for char in prefix:
            if current.children.get(char):
                current=current.children.get(char)
                continue
            else:
                return False
        return True 

# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
# @lc code=end

if __name__ == "__main__":
    root = Trie()
    root.insert("hello")
    print(root)
