package cron

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/cron"
)

func Route(rg *gin.RouterGroup, cronTask *cron.Task) {
	rg.GET("/cron/jobs", func(c *gin.Context) {
		c.JSON(http.StatusOK, cronJobDetails(cronTask.Jobs()))
	})
}

func cronJobDetails(cronJobs []*cron.Job) []any {
	type jobDetail struct {
		Name string `json:"name,omitempty"`
		Spec string `json:"spec,omitempty"`
		Next string `json:"next,omitempty"`
		Prev string `json:"prev,omitempty"`
	}
	jobDetails := make([]any, 0, len(cronJobs))
	for _, job := range cronJobs {
		jobDetails = append(jobDetails, &jobDetail{
			Name: job.Name(),
			Spec: job.Spec(),
			Next: job.Next().String(),
			Prev: job.Prev().String(),
		})
	}
	return jobDetails
}
