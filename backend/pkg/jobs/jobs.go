package jobs

import (
	// "AREA/pkg/controllers"
	"AREA/pkg/models"
	"fmt"
)

// type Job struct {
// 	ActionFunc func(string) bool
// 	ActionFuncParams string
// 	ReactionFunc func(uint)
// 	ReactionFuncParams uint
// }

var currentJobs []models.Job

var ActionMap = map[string]func(string)bool{
	"weather": test,
	"test_action": action,
}

var ReactionMap = map[string]func(uint, string){
	// "discord": SendMessage, 
	"test_reaction": reaction,
}

// func CreateNewJob( action string, reaction string, paramA string, paramR uint ) {
// 	var newJob Job

// 	newJob.ActionFunc = ActionMap[action]
// 	newJob.ReactionFunc = ReactionMap[reaction]
// 	newJob.ActionFuncParams = paramA
// 	newJob.ReactionFuncParams = paramR

// 	currentJobs = append(currentJobs, newJob)
// }

func test(ok string) bool {
	fmt.Println("action")
	return true
}

func action(ok string) bool {
	fmt.Println("action", ok)
	return true
}

func reaction(userID uint, params string) {
	fmt.Println("reaction", userID, params)
}

func AddUserJobsOnLogin(userId uint) {
	jobs := models.GetJobsByUserId(userId)
	currentJobs = append(currentJobs, jobs...)
}

func AddJob(newJob models.Job) {
	currentJobs = append(currentJobs, newJob)
}

func RemoveJobByID(jobId uint) {
	var newCurrentJobs []models.Job
	for _, job := range currentJobs {
		if (job.UserId == jobId) {
			continue
		}
		newCurrentJobs = append(newCurrentJobs, job)
	}
	currentJobs = newCurrentJobs
}

func SuprUserJobsOnLogout(userId uint) {
	var newCurrentJobs []models.Job

	for _, job := range currentJobs {
		if (job.UserId == userId) {
			continue
		}
		newCurrentJobs = append(newCurrentJobs, job)
	}
	
	currentJobs = newCurrentJobs
}

func ExecAllJob() {
	fmt.Println(currentJobs)

	for _, job := range currentJobs {

		if ActionMap[job.ActionFunc](job.ActionFuncParams) {
			ReactionMap[job.ReactionFunc](job.UserId, job.ReactionFuncParams)
		}
	}
	fmt.Println()
	fmt.Println()
}