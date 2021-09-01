package jpushclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
	HOST_SCHEDULE = "https://api.jpush.cn/v3/schedules"
	HOST_REPORT   = "https://report.jpush.cn/v3/received"
	BASE64_TABLE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{secret, appKey, auth, HOST_NAME_SSL}
	return pusher
}

func (client *PushClient) Send(data []byte) (string, error) {
	return client.SendPushBytes(data)
}
func (client *PushClient) CreateSchedule(data []byte) (string, error) {
	// client.BaseUrl = HOST_SCHEDULE
	return client.SendScheduleBytes(data, HOST_SCHEDULE)
}
func (client *PushClient) DeleteSchedule(id string) (string, error) {
	// client.BaseUrl = HOST_SCHEDULE
	return client.SendDeleteScheduleRequest(id, HOST_SCHEDULE)
}
func (client *PushClient) GetSchedule(id string) (string, error) {
	// GET https://api.jpush.cn/v3/schedules/{schedule_id}
	// client.BaseUrl = HOST_SCHEDULE
	return client.SendGetScheduleRequest(id, HOST_SCHEDULE)

}
func (client *PushClient) GetReport(msg_ids string) (string, error) {
	// client.BaseUrl = HOST_REPORT
	return client.SendGetReportRequest(msg_ids, HOST_REPORT)
}
func (client *PushClient) SendPushString(content string) (string, error) {
	ret, err := SendPostString(client.BaseUrl, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (client *PushClient) SendPushBytes(content []byte) (string, error) {
	//ret, err := SendPostBytes(client.BaseUrl, content, client.AuthCode)
	ret, err := SendPostBytes2(client.BaseUrl, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (client *PushClient) SendScheduleBytes(content []byte, url string) (string, error) {
	ret, err := SendPostBytes2(url, content, client.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}

}

func (client *PushClient) SendGetReportRequest(msg_ids string, url string) (string, error) {
	return Get(url).SetBasicAuth(client.AppKey, client.MasterSecret).Param("msg_ids", msg_ids).String()
}

func UnmarshalResponse(rsp string) (map[string]interface{}, error) {
	mapRs := map[string]interface{}{}
	if len(strings.TrimSpace(rsp)) == 0 {
		return mapRs, nil
	}
	err := json.Unmarshal([]byte(rsp), &mapRs)
	if err != nil {
		return nil, err
	}
	if _, ok := mapRs["error"]; ok {
		return nil, fmt.Errorf(rsp)
	}
	return mapRs, nil
}

func (client *PushClient) SendDeleteScheduleRequest(scheduleId string, url string) (string, error) {
	rsp, err := Delete(strings.Join([]string{url, scheduleId}, "/")).Header("Authorization", client.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
func (client *PushClient) SendGetScheduleRequest(scheduleId string, url string) (string, error) {
	rsp, err := Get(strings.Join([]string{url, scheduleId}, "/")).Header("Authorization", client.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
