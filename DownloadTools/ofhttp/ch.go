package ofhttp

func SendPostRequset(url string, reqChan chan map[string]interface{}, respChan chan map[string]interface{}) {
	for req := range reqChan {
		resp, _ := HttpPost(url, req)
		res, _ := ConvertRespToJson(resp)
		respChan <- res
	}
}
