package main

import (
	"fmt"
	btsConfig "go1/config"
	"go1/pkg/config"
	"time"
)

type Person struct {
	Name string `json:"p_name"`
}

func (p Person) Sal() {
	fmt.Println(p.Name)
}

var Persons map[string]Person

func init() {
	Persons = make(map[string]Person)
	btsConfig.Initialize()
}

var sKey []byte

func main() {

	config.InitConfig("")
	name := config.Get("app.test")

	fmt.Println(name)

	return

	//Persons = append(Persons, Person{"hahaha"})
	//Persons = append(Persons, Person{"bbb"})
	//Persons = append(Persons, Person{"cccc"})

	//Persons["haha"] = Person{"hahaha"}

	//Persons = map[string]Person{
	//	"a": Person{"a"},
	//}

	//Persons["b"] = Person{"b"}
	//
	//js, _ := json.Marshal(Persons)
	//fmt.Printf("JSON format: %s \n", js)
	//
	//fmt.Printf("p type is %T value is %v \n", Persons, Persons)
	//
	////for name, p := range Persons {
	////	fmt.Println(name)
	////	p.Sal()
	////}
	//
	//sKey = []byte("aaa")
	//
	//fmt.Printf("sKye type is %T value is %v \n", sKey, sKey)
	//
	////var numCores = flag.Int("n", 3, "number of CPU cores to use")
	//
	//fmt.Println(runtime.NumCPU())
	//
	//return
	//bootstrap.SetupRedis()
	//
	//if cahce.NewCache().SetCache("name", "hello 222") {
	//	name := cahce.NewCache().GetCache("name", false)
	//
	//	fmt.Println(name)
	//}
	//
	//start := time.Now()
	//time.Sleep(5 * 1e9)
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s  start %s \n", delta, start)

	ch := make(chan int)

	sendData(ch)

	go getData(ch)
	time.Sleep(1e9)

}

func sendData(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	//
	//ch <- "aaa"
	//ch <- "bbb"
	//ch <- "ccc"

}

func getData(ch chan int) {

	var input int
	for {
		input = <-ch
		fmt.Println(input)
	}
}
