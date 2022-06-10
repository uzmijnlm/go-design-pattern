package main

import (
	"sync"
)

func main() {
	for {
		wg := &sync.WaitGroup{}
		wg.Add(5)
		for i := 0; i < 5; i++ {
			var j = i
			go getInstance(j, wg)
		}
		wg.Wait()
		clear()
	}
	// 下面的代码可以证明go中没有volatile关键字，用dcl实现单例模式会遇到和java一样的问题，建议使用once方式来实现
	//for {
	//	wg := &sync.WaitGroup{}
	//	wg.Add(2)
	//	for i := 0; i < 2; i++ {
	//		go func() {
	//			instance := getInstance1(wg)
	//			if instance.a == 0 {
	//				fmt.Println("wow")
	//			}
	//		}()
	//	}
	//	// Scanln is similar to Scan, but stops scanning at a newline and
	//	// after the final item there must be a newline or EOF.
	//	wg.Wait()
	//	clear1()
	//}

}
