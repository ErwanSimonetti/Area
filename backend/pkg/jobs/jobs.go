/** @file jobs.go
 * @brief This file contain all the functions to handle the actions and reactions of the Email API
 * @author Juliette Destang
 * @version
 */

// @conv

package jobs

import (
	"AREA/pkg/models"
	// "fmt"
)

// @endconv

var currentJobs []models.Job

var ActionMap = map[string]func(string) bool {
	"The temperature is over a given value": TemperatureIsOverN,
	"The temperature is under a given value": TemperatureIsUnderrN,
	"Check if the player main Teemo": IsPlayingTeemo,
	"The player winrate is over a given %": WinrateIsOverN,
	"The player KDA is over a given value": KDAIsOverN,
	"The covid case are over a given number": CovidCaseIsOverN,
	"The covid critical case are over a given number": CovidCriticalCaseIsOverN,
}

var ReactionMap = map[string]func(uint, string) {
	"Adds a given song to the user's queue": AddSongToQueue,
	"Sends an email from user to given receiver": SendEmail,
	"Sends a webhook message on selected channel": SendMessage,
}

/** @brief This function take a user id and activate his job on login
 * @param userID uint
 */
func AddUserJobsOnLogin(userId uint) {
	jobs := models.GetJobsByUserId(userId)
	currentJobs = append(currentJobs, jobs...)
}

/** @brief This function take a Job model and append a new job to the currentJob
 * @param newJob models.Job
 */
func AddJob(newJob models.Job) {
	currentJobs = append(currentJobs, newJob)
}

/** @brief Remove a given job to the currentJob
 * @param jobId uint
 */
func RemoveJobByID(jobId uint) {
	var newCurrentJobs []models.Job
	for _, job := range currentJobs {
		if (job.ID == jobId) {
			continue
		}
		newCurrentJobs = append(newCurrentJobs, job)
	}
	currentJobs = newCurrentJobs
}

/** @brief Remove* all job from the currentJob when a user logout
 * @param userID uint
 */
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

/** @brief Execute all active jobs each X seconds thanks to a crontab
 */
func ExecAllJob() {
	// fmt.Println(currentJobs)
	for index := range currentJobs {
		if ActionMap[currentJobs[index].ActionFunc] != nil && ActionMap[currentJobs[index].ActionFunc](currentJobs[index].ActionFuncParams) {
			if (currentJobs[index].ActionExecuted == false) {
				currentJobs[index].ActionExecuted = true
				ReactionMap[currentJobs[index].ReactionFunc](currentJobs[index].UserId, currentJobs[index].ReactionFuncParams)
			}
		} else {
			currentJobs[index].ActionExecuted = false
		}
	}
	// fmt.Println()
	// fmt.Println()
}

/** @brief On ping from github api, execute the réaction of a given gitAction
 * @param userID uint, githAction string
 */
func ExecGithJob(userID uint, githAction string) {
	for _, job := range currentJobs {
		if (job.ActionFunc == githAction) && (job.UserId == userID){
			ReactionMap[job.ReactionFunc](job.UserId, job.ReactionFuncParams)
		}
	}
}