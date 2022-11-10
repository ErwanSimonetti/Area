package jobs

import (
	"fmt"

	"AREA/pkg/models"
)

var currentJobs []models.Job

var ActionMap = map[string]func(string) bool {
	"The temperature is over N degrees": TemperatureIsOverN,
	"The temperature is under N degrees": TemperatureIsUnderrN,
	"Check if the player main Teemo": IsPlayingTeemo,
	"The player winrate is over a given %": WinrateIsOverN,
	"The player KDA is over a given value": KDAIsOverN,
	"The covid case are over a given number": CovidCaseIsOverN,
	"The covid critical case are over a given number": CovidCriticalCaseIsOverN,
}

var ReactionMap = map[string]func(uint, string) {
	"Adds a given song to the user's queue": AddSongToQueue,
	"Send an email from user to given receiver": SendEmail,
	"Send a webhook message on selected channel": SendMessage,
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
		if ActionMap[job.ActionFunc] != nil && ActionMap[job.ActionFunc](job.ActionFuncParams) {
			ReactionMap[job.ReactionFunc](job.UserId, job.ReactionFuncParams)
		}
	}
	fmt.Println()
	fmt.Println()
}

func ExecGithJob(userID uint, githAction string) {
	for _, job := range currentJobs {
		if (job.ActionFunc == githAction) && (job.UserId == userID){
			ReactionMap[job.ReactionFunc](job.UserId, job.ReactionFuncParams)
		}
	}
}