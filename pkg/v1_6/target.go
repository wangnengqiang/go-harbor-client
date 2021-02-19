package v1_6

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wangnengqiang/go-harbor-client/pkg/v1_6/models"
	"net/http"
)

const (
	API_PATH_TARGET = "/api/targets"
)

type Target interface {
	CreateTarget(ctx context.Context, target models.ReqTarget) (status bool, err error)
	GetTargets(ctx context.Context, name string) (targets *[]models.Target, err error)
	GetTargetById(ctx context.Context, id int64) (target *models.Target, err error)
}

//This endpoint is for user to create a new replication target.
//registry
func (hc *HarborClient) CreateTarget(ctx context.Context, target models.ReqTarget) (status bool, err error) {
	bytesData, err := json.Marshal(target)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(http.MethodPost, hc.URL+API_PATH_TARGET, reader)
	if err != nil {
		return false, err
	}
	code, body, err := hc.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()
	if code != 201 {
		return false, err
	}
	return true, nil
}

//This endpoint let user list filters targets by name, if name is nil, list returns all targets.
//name The replication's target name.
func (hc *HarborClient) GetTargets(ctx context.Context, name string) (targets *[]models.Target, err error) {
	var ret []models.Target
	path := fmt.Sprintf("%s?name=%s", API_PATH_TARGET, name)
	req, err := http.NewRequest(http.MethodGet, hc.URL+path, nil)
	if err != nil {
		return &ret, err
	}
	err = hc.doJson(ctx, req, &ret)
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}

//This endpoint is for get specific replication's target.
//id The replication's target ID.
func (hc *HarborClient) GetTargetById(ctx context.Context, id int64) (target *models.Target, err error) {
	ret := new(models.Target)
	path := fmt.Sprintf("%s/%d", API_PATH_TARGET, id)
	req, err := http.NewRequest(http.MethodGet, hc.URL+path, nil)
	if err != nil {
		return ret, err
	}
	err = hc.doJson(ctx, req, ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}