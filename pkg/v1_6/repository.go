package v1_6

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wangnengqiang/go-harbor-client/pkg/v1_6/models"
	"github.com/wangnengqiang/go-harbor-client/utils"
	"io/ioutil"
	"net/http"
	"sort"
)

const (
	API_PATH_REPOSITORY = "/api/repositories"
)

type Repository interface {
	GetRepositoriesByProjectId(ctx context.Context, projectId int64) (repositories *[]models.Repository, err error)
	PutRepositoryByName(ctx context.Context, name string, description string) (put bool, err error)
	DeleteRepositoryByName(ctx context.Context, name string) (deleted bool, err error)
	GetLabelByRepositoryName(ctx context.Context, name string) (repositories *[]models.Label, err error)
	PostLabelByRepositoryName(ctx context.Context, name string, label models.Label) (post bool, err error)
	DeleteLabelByRepositoryName(ctx context.Context, name string, id int64) (deleted bool, err error)
	GetTagsByRepositoryName(ctx context.Context, name string) (tags *[]models.DetailedTag, err error)
	PostTagsByRepositoryName(ctx context.Context, name string, tag models.RetagReq) (post bool, err error)
	GetTagByRepositoryName(ctx context.Context, name, tag string) (detailedTag *models.DetailedTag, err error)
	PostTagsLabelByRepositoryName(ctx context.Context, name, tag string, label models.Label) (post bool, err error)
	DeleteTagsLabelByRepositoryName(ctx context.Context, name, tag string, id int64) (deleted bool, err error)
	GetTagsOfRepository(ctx context.Context, registryName string) (tagReps *[]models.TagResp, err error)
	AddLabelToImage(ctx context.Context, repository, tagName string, labelId int64) (status bool, err error)
	GetRepositoriesByNameAndProjectId(ctx context.Context, projectId int64, repositoryName string) (repositories *[]models.Repository, err error)
	GerLastTagsOfRepository(ctx context.Context, registryName string, num int) (tagReps *[]models.TagResp, err error)
	TagExists(ctx context.Context, repositoryFullName, tagName string) bool
}

