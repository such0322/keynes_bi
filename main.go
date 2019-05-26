package main

import (
	"encoding/json"
	"fmt"
	"keynesTe/models"
	"time"
)

func main() {

	NewDB()

	Run()

}

func NewDB() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1:33306", "keynes_bi")
	conf := models.Config{
		DSN:      connStr,
		Idle:     100,
		Active:   10,
		IdleTime: 10,
	}
	models.NewMysql(&conf)

}

func Run() {
	var t1 models.Te
	t1.LoadTeById(1)
	fmt.Println(t1)
	fmt.Printf("%#v\n", t1)
	//
	//var t2 models.Te
	//t2.Title = "ccc"
	//fmt.Printf("%#v\n", time.Now())
	//t2.CreateTime = time.Now()
	//t2.Save()

	var t3 models.Te
	jsonStr := `{"title":"abc","create_time":"2019-01-01 12:11:15"}`
	json.Unmarshal([]byte(jsonStr), &t3)
	fmt.Println(t3)
	fmt.Printf("%#v\n", t3)
	fmt.Printf("%#v\n", time.Now())

	b, _ := json.Marshal(t3)
	fmt.Println(string(b))
	t3.Save()
}
