package arrays_test

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
)

var InsertEndTests = []struct {
	arr      []int
	val      int
	expected int
}{
	{},
	{make([]int, 0, 5), 6, 1},
	{make([]int, 0, 1), 6, 1},
	{make([]int, 0), 6, 1},
}

func TestInsertEnd(t *testing.T) {
	for i, e := range InsertEndTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := InsertCappedEnd(e.arr, e.val)
			fmt.Printf("%v : %p, %p cap: %v\n", actual, e.arr, actual, cap(actual))
			// if actual != e.expected {
			// 	t.Fatalf("MaybeInsertEnd(%v, %v) = %v, want %v", e.arr, e.val, actual, e.expected)
			// }

		})
	}
}

func TestA(t *testing.T) {
	// a := make([]int, 0, 1)
	fmt.Println(strings.Count("test", ""), len("test"))
}

func InsertCappedEnd(arr []int, val int) []int {
	if cap(arr) == len(arr) {
		// don't insert
		return arr
	}
	arr = append(arr, val)
	return arr
}

var InsertMiddleTests = []struct {
	arr      []int
	i        int
	val      int
	expected []int
	err      error
}{
	// {
	// 	// can't insert if no capacity
	// 	[]int{}, 0, 0, []int{}, ErrOverflow,
	// },
	// {
	// 	// can insert at end
	// 	make([]int, 0, 1), 0, 0, []int{0}, nil,
	// },
	// {
	// 	// can't insert past end, even if capacity
	// 	make([]int, 0, 2), 1, 0, []int{}, ErrOutOfBounds,
	// },
	{
		// can't insert past end, even if capacity
		make([]int, 2, 10), 0, 100, []int{100, 0, 0}, nil,
	},
}

func TestInsertMiddle(t *testing.T) {

	for i, e := range InsertMiddleTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result, err := insertMiddle(e.arr, e.i, e.val)
			if !errors.Is(err, e.err) {
				t.Fatalf("unexpected error, got %v, want %v", err, e.err)
			}
			if !slices.Equal(result, e.expected) {
				t.Fatalf("insertMiddle(%v, %v, %v) = %v, want %v", e.arr, e.i, e.val, result, e.expected)
			}
			fmt.Printf("%p:%p", e.arr, result)
		})
	}
}

var (
	ErrOverflow    = errors.New("overflow")
	ErrOutOfBounds = errors.New("out of bounds")
)

// insert at i if and only if i < len(arr) and cap(arr) has room
// return length
// [1,2,3,4,5] 0, 100
// [

func insertMiddle(arr []int, i, val int) ([]int, error) {
	// only allow insert within bounds of current lenght
	// including the end
	if i < 0 || i > len(arr) {
		return arr, fmt.Errorf("%w: %v", ErrOutOfBounds, i)
	}
	// if we are out of capacity, overflow error
	if cap(arr) == len(arr) {
		return arr, ErrOverflow
	}
	// append and swap into place
	arr = append(arr, val)
	curr := len(arr) - 1
	for curr > i {
		swap(arr, curr, curr-1)
		curr--
	}
	return arr, nil
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type DynamicArray struct {
	cap    int
	length int
	arr    []int
}

func NewDynamicArray(cap int) *DynamicArray {
	return &DynamicArray{
		cap:    cap,
		length: 0,
		arr:    make([]int, cap), // not idiomatic go, just following concept
	}
}

func (da *DynamicArray) String() string {
	return fmt.Sprintf("%v", da.arr[:da.length])
}

func (da *DynamicArray) Get(i int) int {
	if i < 0 || i > da.length {
		panic(ErrOutOfBounds)
	}
	return da.arr[i]
}

func (da *DynamicArray) Set(i, n int) {
	if i < 0 || i > da.length {
		panic(ErrOutOfBounds)
	}
	da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.length == da.cap {
		da.resize()
	}
	da.arr[da.length] = n
	da.length++
}

func (da *DynamicArray) resize() {
	da.cap = da.cap * 2
	arr := make([]int, da.cap)
	// for i := range da.arr {
	// 	arr[i] = da.arr[i]
	// }
	copy(arr, da.arr)
	da.arr = arr
}

func (da *DynamicArray) Popback() int {
	da.length--
	return da.arr[da.length]
}

func (da *DynamicArray) GetSize() int {
	return da.length
}

func (da *DynamicArray) GetCapacity() int {
	return da.cap
}

var DynamicArrayTests = []struct {
	cap     int
	inserts []int
}{
	{10, []int{1, 2, 3, 4, 5}},
}

func TestDynamicArray(t *testing.T) {
	for i, e := range DynamicArrayTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			arr := NewDynamicArray(e.cap)
			for _, insert := range e.inserts {
				arr.Pushback(insert)
				fmt.Println(arr.GetCapacity(), arr)
			}
		})
	}
}

/*
A company is planning to interview 2n people. Given the array costs where costs[i] = [aCosti, bCosti], the cost of flying the ith person to city a is aCosti, and the cost of flying the ith person to city b is bCosti.

Return the minimum cost to fly every person to a city such that exactly n people arrive in each city.



Example 1:
A	B
10	20
30	200
400	50
30	20
Input: costs = [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city B for a cost of 20.

The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.
Example 2:

Input: costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
	A	B	A-B	B-A
0	259	770	-511	511
1	448	54	394	-394
2	926	667	259	-259
3	184	139	45	-45
4	840	118	722	-722
5	577	469	108	-108
Output: 1859
Example 3:

Input: costs = [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]
Output: 3086


Constraints:

2 * n == costs.length
2 <= costs.length <= 100
costs.length is even.
1 <= aCosti, bCosti <= 1000
*/

