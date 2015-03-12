# golang_trie
Multiple String Matching


# 所实现的字典树
*  [Aho-Corasick](http://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_string_matching_algorithm)
    ([acmaptrie.go](https://github.com/renrenxing/golang_trie/blob/master/acmaptrie.go))
    多模式字符串精确匹配算法. O(n) 的时间复杂度, 对输入文本做一次遍历.
*  AC状态机的所使用的trie树是基于golang内建的map开发的.
*  AC算法, 对于添加新key支持不好, 每次添加需要重建AC状态机. 敏感词的更新是缓慢的, 所以本算法可以接受.
*  [双数组DAT](http://linux.thai.net/~thep/datrie/datrie.html#Double)
    ([dattrie.go](https://github.com/renrenxing/golang_trie/blob/master/dattrie.go)),
    使用的是列初始化, 并不实现白皮书中的行初始化. 列初始化但需要插入key时候, 需要重建DAT.
    行初始化产生的冲突较多, 并且数组使用效率低.

# 字典树后续改进
*  AC状态机结合DAT, 不直接使用map, DAT是一个压缩的状态机, 运行起来贼快.
*  使用[Wu-Manber](http://webglimpse.net/pubs/TR94-17.pdf)算法, WM对于添加删除支持很好.
    但需要考察, 当最小敏感词长度为2时, 是否有AC+DAT快.

# 一点背景
*  精确匹配
*  模糊匹配
*  单模式匹配
*  多模式匹配
*  trie与Double Array Trie
*  KMP与Aho-Corasick算法
*  Boyer Moore与Set Horspool与Wu-Manber算法

# 相关书籍
* Flexible Pattern Matching in Strings: Practical On-line Search Algorithms for Texts and Biological Sequences





