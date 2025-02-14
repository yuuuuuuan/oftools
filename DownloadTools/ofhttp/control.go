package ofhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"`
}

type ReqStations struct {
    Category     int    `json:"category"`
    ManufactureId int    `json:"manufactureId"`
    ProjName      string `json:"projName"`
}

type ReqProjects struct {
    Category     int    `json:"category"`
    ManufactureId int    `json:"manufactureId"`
    ProjName      string `json:"projName"`
    Mp           int    `json:"mp"`
}

func ConvertRespToStruct(resp *http.Response) (Response, error) {
	// 检查响应是否为nil
	if resp == nil {
		return Response{}, fmt.Errorf("响应对象为nil")
	}

	// 读取响应体内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("读取响应体失败: %v", err)
	}

	// 解析JSON
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, fmt.Errorf("解析JSON失败: %v", err)
	}

	return response, nil
}
