package libs

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/resty.v1"
)

func getRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func JsonRpcRequest(method string, param string, res interface{}) (*resty.Response, error) {
	requestId := getRandomNumber()
	body := fmt.Sprintf(`{
			"id": %d,
			"method":"%s",
			"params":[
				%s
			],
			"jsonrpc":"2.0"
		}`,
		requestId,
		method,
		param,
	)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(res).
		Post(rpcEndpoint)
	return resp, err
}