//This endpoint lets user search repositories accompanying with relevant project ID and repo name.
//Repositories can be sorted by repo name, creation_time, update_time in either ascending or descending order.
//id 项目ID
func (hc *HarborClient) GetRepositoriesByProjectId(ctx context.Context, projectId int64) (repositories *[]models.Repository, err error) {
	ret := new([]models.Repository)
	path := fmt.Sprintf("%s?project_id=%d", API_PATH_REPOSITORY, projectId)
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

//This endpoint is used to update description of the repository.
//name 仓库的名称
func (hc *HarborClient) PutRepositoryByName(ctx context.Context, name string, description string) (put bool, err error) {
	bytesData, err := json.Marshal(struct {
		Description string `json:"description"`
	}{description})
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	path := fmt.Sprintf("%s/%s", API_PATH_REPOSITORY, name)
	req, err := http.NewRequest(http.MethodPut, hc.URL+path, reader)
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

//This endpoint let user delete a repository with name.
//name 仓库的名称
func (hc *HarborClient) DeleteRepositoryByName(ctx context.Context, name string) (deleted bool, err error) {
	path := fmt.Sprintf("%s/%s", API_PATH_REPOSITORY, name)
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

//Get labels of a repository specified by the repo_name.
//name 仓库的名称
func (hc *HarborClient) GetLabelByRepositoryName(ctx context.Context, name string) (repositories *[]models.Label, err error) {
	ret := new([]models.Label)
	path := fmt.Sprintf("%s/%s/labels", API_PATH_REPOSITORY, name)
	fmt.Println(path)
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

//Add a label to the repository.
//name 仓库的名称
//label 标签结构体
func (hc *HarborClient) PostLabelByRepositoryName(ctx context.Context, name string, label models.Label) (post bool, err error) {
	bytesData, err := json.Marshal(label)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	path := fmt.Sprintf("%s/%s/labels", API_PATH_REPOSITORY, name)
	req, err := http.NewRequest(http.MethodPost, hc.URL+path, reader)
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

//Delete the label from the repository specified by the repo_name.
//name 仓库的名称
//id label的ID
func (hc *HarborClient) DeleteLabelByRepositoryName(ctx context.Context, name string, id int64) (deleted bool, err error) {
	path := fmt.Sprintf("%s/%s/labels/%d", API_PATH_REPOSITORY, name, id)
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

//This endpoint aims to retrieve tags from a relevant repository.
//If deployed with Notary, the signature property of response represents whether the image is singed or not.
//If the property is null, the image is unsigned.
//name 仓库的名称
//tag  image的标签
func (hc *HarborClient) GetTagsByRepositoryName(ctx context.Context, name string) (tags *[]models.DetailedTag, err error) {
	ret := new([]models.DetailedTag)
	path := fmt.Sprintf("%s/tags/%s", API_PATH_REPOSITORY, name)
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

//This endpoint tags an existing image with another tag in this repo, source images can be in different repos or projects.
//name 仓库的名称
//tag  image添加的标签
func (hc *HarborClient) PostTagsByRepositoryName(ctx context.Context, name string, tag models.RetagReq) (post bool, err error) {
	bytesData, err := json.Marshal(tag)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	path := fmt.Sprintf("%s/%s/tags", API_PATH_REPOSITORY, name)
	req, err := http.NewRequest(http.MethodPost, hc.URL+path, reader)
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

//This endpoint aims to retrieve the tag of the repository.
//If deployed with Notary, the signature property of response represents whether the image is singed or not.
//If the property is null, the image is unsigned.
//name 仓库的名称
//tag  image的标签
func (hc *HarborClient) GetTagByRepositoryName(ctx context.Context, name, tag string) (detailedTag *models.DetailedTag, err error) {
	ret := new(models.DetailedTag)
	path := fmt.Sprintf("%s/%s/tags/%s", API_PATH_REPOSITORY, name, tag)
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

//This endpoint let user delete tags with repo name and tag.
//name 仓库的名称
//tag  仓库的标签
func (hc *HarborClient) DeleteTagByRepositoryName(ctx context.Context, name string, tag string) (deleted bool, err error) {
	path := fmt.Sprintf("%s/%s/tags/%s", API_PATH_REPOSITORY, name, tag)
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

//Get labels of an image specified by the repo_name and tag.
//根据RepositoryName获取Repository的tag的label
//name 仓库的名称
//tag  image的标签
func (hc *HarborClient) GetTagsLabelByRepositoryName(ctx context.Context, name, tag string) (label *[]models.Label, err error) {
	ret := new([]models.Label)
	path := fmt.Sprintf("%s/%s/tags/%s/labels", API_PATH_REPOSITORY, name, tag)
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

//This endpoint tags an existing image with another tag in this repo, source images can be in different repos or projects.
//name 仓库的名称
//tag image的标签
func (hc *HarborClient) PostTagsLabelByRepositoryName(ctx context.Context, name, tag string, label models.Label) (post bool, err error) {
	bytesData, err := json.Marshal(label)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	path := fmt.Sprintf("%s/%s/tags/%s/labels", API_PATH_REPOSITORY, name, tag)
	req, err := http.NewRequest(http.MethodPost, hc.URL+path, reader)
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

//Delete the label from the image specified by the repo_name and tag.
//name 仓库的名称
//tag  仓库的标签
//id 标签ID
func (hc *HarborClient) DeleteTagsLabelByRepositoryName(ctx context.Context, name, tag string, id int64) (deleted bool, err error) {
	path := fmt.Sprintf("%s/%s/tags/%s/labels/%d", API_PATH_REPOSITORY, name, tag, id)
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

func (hc *HarborClient) GetTagsOfRepository(ctx context.Context, registryName string) (tagReps *[]models.TagResp, err error) {
	ret := new([]models.TagResp)
	tagsURL := hc.URL + fmt.Sprintf("%s/%s/tags/", API_PATH_REPOSITORY, registryName)
	req, err := http.NewRequest(http.MethodGet, tagsURL, nil)
	if err != nil {
		return ret, nil
	}
	code, body, err := hc.do(ctx, req)
	if err != nil {
		return ret, err
	}
	if code == 200 {
		bytes, _ := ioutil.ReadAll(body)
		err = json.Unmarshal(bytes, &ret)
		if err != nil {
			fmt.Println("errors:", err)
		}
	}
	defer body.Close()
	return ret, nil
}
func (hc *HarborClient) AddLabelToImage(ctx context.Context, repository, tagName string, labelId int64) (status bool, err error) {
	//"https://10.230.34.226/api/repositories/wnq/nginx-photon/tags/v1/labels"
	labelURL := hc.URL + fmt.Sprintf("%s/%s/tags/%s/labels", API_PATH_REPOSITORY, repository, tagName)
	label, err := hc.GetLabelById(ctx, labelId)
	if err != nil {
		return false, err
	}
	bytesData, err := json.Marshal(label)
	if err != nil {
		return false, err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(http.MethodPost, labelURL, reader)
	if err != nil {
		return false, nil
	}
	_, body, err := hc.do(ctx, req)
	if err != nil {
		return false, err
	}
	defer body.Close()
	return true, nil
}

func (hc *HarborClient) GetRepositoriesByNameAndProjectId(ctx context.Context, projectId int64, repositoryName string) (repositories *[]models.Repository, err error) {
	ret := new([]models.Repository)
	path := fmt.Sprintf("%s?project_id=%d&q=%s", API_PATH_REPOSITORY, projectId, repositoryName)
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

func (hc *HarborClient) GerLastTagsOfRepository(ctx context.Context, registryName string, num int) (tagReps *[]models.TagResp, err error) {
	ret := new([]models.TagResp)
	tagsURL := hc.URL + fmt.Sprintf("%s/%s/tags/", API_PATH_REPOSITORY, registryName)
	req, err := http.NewRequest(http.MethodGet, tagsURL, nil)
	if err != nil {
		return ret, nil
	}
	code, body, err := hc.do(ctx, req)
	if err != nil {
		return ret, err
	}
	if code == 200 {
		data, _ := ioutil.ReadAll(body)
		err = json.Unmarshal(data, &ret)
		if err != nil {
			return ret, err
		}
		defer body.Close()
		// 判断返回的tag的数量与num的关系
		if len(*ret) >= num {
			rTags := utils.TimeSort{}
			for _, t := range *ret {
				rTags.Slice = append(rTags.Slice, t)
			}
			rTags.By = tagsTimeBy
			sort.Sort(rTags)
			sortTags := rTags.Slice
			if len(sortTags) > num {
				sortTags = sortTags[:num]
			}
			var repTags []models.TagResp
			for _, d := range sortTags {
				repTags = append(repTags, d.(models.TagResp))
			}
			return &repTags, nil
		} else {
			return ret, nil
		}

	} else {
		return ret, errors.New("获取 tag 失败")
	}

}

func tagsTimeBy(a interface{}, b interface{}) bool {
	return a.(models.TagResp).Name > (b.(models.TagResp).Name)
}

func (hc *HarborClient) TagExists(ctx context.Context, repositoryFullName, tagName string) bool {
	tags, _ := hc.GetTagsOfRepository(ctx, repositoryFullName)
	flag := false
	if *tags != nil {
		for _, tag := range *tags {
			if tagName == tag.Name {
				flag = true
				break
			}
		}
	}
	return flag
}
