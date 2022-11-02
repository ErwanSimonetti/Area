package jobs

import (
	"fmt"
	// "time"
	"github.com/jasonlvhit/gocron"
	"AREA/pkg/controllers"
)


func ExecJob(action func() (bool, error), reaction func(string), react_value string) {
	newAction, _ := action()
	if (newAction == true) {
		reaction(react_value)
	}
}

func task() {
	
}
// // func DeleteJob() {
// // 	gocron.Remove(task)
// // }

func NewScheduler() {
	fmt.Println("newmachin lol")
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(ExecJob, controllers.TemperatureIsUnder24, controllers.SendMessage, "13")
	// s.Every(3).Seconds().Do(task)
	<- s.Start()
}