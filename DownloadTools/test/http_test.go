package test

import (
	"encoding/json"
	"fmt"
	"log"
	"oftools/DownloadTools/ofhttp"
	"reflect"
	"testing"
)

const baseurl = "http://192.168.124.126/client"

func TestAdd(t *testing.T) {
	// 使用t.Run创建子测试用例
	t.Run("add positive numbers", func(t *testing.T) {
		var input, result map[string]interface{}
		var err error
		url := baseurl + "/manufactures"
		input_json := `{}`
		err = json.Unmarshal([]byte(input_json), &input)
		if err != nil {
			log.Fatal(err)
		}
		result_json := `{
			"code": 0,
			"data": {
				"252": "二号园区",
				"254": "未来城园区",
				"256": "巢湖园区",
				"260": "三号园区",
				"339284": "未来城园区VR",
				"339284": "未来城园区AR"
			},
			"msg": "成功"
		}`
		err = json.Unmarshal([]byte(result_json), &result)
		if err != nil {
			log.Fatal(err)
		}
		resp, _ := ofhttp.HttpPost(url, input)
		res, _ := ofhttp.ConvertRespToJson(resp)
		if !DeepEqual(result, res) {
			t.Fatalf("failed")
		}
	})
}

func DeepEqual(a, b map[string]interface{}) bool {
	// 比较两个映射的长度
	if len(a) != len(b) {
		return false
	}

	// 遍历每个键
	for key := range a {
		// 检查键是否存在
		if _, ok := b[key]; !ok {
			return false
		}

		// 获取两个值
		valA := a[key]
		valB := b[key]

		// 比较类型
		if getType(valA) != getType(valB) {
			return false
		}

		// 根据类型比较值
		if !valueEqual(valA, valB) {
			return false
		}

	}
	return true
}

// 获取值的类型
func getType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 比较两个值是否相等
func valueEqual(a, b interface{}) bool {
	// 使用 reflect.DeepEqual 来比较复杂的类型
	return reflect.DeepEqual(a, b)
}
