package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/pubsub"
)

func Route(rg *gin.RouterGroup, pubsubTask *pubsub.Task) {
	rg.GET("/pubsub/jobs", func(c *gin.Context) {
		c.JSON(http.StatusOK, subJobDetails(pubsubTask.Jobs()))
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
