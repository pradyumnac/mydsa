package main

import "fmt"

var done chan bool

func firstPairSum(inp []int, k int) {
	// var found bool
	// OuterLoop:
	defer wg.Done()
	for i := 0; i <= len(inp)-1; i = i + 1 {
		for j := i + 1; j <= len(inp)-1; j = j + 1 {
			if inp[i]+inp[j] == k {
				// found
				fmt.Println(i, j)
				<-done
				return
				// found = true
				// break OuterLoop
			}
		}
	}

	fmt.Println(-1, -1)
	<-done
	// if !found {
	// 	fmt.Println(-1, -1)
	// }
}

func main() {
	// inp [1,2,3,4,5]
	// k
	// result: first pair sum of which is eq k
	done = make(chan bool)
	wg.Add(1)
	go firstPairSum([]int{1, 2, 3, 4, 5}, 50)

	done <- true
}
