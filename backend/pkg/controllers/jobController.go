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

func AddJobToUser(w http.ResponseWriter, r *http.Request) {
	newJob := &models.Job{}
	utils.ParseBody(r, newJob)

	requestUser, _ := GetUser(w, r)
	newJob.UserId = requestUser.ID
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
	jobs.RemoveJobByID(uint(jobId))
	w.WriteHeader(http.StatusOK)

}

func GetUserJobs(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	jobs := models.GetJobsByUserId(requestUser.ID)

	res, _ := json.Marshal(jobs)
	w.Header().Set("Content-Type","pkglication/json")
	// w.WriteHeader(http.StatusOK)
	w.Write(res)
}