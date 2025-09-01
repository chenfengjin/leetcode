 #ifndef LEETCODE_H
  struct ListNode {
     int val;
     ListNode *next;
     ListNode() : val(0), next(nullptr) {}
     ListNode(int x) : val(x), next(nullptr) {}
     ListNode(int x, ListNode *next) : val(x), next(next) {}
  };

 
// 输入输出流
#include <iostream>    // 标准输入输出流
#include <fstream>     // 文件流
#include <sstream>     // 字符串流
#include <iomanip>     // 输入输出格式化

// 容器类
#include <vector>      // 动态数组
#include <list>        // 双向链表
#include <deque>       // 双端队列
#include <queue>       // 队列（包含priority_queue）
#include <stack>       // 栈
#include <set>         // 集合（包含multiset）
#include <map>         // 映射（包含multimap）
#include <unordered_set> // 无序集合
#include <unordered_map> // 无序映射
#include <array>       // 固定大小数组(C++11)
#include <forward_list> // 单向链表(C++11)

// 算法与迭代器
#include <algorithm>   // 算法库（排序、查找等）
#include <iterator>    // 迭代器
#include <functional>  // 函数对象和函数适配器

// 字符串处理
#include <string>      // 字符串类
#include <cstring>     // C风格字符串函数
#include <cctype>      // 字符处理函数

// 内存管理
#include <memory>      // 智能指针(C++11)

// 数学相关
#include <cmath>       // 数学函数
#include <complex>     // 复数
#include <valarray>    // 数值数组
#include <random>      // 随机数(C++11)
#include <numeric>     // 数值算法

// 时间与日期
#include <ctime>       // C风格时间函数
#include <chrono>      // 时间库(C++11)

// 异常处理
#include <exception>   // 异常基类
#include <stdexcept>   // 标准异常类

// 多线程
#include <thread>      // 线程(C++11)
#include <mutex>       // 互斥锁(C++11)
#include <atomic>      // 原子操作(C++11)
#include <future>      // 异步操作(C++11)

// 其他工具
#include <utility>     // 工具类（pair等）
#include <tuple>       // 元组(C++11)
#include <typeinfo>    // 类型信息
#include <type_traits> // 类型特性(C++11)
#include <limits>      // 数值极限
#include <cassert>     // 断言
#include <cstdlib>     // C标准库函数
#include <cstddef>     // 标准定义（size_t等）
#include <cstdint>     // 固定宽度整数类型(C++11)

// 初始化和工具函数
#include <initializer_list> // 初始化列表(C++11)
#include <bitset>      // 位集
#include <ratio>       // 编译时有理数算术(C++11)
#include <gtest/gtest.h>
using namespace std;


ListNode* create_list_from_vector_int(vector<int>nums){
    ListNode *dummy = new ListNode{};
    auto cur = dummy;
    for (auto num:nums){
        cur->next = new ListNode(num);
        cur = cur->next;
    }
    return dummy->next;
}
 struct TreeNode {
     int val;
     TreeNode *left;
      TreeNode *right;

      TreeNode() : val(0), left(nullptr), right(nullptr) {}
      TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
      TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left),
        right(right) {}
  };
bool link_list_equal(ListNode* l1,ListNode*l2){
    cout << "l1:";
    while (l1)
    {
        cout << l1->val << "->";
        l1 = l1->next;
    }
    cout << endl;
    cout << "l2:";
    while (l2){
        cout << l2->val << "->" ;
        l2 = l2->next;
    }
    while (l1 && l2)
    {
        if (l1->val!=l2->val){
            return false;
        }
        l1 = l1->next;
        l2 = l2->next;
    }
    return !l1 && !l2;
}

void print_link_list(ListNode*l){
    while (l){
        cout << l->val << "->";
        l = l->next;
    }
    cout << endl;
}

void print_vector(vector<int>nums){
    for (int num:nums){
        cout << num << "->";
    }
    cout << endl;
}
#define LEETCODE_H
 #endif // !LEETCODE_H
