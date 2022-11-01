package jobs

import (

)

type Job struct {
	ActionFunc func(string) bool
	ActionFuncParams string
	ReactionFunc func(string)
	ReactionFuncParams string
}

// func CreateNewJob(jobs []Job) Job {

// }

func ExecAllJob(jobs []Job) {
	for _, job := range jobs {

		if job.ActionFunc(job.ActionFuncParams) {
			job.ReactionFunc(job.ReactionFuncParams)
		}
	}
}