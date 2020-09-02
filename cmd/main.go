package main

import (
	"fmt"
	"myBlockchain/utils"
)

func main() {
	fmt.Println(utils.FindNeighbor("127.0.0.1", 5000, 0, 3, 5000, 5003))
}
