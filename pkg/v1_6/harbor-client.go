package v1_6

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// harbor客户端结构体
type HarborClient struct {
	Conn *http.Client
	URL  string
	IP   string
	User string
	Pass string
}

// harbor api的实现
type HarborApiImpl interface {
	Label
	Project
	Target
	Policy
	Repository
	Job
}

//NewClient 创建新的Client结构体
func NewHarborClient(cli *http.Client, url, user, pass string) *HarborClient {
	array := strings.SplitN(url, "//", 2)
	if cli == nil {
		switch array[0] {
		case "http:":
			cli = &http.Client{
				Timeout: time.Duration(10) * time.Second,
			}

		case "https:":
			cli = &http.Client{
				Timeout: time.Duration(10) * time.Second,
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}}
		}
	}
	ip := array[1]
	client := &HarborClient{
		Conn: cli,
		URL:  url,
		IP:   ip,
		User: user,
		Pass: pass,
	}
	return client
}

//do 处理请求后的返回值
func (hc *HarborClient) do(ctx context.Context, req *http.Request) (int, io.ReadCloser, error) {
	resp, err := hc.toDo(ctx, req)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, resp.Body, nil
}

//toDo 向harbor发起请求
func (hc *HarborClient) toDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(hc.User, hc.Pass)
	req = req.WithContext(ctx)
	return hc.Conn.Do(req)
}

//doJson 转换返回数据
func (hc *HarborClient) doJson(ctx context.Context, req *http.Request, v interface{}) error {
	code, body, err := hc.do(ctx, req)
	if err != nil {
		return err
	}
	if body != nil {
		defer body.Close()
	}
	if code >= 400 {
		bytes, _ := ioutil.ReadAll(body)
		return fmt.Errorf("http request failed(%d): %s", code, string(bytes))
	}
	err = json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}
