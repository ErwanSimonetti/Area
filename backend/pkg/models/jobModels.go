package models

import "github.com/jinzhu/gorm"

type Job struct {
	gorm.Model
	UserId uint `json:"user_id"`
	ActionService string `json:"action_service"`
	ActionFunc string `json:"action_func"`
	ActionFuncParams string `json:"action_func_params"`
	ReactionService string `json:"reaction_service"`
	ReactionFunc string `json:"reaction_func"`
	ReactionFuncParams string `json:"reaction_func_params"`
}

func (job *Job) CreateJob() uint {
	db.NewRecord(job)
	db.Create(&job)
	return job.ID
}

func GetJobById(Id uint) ([]Job){
	var jobs []Job
	db.Where("ID=?", Id).Find(&jobs)
	return jobs
}

func GetJobsByUserId(userId uint) ([]Job){
	var jobs []Job
	db.Where("user_id=?", userId).Find(&jobs)
	return jobs
}

func DeleteUserJob(ID uint) Job{
	var job Job
	db.Unscoped().Where("ID = ?", ID).Delete(&job)
	return job
}

func DeleteAllUserJob(userId uint) []Job{
	var job []Job
	db.Unscoped().Where("user_id = ?", userId).Delete(&job)
	return job
}

func CheckExistingGitAction(id uint, action string) bool{
	var job Job
	db.Where("user_id = ?", id).Where("action_func = ?", action).Find(&job)
	if (job.ActionFunc == "") {
		return false
	} else {
		return true
	}
}