// [[10,20],[30,200],[400,50],[30,20]]
// refunds = [10, 170, -350, -10]
// mincost = 10, 40, 440, 470
// refundsSorted = [-350, -10, 10, 170]
// mincost after refund 470+-360 = 110

// [[259,770],
//  [448,54],
//  [926,667],
//  [184,139],
//  [840,118],
//  [577,469]]

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	refund := make([]int, n) // SUM(B-A)
	// refund = [0,0,0,0,0,0]

	minCost := 0 // SUM(A)
	for i, cost := range costs {
		refund[i] = cost[1] - cost[0]
		minCost += cost[0]
	}
	slices.Sort(refund)
	for i := 0; i < n/2; i++ {
		minCost += refund[i]
	}

	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(nums []int) int {
	acc := 0
	for i := 0; i < len(nums); i++ {
		acc += nums[i]
	}
	return acc
}

// Sum(A) + SUM(MIN_HALF(B-A))
// Sum()
func two(costs [][]int) int {
	totalA := 0
	totalBMinusA := 0
	BMinusA := make([]int, len(costs))

	for i, amounts := range costs {
		totalA += amounts[0]
		totalBMinusA += amounts[1] - amounts[0]
		BMinusA[i] = amounts[1] - amounts[0]
	}
	// add back in cheapest half of A
	slices.Sort(BMinusA)
	minHalf := sum(BMinusA[:len(BMinusA)/2])

	return totalA + minHalf
}

// SUM(B) + MIN_HALF(SUM(A-B))
func two2(costs [][]int) int {
	n := len(costs)
	B := make([]int, n)
	AMinusB := make([]int, n)
	for i, cost := range costs {
		B[i] = cost[1]
		AMinusB[i] = cost[0] - cost[1]
	}
	sumB := sum(B)
	slices.Sort(AMinusB)
	minHalf := sum(AMinusB[:n/2])

	return sumB + minHalf
}

func sumMinN(arr []int, n int) int {
	slices.Sort(arr)
	return sum(arr[:n])
}

// SUM(B) + MIN_HALF(SUM(A-B))
func two3(costs [][]int) int {
	n := len(costs)
	B := make([]int, n)
	AMinusB := make([]int, n)
	for i, cost := range costs {
		B[i] = cost[1]
		AMinusB[i] = cost[0] - cost[1]
	}
	sumB := sum(B)
	minHalfAMinusB := sumMinN(AMinusB, n/2)

	return sumB + minHalfAMinusB
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// SUM(B) - SUM(MIN_HALF(|A-B|))
func two4(costs [][]int) int {
	n := len(costs)
	B := make([]int, n)
	AMinusB := make([]int, n)
	for i, cost := range costs {
		B[i] = cost[1]
		AMinusB[i] = cost[0] - cost[1]
	}
	sumB := sum(B)
	minHalfAMinusB := Abs(sumMinN(AMinusB, n/2))

	return sumB - minHalfAMinusB
}

func twoChoice(costs [][]int) int {
	// sort based on A-B where a cost[0] (City A) and b cost[1] (City B)
	// This partitions our data ChooseA[:mid] and ChooseB[mid:]
	slices.SortFunc(costs, func(a, b []int) int {
		// cost func/choice A-B, then just sort using that
		// if use B-A need to swap
		// (b[1] - b[0]) - (a[1] - a[0])
		return (a[0] - a[1]) - (b[0] - b[1])
	})
	fmt.Println("PARTIIION: ", costs)
	mid := len(costs) / 2
	// our A choices bassed on choice function A-B
	As := costs[0:mid]
	// Our B choices
	Bs := costs[mid:]
	minCost := 0
	for i := 0; i < len(As); i++ {
		minCost += As[i][0] + Bs[i][1]
	}
	return minCost
}

var twoCitySchedCostTests = []struct {
	costs    [][]int
	expected int
}{
	{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
	{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
	{[][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}, 3086},
}

func TestTwoCitySchedCost(t *testing.T) {
	for i, e := range twoCitySchedCostTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := twoCitySchedCost(e.costs)
			if actual != e.expected {
				t.Fatalf("twoCitySchedCost(%v) = %v, want %v", e.costs, actual, e.expected)
			}
			actual = twoChoice(e.costs)
			if actual != e.expected {
				t.Fatalf("two(%v) = %v, want %v", e.costs, actual, e.expected)
			}
		})
	}
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)

	for _, ch := range s {
		if opener, ok := closeToOpen[ch]; ok {
			lastIndex := len(stack) - 1
			top := stack[lastIndex]
			if top != opener {
				return false
			}
			// pop
			stack = stack[0:lastIndex]

		} else {
			// push
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

var closeToOpen = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"", true},
	{"{}", true},
	{"{()}", true},
	{"()[]{}", true},
}

// ["([{", ")]}"

// ()[  ]{}
func TestIsValid(t *testing.T) {
	for i, e := range isValidTests {
		e := e
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := isValid(e.input)
			if actual != e.expected {
				t.Fatalf("isValid(%q) = %v, want %v", e.input, actual, e.expected)
			}
		})
	}
}
