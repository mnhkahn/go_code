package tests

import (
	"fmt"
	"testing"

	"github.com/mnhkahn/go_code/goget"
)

func TestProcess(t *testing.T) {
	schedule := goget.NewGoGetSchedules(2)
	schedule.SetDownloadBlock(1)
	job := schedule.NextJob()
	fmt.Println(job)
	schedule.FinishJob(job)

	job = schedule.NextJob()
	fmt.Println(job)
	schedule.FinishJob(job)
}
