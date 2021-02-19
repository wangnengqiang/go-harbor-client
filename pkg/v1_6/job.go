package v1_6

import (
	"context"
	"fmt"
	"github.com/wangnengqiang/go-harbor-client/pkg/v1_6/models"
	"net/http"
)

const (
	API_PATH_JOB = "/api/jobs/replication"
)

type Job interface {
	GetJobStatusOfRepositoryPolicy(ctx context.Context, policyId int64) (jobStatuses *[]models.Job, err error)
}

func (hc *HarborClient) GetJobStatusOfRepositoryPolicy(ctx context.Context, policyId int64) (jobStatuses *[]models.Job, err error) {
	job := new([]models.Job)
	// https://harbor.deppontest.com/api/jobs/replication?policy_id=145
	jobURL := hc.URL + fmt.Sprintf("%s?policy_id=%d", API_PATH_JOB, policyId)
	req, err := http.NewRequest(http.MethodGet, jobURL, nil)
	if err != nil {
		return job, err
	}
	err = hc.doJson(ctx, req, &job)
	if err != nil {
		return job, err
	}
	return job, nil
}
