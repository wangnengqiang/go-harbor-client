package v1_6

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wangnengqiang/go-harbor-client/pkg/v1_6/models"
	"net/http"
)

const API_PATH_PROJECT = "/api/projects"

type Project interface {
	CreateProject(ctx context.Context, project models.Project) (bool, error)
	GetProjectByName(ctx context.Context, projectName string) (projects *[]models.Project, err error)
	GetProjectById(ctx context.Context, id int64) (project *models.Project, err error)
	GetProjects(ctx context.Context) (projects *[]models.Project, err error)
}

func (hc *HarborClient) CreateProject(ctx context.Context, project models.Project) (bool, error) {
	createProjectUrl := hc.URL + fmt.Sprintf("%s", API_PATH_PROJECT)
	bytesData, err := json.Marshal(project)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(http.MethodPost, createProjectUrl, reader)
	if err != nil {
		return false, err
	}
	code, body, err := hc.do(ctx, req)
	if body != nil {
		defer body.Close()
	}
	if code == 201 {
		logrus.Info("Project created successfully")
		return true, nil
	} else if code == 409 {
		logrus.Info("Project name already exists.")
		return false, errors.New("project name already exists")
	} else {
		return false, errors.New("create project failed")
	}
}

// harbor中项目名是唯一的,根据projectName查询出来可能是多个
func (hc *HarborClient) GetProjectByName(ctx context.Context, projectName string) (projects *[]models.Project, err error) {
	ret := new([]models.Project)
	// https://10.230.34.226/api/projects?name=wnq
	// name 参数是模糊查询
	// 返回的是一个数组
	projectURL := hc.URL + fmt.Sprintf("%s/?name=%s", API_PATH_PROJECT, projectName)
	req, err := http.NewRequest(http.MethodGet, projectURL, nil)
	if err != nil {
		return ret, err
	}
	err = hc.doJson(ctx, req, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

//This endpoint returns specific project information by project ID.
func (hc *HarborClient) GetProjectById(ctx context.Context, id int64) (project *models.Project, err error) {
	ret := new(models.Project)
	path := fmt.Sprintf("%s/%d", API_PATH_PROJECT, id)
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

//This endpoint returns all projects created by Harbor.
func (hc *HarborClient) GetProjects(ctx context.Context) (projects *[]models.Project, err error) {
	ret := new([]models.Project)

	path := fmt.Sprintf(API_PATH_PROJECT)
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
