package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func MapRace() {
	a := 1
	go func() {
		a = 2
	}()
	a = 3
	fmt.Println("a is =====", a)

	time.Sleep(2 * time.Second)

}

func safeWRMap() {
	var mu sync.RWMutex
	data := make(map[string]int)
	data["key1"] = 10
	data["key2"] = 20

	numRoutines := 50000000

	// 写数据
	var wg sync.WaitGroup
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock() // 获取读锁
			defer mu.Unlock()
			data[strconv.Itoa(i)] = i
		}(i)
	}

	// 读数据
	wg.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func(i int) {
			defer wg.Done()
			mu.RLock() // 获取读锁
			defer mu.RUnlock()
			// 读取字典中的值
			//fmt.Println("key: %v v: : %v", data[strconv.Itoa(i)], strconv.Itoa(i))
		}(i)
	}

	wg.Wait()
}

func ConcurrencyNotSafe() {
	var wg sync.WaitGroup
	data := make(map[int]int)
	numRoutines := 50000

	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		data[i] = i
	//		//fmt.Printf("Wrote: data[%d] = %d\n", i, i)
	//	}(i)
	//}

	for i := 0; i < numRoutines*10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			randomInRange := rand.Intn(numRoutines)
			fmt.Printf("Wrote: data[%d] = %d\n", i, i)
			fmt.Println("dict value:%d, %d", randomInRange, data[randomInRange])
		}(i)
	}

	wg.Wait()

	fmt.Println("\nFinal map content:")

	for i, v := range data {
		fmt.Printf("data[%d] = %d\n", i, v)
	}
}

// 结论1：多个goroutine同时读map不会出现数据竞争状态
// 结论2：多个goroutine同时"读写"同一个map，数据的竞争状态会出现的非常快速
func main() {
	fmt.Print("demos begin")
	//MapRace()
	ConcurrencyNotSafe()

	//safeWRMap()
	fmt.Print("demos end")
}
