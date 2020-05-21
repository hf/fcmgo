package fcmgo

import (
	"bytes"
	"context"
	"encoding/json"
	"golang.org/x/net/context/ctxhttp"
	"net/http"
)

// An FCM client. Just initialize this struct, and you'll be good to go.
type Client struct {
	Authorization string
	Client        *http.Client
}

type DirectMessage struct {
	To   string      `json:"to"`
	Data interface{} `json:"data"`
}

type ResponseData struct {
}

type Response struct {
	HTTP *http.Response
	Data ResponseData
}

func (fcm Client) Send(ctx context.Context, msg *DirectMessage) (*Response, error) {
	reqBody, err := json.Marshal(msg)
	if nil != err {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(reqBody))
	if nil != err {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+fcm.Authorization)

	res, err := ctxhttp.Do(ctx, fcm.Client, req)
	if nil != err {
		return nil, err
	}

	response := Response{
		HTTP: res,
		Data: ResponseData{},
	}

	err = json.NewDecoder(res.Body).Decode(&response.Data)
	if nil != err {
		return nil, err
	}

	return &response, nil
}
