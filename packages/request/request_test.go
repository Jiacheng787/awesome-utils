package request

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

type RequestPayload struct {
	Mobile  string `json:"mobile"`
	MsgType string `json:"msgType"`
}

const url = "https://open.iconntech.com/unifyUser/sendMsg"

func TestHttpPost(t *testing.T) {
	// 如何快速创建字节流
	// bytes := []byte(`{ "name": "测试内容", "age": 2333 }`)
	payload, err := json.Marshal(&RequestPayload{
		Mobile:  "13777558847",
		MsgType: "01",
	})
	if err != nil {
		log.Fatal(err)
	}

	// 将字符串转为 io.Reader 对象
	// 如果是字节流，也可以使用 bytes.NewBuffer()
	reqBody := strings.NewReader(string(payload))
	request, err := http.NewRequest(http.MethodPost, url, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)
	//status := response.Status
	//header := response.Header
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %s", string(resBody))
}
