//****************************************************/
// Addres: https://github.com/11101171
// Author: ningzhong.zeng
//****************************************************/

package jobs

import (
	"sync"

	"github.com/lisijie/cron"
	"github.com/revel/revel"
)

var (
	mainCron *cron.Cron
	workPool chan bool
	lock     sync.Mutex
)

var mate = `
========================
=== 开启corn workpool %d
========================
`

func InitCorn() {
	conf, err := revel.LoadConfig("app.conf")
	if err != nil {
		revel.ERROR.Panicln("读取app.conf失败")
	}
	cronPoolSize, _ := conf.Raw().Int(revel.RunMode, "cron.poolSize")
	workPool = make(chan bool, cronPoolSize)
	mainCron = cron.New()
	// job := NewCommandJob("11111", "zhangsan", "ls")
	// mainCron.Schedule(Every(5*time.Second+5*time.Nanosecond), job)
	// AddJob("1 1 10 14,18 9 ?", job)
	mainCron.Start()
	revel.INFO.Printf(mate, cronPoolSize)
}

func AddJob(spec string, job *Job) bool {
	lock.Lock()
	defer lock.Unlock()

	if GetEntryById(job.id) != nil {
		return false
	}

	err := mainCron.AddJob(spec, job)
	if err != nil {
		revel.ERROR.Println("AddJob: ", err.Error())
		return false
	}
	return true
}

func RemoveJob(id string) bool {
	mainCron.RemoveJob(func(e *cron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return true
			}
		}
		return false
	})
	return true
}

func GetEntryById(id string) *cron.Entry {
	entries := mainCron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}

func GetEntries(size int) []*cron.Entry {
	ret := mainCron.Entries()
	if len(ret) > size {
		return ret[:size]
	}
	return ret
}
