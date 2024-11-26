package course

// https://leetcode.com/problems/course-schedule/description/
/*
 There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
 You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course
  bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.



Example 1:
0->1

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
Example 2:
0->1
1->0

1<->0
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]

Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.


Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
All the pairs prerequisites[i] are unique.
*/

// Steps for solution:
// Step 1:Converting our input into Adj
// Step 2: Marking indegree of nodes
// Step 3: Getting the nodes with 0 in-degree and adding it in queue
// Step 4: Our Main BFS until !q.isEmpty() -> keep on unmarking the indegrees of a node -> if in-degree becomes 0 then add it in queue and when it turns come in next iteration, it willl be added in the topo list.
// Time Complexity: O(V+E), where V = no. of nodes and E = no. of edges. This is a simple BFS algorithm.
// Space Complexity: O(N) + O(N) ~ O(2N), O(N) for the indegree array, and O(N) for the queue data structure used in BFS(where N = no.of nodes). Extra O(N) for storing the topological sorting.
// Total ~ O(3N)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// there are multiple trees, so we have to traverse them all
	graph := MakeAdjList(numCourses, prerequisites)
	indegrees := make([]int, numCourses)

	for _, neighbors := range graph {
		for _, j := range neighbors {
			indegrees[j]++
		}
	}
	var queue Queue = make([]int, 0, numCourses)
	for i := range indegrees {
		if indegrees[i] == 0 {
			queue.Enqueue(i)
		}
	}
	counter := 0
	for len(queue) != 0 {
		counter++
		node := queue.Dequeue()
		for _, key := range graph[node] {
			indegrees[key]--
			if indegrees[key] == 0 {
				queue.Enqueue(key)
			}
		}
	}

	return counter == numCourses

}

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() int {
	old := *q
	first := old[0]
	*q = old[1:]
	return first
}

func MakeAdjList(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)

	for _, edge := range prerequisites {
		course, prereq := edge[0], edge[1]
		graph[prereq] = append(graph[prereq], course)
	}

	return graph
}

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
Example 2:

Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
Output: [0,2,1,3]
Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].
Example 3:

Input: numCourses = 1, prerequisites = []
Output: [0]


Constraints:

1 <= numCourses <= 2000
0 <= prerequisites.length <= numCourses * (numCourses - 1)
prerequisites[i].length == 2
0 <= ai, bi < numCourses
ai != bi
All the pairs [ai, bi] are distinct.
*/

func FindOrder(numCourses int, prerequisites [][]int) []int {
	graph := MakeAdjList(numCourses, prerequisites)
	indegrees := make([]int, numCourses)
	for _, neighbors := range graph {
		for _, key := range neighbors {
			indegrees[key]++
		}
	}
	var queue Queue = make([]int, 0, numCourses)
	for i := range indegrees {
		if indegrees[i] == 0 {
			queue.Enqueue(i)
		}
	}
	result := make([]int, 0, numCourses)
	for len(queue) != 0 {
		key := queue.Dequeue()
		result = append(result, key)
		for _, neighbor := range graph[key] {
			indegrees[neighbor]--
			if indegrees[neighbor] == 0 {
				queue.Enqueue(neighbor)
			}
		}

	}

	if len(result) != numCourses {
		return nil
	}
	return result
}
