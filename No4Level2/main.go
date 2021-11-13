package main

import (
	"fmt"
	"time"
)
func time1(a string){
	for {
		if time.Now().Format("15:04:05") == a {
			fmt.Println("谁能比我卷！")
			break
		}
	}
}
func time2(b string){
	for {
		if time.Now().Format("15:04:05") == b {
			fmt.Println("早八算什么，早六才是吾辈应起之时")
			break
		}
	}
}
func main() {
    go time1("2:00:00")
	go time2("6:00:00")
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("芜湖！起飞！")
		}
	}
}
