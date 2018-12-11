package handlers

import (
	"github.com/go-land/job-service/proto"
	"golang.org/x/net/context"
	"log"
)

type JobService struct {
	jobs map[string]string
}

func (service *JobService) GetJob(ctx context.Context, request *job.GetJobRequest, resp *job.JobResponse) error {

	log.Printf("GetJob called for alias %s\n", request.Name)

	jobValue, ok := service.jobs[request.Name]

	if !ok {
		resp.Job = "<undefined>"
	} else {
		resp.Job = jobValue
	}

	return nil
}

func NewJobHandler() *JobService {

	jobHandler := JobService{}

	jobHandler.jobs = make(map[string]string)

	jobHandler.jobs["maksym"] = "Senior EM"
	jobHandler.jobs["olesia"] = "Housekeeper"

	return &jobHandler
}
