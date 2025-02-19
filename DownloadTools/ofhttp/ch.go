package ofhttp

import "log"

func SendPostRequset(url string, reqChan chan map[string]interface{}, respChan chan map[string]interface{}) {
	for req := range reqChan {
		resp, err := HttpPost(url, req)
		if err != nil {
			log.Fatal(err)
		}
		res, err := ConvertRespToJson(resp)
		if err != nil {
			log.Fatal(err)
		}
		respChan <- res
	}
}
