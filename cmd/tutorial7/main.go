package main

import (
	"fmt"
	"sync"
	"time"
)

var dbData = []string{"db1", "db2", "db3", "db4", "db5"}
var results = []string{}
var wg = sync.WaitGroup{} // a counter to keep tabs on number of go routines working
var m = sync.Mutex{}      // mutual exclusion
var mrw = sync.RWMutex{}  // Read write mutex

func main() {
	t0 := time.Now()

	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1) // increments the counter by 1 for every iteration of loop
	// 	go dbCall(i)
	// }
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait() // waits for all routines to complete i.e counter goes back to 0
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Println(results)
}

func dbCall(i int) {
	//simulate db call
	var delay float32 = 2000 // when delay is constant many weird things happen like memory corruption as all routines try to write at the same time to result array
	fmt.Printf("Delay for %v is: %v\n", dbData[i], delay)
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println(time.Duration(delay) * time.Millisecond)
	//fmt.Println("The result from database is: ", dbData[i])
	//m.Lock() // code snippet between lock and unlock is performed singularly without disruption
	//results = append(results, dbData[i])
	//m.Unlock()
	save(dbData[i])
	log()
	fmt.Println()
	wg.Done() // decrements the counter by 1, signifying one task as done
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}

func save(result string) {
	mrw.Lock()
	results = append(results, result)
	mrw.Unlock()
}

func log() {
	mrw.RLock()
	fmt.Printf("Current results array is: %v\n", results)
	mrw.RUnlock()
}
