package main

import (
	"errors"
	"fmt"
	//	"github.com/blocktree/OpenWallet/log"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
	"net/http"
)

// A Client is a Tron RPC client. It performs RPCs over HTTP using JSON
// request and responses. A Client must be configured with a secret token
// to authenticate with other Cores on the network.
type Client struct {
	BaseURL string
	// AccessToken string
	Debug  bool
	client *req.Req
}

// NewClient create new client to connect
func NewClient(url, token string, debug bool) *Client {
	c := Client{
		BaseURL: url,
		// AccessToken: token,
		Debug: debug,
	}

	api := req.New()
	//trans, _ := api.Client().Transport.(*http.Transport)
	//trans.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c.client = api

	return &c
}


// Call calls a remote procedure on another node, specified by the path.
func (c *Client) Call(path string, param interface{}) (*gjson.Result, error) {

	if c == nil || c.client == nil {
		return nil, errors.New("API url is not setup. ")
	}

	url := c.BaseURL + path
	authHeader := req.Header{"Accept": "application/json"}

	r, err := req.Post(url, req.BodyJSON(&param), authHeader)
	if err != nil {
	//	log.Error("Failed: %+v >\n", err)
		return nil, err
	}

	if r.Response().StatusCode != http.StatusOK {
		message := gjson.ParseBytes(r.Bytes()).String()
		message = fmt.Sprintf("[%s]%s", r.Response().Status, message)
	//	log.Error(message)
		return nil, errors.New(message)
	}

	res := gjson.ParseBytes(r.Bytes())
	return &res, nil
}

func (c * Client) EasyTransferByPrivate(privateKey, toAddress string, amount int)(txhash map[string]string, err error) {

	params := req.Param{
		"privateKey":    privateKey,
		"toAddress": toAddress,
		"amount":     amount ,
	}

	r, err := c.Call("/wallet/easytransferbyprivate", params)
	if err != nil {
		return nil, err
	}

	txhash_hex := r.Get("transaction.txID").String()

	txhash = map[string]string{
		"txhash":    txhash_hex,
	}
	return txhash, nil
}


var quit chan int // 只开一个信道

func foo0(id int) {

	c := NewClient("http://127.0.0.1:16667","",true)
	for i := 1; i <= 600; i++ {
		var privateKey= "a88388f4f53bbc96baed90a0e598794d2d0f4da59d778bb0cb2dc030bec881b5"
		var toAddress= "412bf7b3aae096376c7a74b9462b8257456dc181f4"

		_, err := c.EasyTransferByPrivate(privateKey, toAddress, i+id);
		if err != nil {
			//t.Errorf("EasyTransferByPrivate failed: %v\n", err)
		} else {
			//fmt.Printf("EasyTransferByPrivate return: \n\t%+v\n", r)
		}
	}
	quit <- 0 // ok, finished
}

func foo1(id int) {

	c := NewClient("http://127.0.0.1:17001","",true)

	for i := 1; i <= 600; i++ {
		var privateKey= "a88388f4f53bbc96baed90a0e598794d2d0f4da59d778bb0cb2dc030bec881b5"
		var toAddress= "4165a930a5a294b3cdedced533fcb36a2860a18ee5"

		_, err := c.EasyTransferByPrivate(privateKey, toAddress, i+id);
		if err != nil {
			//t.Errorf("EasyTransferByPrivate failed: %v\n", err)
		} else {
			//fmt.Printf("EasyTransferByPrivate return: \n\t%+v\n", r)
		}
	}
	quit <- 0 // ok, finished
}

func main() {

	count := 300
	quit = make(chan int) // 无缓冲

	for i := 0; i < count; i++ {
		go foo0(i)
		go foo1(i)
		//go foo2(i)
	}
	for i := 0; i < count; i++ {
		<- quit
	}
}

