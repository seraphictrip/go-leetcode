# Patterns

https://www.youtube.com/watch?v=DjYZk8nrXVY

## Prefix sum

303. https://leetcode.com/problems/range-sum-query-immutable/description/

525. https://leetcode.com/problems/contiguous-array/description/

560. https://leetcode.com/problems/subarray-sum-equals-k/

## Two Pointers

167. https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

15. https://leetcode.com/problems/3sum/description/

11. https://leetcode.com/problems/container-with-most-water/description/

### Fast And Slow Pointers
141. https://leetcode.com/problems/linked-list-cycle/description/

202. https://leetcode.com/problems/happy-number/description/

287. https://leetcode.com/problems/find-the-duplicate-number/description/

## Sliding window
643. https://leetcode.com/problems/maximum-average-subarray-i/description/

3. https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

76. https://leetcode.com/problems/minimum-window-substring/description/


## Linked List inplace reversal
```
// I definitely need to deep dive on this code, but re-re for now
func reverselinkedlist(head *ListNode) *ListNode {
    prev := nil
    current := head

    for current != nil {
        next := current.next
        current.next = prev
        prev = current
        current = next
    }

    return prev
}
```
206. https://leetcode.com/problems/reverse-linked-list/description/

92. https://leetcode.com/problems/reverse-linked-list-ii/description/

24. https://leetcode.com/problems/swap-nodes-in-pairs/description/

## Monotonic Stack
find next greater or next smaller element in a stack

496. https://leetcode.com/problems/next-greater-element-i/description/

739. https://leetcode.com/problems/daily-temperatures/description/

84. https://leetcode.com/problems/largest-rectangle-in-histogram/description/

# Top 'K' Elements

* K Largest: min-heap
* K Smallest: max-heap
* K Most Frequent
* quickselect: https://en.wikipedia.org/wiki/Quickselect (academic)

215. https://leetcode.com/problems/kth-largest-element-in-an-array/description/

347. https://leetcode.com/problems/top-k-frequent-elements/description/

373. https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/

## Ovderlappin intervals or ranges
* Merging Intervals
* Interval Intersection
* Insert Interval
* Finding Minimum Meeting rooms

56. https://leetcode.com/problems/merge-intervals/description/
57. https://leetcode.com/problems/insert-interval/description/
435. https://leetcode.com/problems/non-overlapping-intervals/description/


## Modified Binary Search
* Searching in a "Nearly" sorted array
* Searching in a "rotated" sorted array
* searching in list unknown length
* searchign in array with duplicates
* first or last occurence
* finding peak


## Binary Tree Traversal
* PreOrder: root->left->right
* InOrder: left->root->right
* PostOrder: left->right->root
* LevelOrder: level by level

```
// need to spend more time, but do so re-re right now

// root -> left -> right
func preorder(node *BinaryTreeNode) {
    if node != nil {
        fmt.Print(node.val, " ")
    }
    preorder(node.left)
    preorder(node.right)
}
// left -> root -> right
func inorder(node *BinaryTreeNode) {
    if node != nil {
        inorder(node.left)
        fmt.Print(node.val, " ")
        inorder(node.right)
    }
}
// left -> right -> root
func postorder(node *BinaryTreeNode) {
    if node != nil {
        postorder(node.left)
        postorder(node.right)
        fmt.Print(node.val, " ")
    }
}


func levelorder(node *BinaryTreeNode) []*BinaryTreeNode {
    result := make([]*BinaryTreeNode, 0)
    if node == nil {
        return result
    }
    // ... TODO, implement
}
```
### Inorder
in sorted order

230. https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/
### Preorder
Creating a copy of a tree

257. https://leetcode.com/problems/binary-tree-paths/description/
### Postorder
Want to process children before parent, 
124. https://leetcode.com/problems/binary-tree-maximum-path-sum/description/
### Level order
need to explore all nodes of current level
107. https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/

257. https://leetcode.com/problems/binary-tree-paths/description/

230. https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

124. https://leetcode.com/problems/binary-tree-maximum-path-sum/description/

107. https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/


# Depth-First Search : DFS
* Path between two nodes
* Checking for cycle in graph
* Finding a topological order in a DAG (seems interesting to explore...)

```
func dfs(graph Graph, v *TreeNode, visted map[*TreeNode]bool) {
    visited[v] = true
    fmt.Print(v, " ")

    for _, neighor := range graph[v] {
        if !visited[neighbor] {
            dfs(graph, nieghbor, visited)
        }
    }
}

```

133. https://leetcode.com/problems/clone-graph/description/
113. https://leetcode.com/problems/path-sum-ii/description/

210. https://leetcode.com/problems/course-schedule-ii/description/

# Breadth-First Search: BFS
* Finding Shortest Path bewteen two nodes
* Printing nodes level by level
* finding all connected components
* finding shortest transformation sequence from one word to another

```
func bfs(graph Graph, start *Node) {
    visited := map[*Node]bool{}
    queue = make([]*Node, 0)
    for len(queue) != 0 {
        node := queue[0]
        queue := queue[1:]
        fmt.Print(node, " ")

        for neighbor := graph[node] {
            if !visited[neighbor] {
                visited[neighbor] = true
                queue = append(queue, neighbor)
            }
        }

    }
}
```

102. https://leetcode.com/problems/binary-tree-level-order-traversal/description/

994. https://leetcode.com/problems/rotting-oranges/description/

127. https://leetcode.com/problems/word-ladder/description/


(unrelated, but 730 keeps calling to me, check it out)
## Matrix Traversal
733. https://leetcode.com/problems/flood-fill/description/

200. https://leetcode.com/problems/number-of-islands/description/

130. https://leetcode.com/problems/surrounded-regions/description/


# Backtracking (permutation/combination)
Exploring all potential solution paths and backtracking paths without valid solutions

* Generate all possible Permutations/Combinations of a given set of elements
* suduko or n queens problem

46. https://leetcode.com/problems/permutations/description/

78. https://leetcode.com/problems/subsets/description/

51. https://leetcode.com/problems/n-queens/description/


## Dynamic Programming / DP
Solving Optimization problems by breaking them down into smaller
sub problems

### ideals
* Overlapping subproblems
* Optimal substructure (investigate)

### Problem type
* Maximize or minimize a certain value
* count number of ways

### Types
* Fibonacci Numbers (memo)
* 0/1 Knapsack
* Longest common subsequence
* Longest increasing subsequence
* Subset sum
* Matrix chain multiplication

70. https://leetcode.com/problems/climbing-stairs/description/
322. https://leetcode.com/problems/coin-change/description/
1143. https://leetcode.com/problems/longest-common-subsequence/description/
300. https://leetcode.com/problems/longest-increasing-subsequence/description/
416. https://leetcode.com/problems/partition-equal-subset-sum/description/
312. https://leetcode.com/problems/burst-balloons/description/


## Stephen's 

## functional