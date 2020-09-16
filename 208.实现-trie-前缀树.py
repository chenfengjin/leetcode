#
# @lc app=leetcode.cn id=208 lang=python3
#
# [208] 实现 Trie (前缀树)
#

# @lc code=start
class TrieNode:
    def __init__(self):
        self.subTree = {}
        self.is_leaf = False
        # self.value = 
# 关键是对不同的类的职责划分，不然写出来很难受  
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
        for c in word:
            if not c in current.subTree:
                current.subTree[c] = TrieNode()
            current = current.subTree.get(c)
        current.is_leaf = True                
            
            


    def search(self, word: str) -> bool:
        """
        Returns if the word is in the trie.
        """
        current = self.root
        for c in word:
            if not c in current.subTree:
                return False
            current = current.subTree.get(c)
        if not current.is_leaf:
            return False
        return True


    def startsWith(self, prefix: str) -> bool:
        """
        Returns if there is any word in the trie that starts with the given prefix.
        """
        current = self.root
        for c in prefix:
            if not c in current.subTree:
                return False
            current = current.subTree.get(c)
        return True



# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
# @lc code=end

