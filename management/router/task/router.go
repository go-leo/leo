package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/runner/task/cron"
	"github.com/go-leo/leo/v2/runner/task/pubsub"
)

func Route(rg *gin.RouterGroup, cronJobs []*cron.Job, subJobs []*pubsub.Job) {
	rg.GET("/task/cron/jobs", func(c *gin.Context) {
		c.JSON(http.StatusOK, cronJobDetails(cronJobs))
	})
	rg.GET("/task/pubsub/jobs", func(c *gin.Context) {
		c.JSON(http.StatusOK, subJobDetails(subJobs))
	})
}

func subJobDetails(subJobs []*pubsub.Job) []any {
	type jobDetail struct {
		SubscribeTopic string `json:"subscriber,omitempty"`
		PublishTopic   string `json:"publish_topic,omitempty"`
	}
	jobDetails := make([]any, 0, len(subJobs))
	for _, job := range subJobs {
		jobDetails = append(jobDetails, &jobDetail{
			SubscribeTopic: job.SubscribeTopic(),
			PublishTopic:   job.PublishTopic(),
		})
	}
	return jobDetails
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
