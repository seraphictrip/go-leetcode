# Stack

| Operation     |	Big-O Time Complexity   |	Notes                   |
|---------------|---------------------------|---------------------------|
| Push          | O(n)                      |  array backed migt be capped   |
| Pop           | O(n)*                     | check if isEmpty          |
| Top/Peek      | O(n)*                     | Retrieve without removing |


## Dynamic Array version easy, with slice
* Push is just append to end of the array
* Peek/Top is looking at last element `stack[len(stack)-1]`
* Pop is truncating slice `stack = stack[:len(stack)-1]`
