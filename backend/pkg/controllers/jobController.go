/** @file jobController.go
 * @brief This file contain all the functions to handle the job
 * @author Timothee de Boynes
 * @version
 */

// @cond
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"

	"AREA/pkg/jobs"
	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// var gitHubActions = []string{
// 	"push",
// 	"pull_request",
// 	"branch_protection_rule",
// 	"check_run",
// 	"check_suite",
// 	"create",
// 	"delete",
// 	"deployment",
// 	"deployment_status",
// 	"discussion",
// 	"discussion_comment",
// 	"fork",
// 	"gollum",
// 	"issue_comment",
// 	"issues",
// 	"label",
// 	"merge_group",
// 	"milestone",
// 	"page_build",
// 	"project_card",
// 	"project_column",
// 	"public",
// 	"pull_request_comment",
// 	"pull_request_review",
// 	"pull_request_review_comment",
// 	"pull_request_target",
// 	"registry_package",
// 	"release",
// 	"repository_dispatch",
// 	"schedule",
// 	"status",
// 	"watch",
// 	"workflow_call",
// 	"workflow_dispatch",
// 	"workflow_run",
// }

var GitHubActions = map[string]string{
	"Push action happened on your repository" : "push",
	"Pull request action happened on your repository": "pull_request",
	"If branch protection changed on your repository": "branch_protection_rule",
	"Checks run action on your repository": "check_run",
	"Checks suite action on your repository": "check_suite",
	"Someone created a git reference on your repository": "create",
	"Someone deleted a git reference on your repository": "delete",
	"Check deploy action on your repository": "deployment",
	"Check deploy status action on your repository": "deployment_status",
	"A discussion started on you repository": "discussion",
	"A comment was added to discussion on you repository": "discussion_comment",
	"Fork action on your repository": "fork",
	"Gollum action on your repository": "gollum",
	"Someone wrote an issue comment": "issue_comment",
	"Someone created an issue": "issues",
	"Someone created a label on your repository": "label",
	"A merge was perfomed on your repository": "merge_group",
	"A milestone was created on your repository" : "milestone",
	"Someone pushed on your publishing branch" : "page_build",
	"Created a project card on your repository" : "project_card",
	"Moved a card to a column in you repository" : "project_column",
	"Set your repoistory to public" : "public",
	"Add a comment to a pull request" : "pull_request_comment",
	"Someone reviewed your pull request" : "pull_request_review",
	"Added a comment to a review on some pull request" : "pull_request_review_comment",
	"Some activity was detected on a pull request" : "pull_request_target",
	"Some activity was detected on your package" : "registry_package",
	"Released action on your repository" : "release",
	"Dispatched action on your repository" : "repository_dispatch",
	"Added a schedule on your repository" : "schedule",
	"Changed the status of your repository" : "status",
	"Someone watched your repository" : "watch",
	"There was some activity on your workflow" : "workflow_call",
	"Dispatched action on your workflow" : "workflow_dispatch",
	"Ran your repository workflow" : "workflow_run",
}

// @endcond

/** @brief on a request, add a new job to a given user
 * @param w http.ResponseWriter, r *http.Request
 */
func AddJobToUser(w http.ResponseWriter, r *http.Request) {
	newJob := &models.Job{}
	utils.ParseBody(r, newJob)

	requestUser, _ := GetUser(w, r)
	newJob.UserId = requestUser.ID
	newJob.ActionExecuted = false

	gitActions := utils.GetKeyFromMap(GitHubActions)
	if utils.ArrayContainsString(gitActions, newJob.ActionFunc) {
		CreateWebhook(requestUser.ID, GitHubActions[newJob.ActionFunc], newJob.ActionFuncParams)
	}
	
	jobId := newJob.CreateJob()
	userToken := models.FindUserToken(newJob.UserId)
	
	newJob.ReactionFuncParams = newJob.ReactionFuncParams+"@@@"+strconv.FormatUint(uint64(jobId), 10)
	if newJob.ReactionService == "discord" {
		models.SetDiscordWebhook(newJob.UserId, jobId, userToken.CurrentDiscordWebhookId, userToken.CurrentDiscordWebhookToken)
		models.UpdateJobField(jobId, "reaction_func_params", newJob.ReactionFuncParams)
	}

	jobs.AddJob(*newJob)
	res, _ := json.Marshal(jobId)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, remove a job to a given user
 * @param w http.ResponseWriter, r *http.Request
 */
func RemoveJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobId, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	jobs.RemoveJobByID(uint(jobId))
	models.DeleteUserJob(uint(jobId))
}

/** @brief on a request, retrieve all a user's job
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUserJobs(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	jobs := models.GetJobsByUserId(requestUser.ID)

	res, _ := json.Marshal(jobs)
	// w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, retrieve the actions and reaction if the user is connected to the service
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUserPropositions(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	FillServices()
	services := Services
	var servicesOptions []Service

	tokens := *models.FindUserToken(requestUser.ID)
	for _, service := range services {
		if (service.Name == "email" && !models.CheckIfConnectedToService(tokens, "email")) {
			continue
		}
		if (service.Name == "discord" && !models.CheckIfConnectedToService(tokens, "discord")) {
			continue
		}
		if (service.Name == "spotify" && !models.CheckIfConnectedToService(tokens, "spotify")) {
			continue
		}
		if (service.Name == "github" && !models.CheckIfConnectedToService(tokens, "github")) {
			continue
		}
		servicesOptions = append(servicesOptions, service)
	}
	res, _ := json.Marshal(servicesOptions)
	// w.WriteHeader(http.StatusOK)
	w.Write(res)
}
