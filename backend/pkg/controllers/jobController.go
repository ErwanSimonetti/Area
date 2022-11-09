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

var gitHubActions = []string{
	"push",
	"pull_request",
	"branch_protection_rule",
	"check_run",
	"check_suite",
	"create",
	"delete",
	"deployment",
	"deployment_status",
	"discussion",
	"discussion_comment",
	"fork",
	"gollum",
	"issue_comment",
	"issues",
	"label",
	"merge_group",
	"milestone",
	"page_build",
	"project_card",
	"project_column",
	"public",
	"pull_request_comment",
	"pull_request_review",
	"pull_request_review_comment",
	"pull_request_target",
	"registry_package",
	"release",
	"repository_dispatch",
	"schedule",
	"status",
	"watch",
	"workflow_call",
	"workflow_dispatch",
	"workflow_run",
}

func AddJobToUser(w http.ResponseWriter, r *http.Request) {
	newJob := &models.Job{}
	utils.ParseBody(r, newJob)

	requestUser, _ := GetUser(w, r)
	newJob.UserId = requestUser.ID

	if (utils.ArrayContainsString(gitHubActions, newJob.ActionFunc)) {
		CreateWebhook(requestUser.ID, newJob.ActionFunc, newJob.ActionFuncParams)
	}
	jobs.AddJob(*newJob)

	jobId := newJob.CreateJob()
	res, _ := json.Marshal(jobId)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RemoveJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobId, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
        fmt.Println(err)
    }
	models.DeleteUserJob(uint(jobId))
}

func GetUserJobs(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	jobs := models.GetJobsByUserId(requestUser.ID)

	res, _ := json.Marshal(jobs)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}