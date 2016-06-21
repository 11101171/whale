//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package core

import (
	"whale/app/core/jobs"
	"whale/app/models"

	"github.com/revel/revel"
)

func init() {
	revel.INFO.Println("--------init--------")
}

func Init() {
	revel.INFO.Println("start Init")
	jobs.InitCorn()
}

func InitJobs() {
	list := models.SelectTaskList(1)
	for _, task := range list {
		job, err := jobs.NewJobFromTask(&task)
		if err != nil {
			revel.ERROR.Println("InitJobs:", err.Error())
			continue
		}
		jobs.AddJob(task.CronSpec, job)
	}
}
