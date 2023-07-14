package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/guid"
	"ruoyi-vue3-go/src/com/gohiding/utils"
)

func main() {
	//guid
	RandomUUID := utils.RandomUUID()
	fmt.Println("guid=", RandomUUID)
	//simpleGuid
	SimpleUUID := utils.SimpleUUID()
	fmt.Println("simpleGuid=", SimpleUUID)

	// gf/guid
	Sequence := guid.S()
	fmt.Println("Sequence=", Sequence)
}
