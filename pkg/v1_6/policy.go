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
	"strconv"
)

const (
	API_PATH_POLICY      = "/api/policies/replication"
	API_PATH_REPLICATION = "/api/replications"
)

type Policy interface {
	GetPolicies(ctx context.Context, policyName string, projectId int64) (policies *[]models.Policy,
		err error)
	CreateReplicationPolicy(ctx context.Context, policy models.ReqPolicy) (status bool, err error)
	CreatePolicy(ctx context.Context, policy models.HarborPolicy) (status bool, err error)
	DeletePolicyById(ctx context.Context, id int64) (delete bool, err error)
	GetReplicationPolicyById(ctx context.Context, policyId int64) (policy *models.Policy, error error)
	// 执行复制测率
	TriggerReplicationPolicy(ctx context.Context, policyId int64) (status bool, err error)
}

//This endpoint let user list filters policies by name and project_id, if name and project_id are nil, list returns all policies.
//name The replication's policy name.
//id Relevant project ID.
func (hc *HarborClient) GetPolicies(ctx context.Context, policyName string, projectId int64) (policies *[]models.Policy,
	err error) {
	var ret []models.Policy
	// https://harbor.deppontest.com/api/policies/replication?name=grayscale&project_id=10491
	path := fmt.Sprintf("%s", API_PATH_POLICY)
	policiesUrl := hc.URL + path
	if policyName != "" && projectId != 0 {
		policiesUrl = policiesUrl + "?name=" + policyName + "&project_id=" + strconv.FormatInt(projectId, 10)
	} else if policyName == "" && projectId == 0 {

	} else if policyName != "" && projectId == 0 {
		policiesUrl = policiesUrl + "?name=" + policyName
	} else if policyName == "" && projectId != 0 {
		policiesUrl = policiesUrl + "?project_id=" + strconv.FormatInt(projectId, 10)
	}
	req, err := http.NewRequest(http.MethodGet, policiesUrl, nil)
	if err != nil {
		return &ret, err
	}
	err = hc.doJson(ctx, req, &ret)
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}

func (hc *HarborClient) CreateReplicationPolicy(ctx context.Context, repPolicy models.ReqPolicy) (status bool, err error) {
	// 判断target个数，如果大于1需要处理,一个policy不支持有多个target,默认只会添加target数组中的第一个
	addPolicyURL := hc.URL + fmt.Sprintf("%s/", API_PATH_POLICY)
	policyJsonData, err := json.Marshal(repPolicy)
	if err != nil {
		return false, errors.New("policy对象解析失败")
	}
	policyByteData := bytes.NewReader(policyJsonData)
	req, err := http.NewRequest(http.MethodPost, addPolicyURL, policyByteData)
	code, body, err := hc.do(ctx, req)
	if body != nil {
		defer body.Close()
	}
	if err != nil {
		return false, errors.New("policy创建失败")
	}
	if code == 201 {
		return true, nil
	} else {
		bytes, _ := ioutil.ReadAll(body)
		return false, errors.New(string(bytes))
	}

}

func (hc *HarborClient) CreatePolicy(ctx context.Context, policy models.HarborPolicy) (status bool, err error) {
	createPolicyUrl := hc.URL + fmt.Sprintf("%s/", API_PATH_POLICY)
	policyJsonData, err := json.Marshal(policy)
	if err != nil {
		return false, errors.New("policy对象解析失败")
	}
	policyByteData := bytes.NewReader(policyJsonData)
	req, err := http.NewRequest(http.MethodPost, createPolicyUrl, policyByteData)
	code, body, err := hc.do(ctx, req)
	if body != nil {
		defer body.Close()
	}
	if err != nil {
		return false, errors.New("policy创建失败")
	}
	if code == 201 {
		return true, nil
	} else {
		bytes, _ := ioutil.ReadAll(body)
		return false, errors.New(string(bytes))
	}
}

func (hc *HarborClient) DeletePolicyById(ctx context.Context, id int64) (delete bool, err error) {
	path := fmt.Sprintf("%s/%d", API_PATH_POLICY, id)
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

func (hc *HarborClient) GetReplicationPolicyById(ctx context.Context, policyId int64) (policy *models.Policy, error error) {
	var ret models.Policy
	policyUrl := hc.URL + fmt.Sprintf("%s/%d", API_PATH_POLICY, policyId)
	req, err := http.NewRequest(http.MethodGet, policyUrl, nil)
	if err != nil {
		return &ret, err
	}
	err = hc.doJson(ctx, req, &ret)
	if err != nil {
		return &ret, err
	}
	return &ret, nil
}

// 执行复制策略
func (hc *HarborClient) TriggerReplicationPolicy(ctx context.Context, policyId int64) (status bool, err error) {
	triggerURL := hc.URL + fmt.Sprintf("%s", API_PATH_REPLICATION)
	// { "policy_id": 0 }  在body里面传递policy_id
	strPolicyId := strconv.FormatInt(policyId, 10)
	jsonData := "{\"policy_id\":" + strPolicyId + "}"
	buf := bytes.NewBufferString(jsonData)
	req, err := http.NewRequest(http.MethodPost, triggerURL, buf)
	if err != nil {
		return false, err
	}
	code, body, err := hc.do(ctx, req)
	data, err := ioutil.ReadAll(body)
	if code == 200 {
		return true, nil
	} else {
		fmt.Println(string(data), err)
		return false, err
	}
}
