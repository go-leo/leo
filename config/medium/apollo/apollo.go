package apollo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ConfigResponse apollo配置响应
type ConfigResponse struct {
	AppID          string         `json:"appId"`
	Cluster        string         `json:"cluster"`
	NamespaceName  string         `json:"namespaceName"`
	ReleaseKey     string         `json:"releaseKey"`
	Configurations map[string]any `json:"configurations"`
}

// Notification 用于保存 apollo Notification 信息
type Notification struct {
	NamespaceName  string `json:"namespaceName"`
	NotificationID int64  `json:"notificationId"`
}

func getConfigContent(ctx context.Context, uri string, appID, secret, contentType string, client *http.Client) (string, error) {
	var resp ConfigResponse
	rawResp, err := requestApollo(ctx, uri, appID, secret, client)
	if err != nil {
		return "", err
	}
	defer rawResp.Body.Close()
	if rawResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed get config from apollo, StatusCode: %d, Status: %s", rawResp.StatusCode, rawResp.Status)
	}
	data, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	var content string
	switch contentType {
	case "yaml", "yml":
		if val, ok := resp.Configurations["content"]; ok {
			content, ok = val.(string)
			if !ok {
				return "", fmt.Errorf("content is not string, %v, %T", val, val)
			}
		}
	}
	return content, nil
}

func requestApollo(ctx context.Context, uri, appID, secret string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	authSignatureHeaders := authSignatureHeaders(uri, appID, secret)
	for key, val := range authSignatureHeaders {
		req.Header.Set(key, val)
	}
	return client.Do(req)
}
