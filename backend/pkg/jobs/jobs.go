package jobs

import (
	"AREA/pkg/controllers"
	"fmt"
)

type Job struct {
	ActionFunc func(string) bool
	ActionFuncParams string
	ReactionFunc func(uint)
	ReactionFuncParams uint
}

var currentJob []Job

var ActionMap = map[string]func(string)bool{
	"weather": test,
}

var ReactionMap = map[string]func(uint){
	"discord": controllers.SendMessage, 
}

func CreateNewJob( action string, reaction string, paramA string, paramR uint ) {
	var newJob Job

	newJob.ActionFunc = ActionMap[action]
	newJob.ReactionFunc = ReactionMap[reaction]
	newJob.ActionFuncParams = paramA
	newJob.ReactionFuncParams = paramR

	currentJob = append(currentJob, newJob)
}

func test(ok string) bool {
	fmt.Println("action")
	return true
}

func ExecAllJob() {
	fmt.Println(currentJob)

	for _, job := range currentJob {

		if job.ActionFunc(job.ActionFuncParams) {
			job.ReactionFunc(job.ReactionFuncParams)
		}
	}
}