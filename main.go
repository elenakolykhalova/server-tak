package main


import (
	"fmt"
)

//для проверки сервисов
type HealthCheck struct {
	ServiceID int
	status string
}

var PassStatus = "pass"
var FailStatus = "fail"

func GenerateCheck() []HealthCheck{
	var s []HealthCheck

	for i :=0; i <=5; i++ {
		var temp HealthCheck
		temp.status = FailStatus
		temp.ServiceID = i
		if i%2 == 0 {
			temp.status = PassStatus
		}
		s = append(s, temp)
	}
	return s
}

func main() {
	fmt.Println("Тут будет выведен индикатор")
	s := GenerateCheck()

	for i:=0; i < len(s); i++ {
		if s[i].status == "pass" {
			fmt.Println(s[i].ServiceID)
		}
	}

}
