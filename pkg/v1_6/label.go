package v1_6

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wangnengqiang/go-harbor-client/pkg/v1_6/models"
	"io/ioutil"
	"net/http"
)

const (
	API_PATH_LABEL = "/api/labels"
)

type Label interface {
	CreateLabel(ctx context.Context, label models.Label) (status bool, err error)
	GetLabelById(ctx context.Context, id int64) (label *models.Label, err error)
	GetLabelsByProjectId(ctx context.Context, name, scope string, id int64) (labels *[]models.Label,
		err error)
	DeleteLabelById(ctx context.Context, id int64) (deleted bool, err error)
	GetLabelsByLabelNameAndProjectId(ctx context.Context, labelName string, projectId int64) (label *[]models.Label, err error)
}

func (hc *HarborClient) CreateLabel(ctx context.Context, label models.Label) (bool, error) {
	createLabelURL := hc.URL + API_PATH_LABEL
	bytesData, err := json.Marshal(label)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(http.MethodPost, createLabelURL, reader)
	if err != nil {
		return false, err
	}
	code, body, err := hc.do(ctx, req)
	if err != nil {
		return false, err
	}
	datas, _ := ioutil.ReadAll(body)
	defer body.Close()
	// 201 Create successfully
	// 409 Label with the same name and same scope already exists.
	if code != 201 {
		return false, errors.New(string(datas))
	}
	return true, nil
}

func (hc *HarborClient) GetLabelById(ctx context.Context, id int64) (label *models.Label, err error) {
	ret := new(models.Label)
	path := fmt.Sprintf("%s/%d", API_PATH_LABEL, id)
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

//This endpoint let user list labels by name, scope and project_id.
//name The label name.
//scope The label scope. Valid values are g and p. g for global labels and p for project labels.
//id Relevant project ID, required when scope is p.
func (hc *HarborClient) GetLabelsByProjectId(ctx context.Context, name, scope string, id int64) (labels *[]models.Label,
	err error) {
	var ret []models.Label
	path := fmt.Sprintf("%s?scope=%s&name=%s&project_id=%d", API_PATH_LABEL, scope, name, id)
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

//Delete the label specified by ID.
//id Label ID.
func (hc *HarborClient) DeleteLabelById(ctx context.Context, id int64) (deleted bool, err error) {
	path := fmt.Sprintf("%s/%d", API_PATH_LABEL, id)
	req, err := http.NewRequest(http.MethodDelete, hc.URL+path, nil)
	if err != nil {
		return false, err
	}
	_, body, err := hc.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()
	return true, nil
}

// 传入的labelName唯一才能获取唯一的label
func (hc *HarborClient) GetLabelsByLabelNameAndProjectId(ctx context.Context, labelName string, projectId int64) (label *[]models.Label,
	err error) {
	var ret []models.Label
	labelUrl := hc.URL + fmt.Sprintf("%s?name=%s&scope=p&project_id=%d", API_PATH_LABEL, labelName, projectId)
	req, err := http.NewRequest(http.MethodGet, labelUrl, nil)
	if err != nil {
		return &ret, err
	}
	err = hc.doJson(ctx, req, &ret)
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}
