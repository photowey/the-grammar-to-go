package oncet

import (
	"fmt"
	"sync"
)

// Singleton singleton struct
type Singleton struct {
	Id   int
	Name string
	Age  int
}

var once sync.Once

var Instance *Singleton

// GetInstance get the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("handle create the Singleton Instance.")
		Instance = &Singleton{Id: 9527, Name: "photowey", Age: 18}
	})

	return Instance
}
