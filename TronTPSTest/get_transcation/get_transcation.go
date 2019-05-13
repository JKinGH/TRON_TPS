package main

import (
	"errors"
	"fmt"
	"github.com/blocktree/OpenWallet/log"
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
		log.Error("Failed: %+v >\n", err)
		return nil, err
	}
	// log.Std.Info("%+v", r)

	if r.Response().StatusCode != http.StatusOK {
		message := gjson.ParseBytes(r.Bytes()).String()
		message = fmt.Sprintf("[%s]%s", r.Response().Status, message)
		log.Error(message)
		return nil, errors.New(message)
	}

	res := gjson.ParseBytes(r.Bytes())
	return &res, nil
}


/*
wallet/getnowblock
作用：查询最新块。
demo: curl -X POST  http://127.0.0.1:8090/wallet/getnowblock
参数说明：无
返回值：当前块。
{
    "blockID": "00000000000013d785c9a574621f4ebfa69593bc6cf04157d02fa8709d9e49f9",
    "block_header": {
        "raw_data": {
            "number": 5079,
            "txTrieRoot": "0000000000000000000000000000000000000000000000000000000000000000",
            "witness_address": "41a8beb4082701cc44a4134733c8d2cea4801eb48e",
            "parentHash": "00000000000013d6d2f1b6eee88e03a55d3a9afaf60cac9e85613eb4b4a30c10",
            "version": 7,
            "timestamp": 1555136958000
        },
        "witness_signature": "30c6d2dd6b4c92c80c3391f7d73f081989a432c773df53c22e7e4e7c9f4257d173513b3ffad5e9a2774adbfda2d6465082839d8bf1890064486611bc23de19be01"
    }
}
*/
func (c * Client) GetNowBlock()(blockcount map[string]int64, err error) {

	params := req.Param{
	}

	r, err := c.Call("/wallet/getnowblock", params)
	if err != nil {
		return nil, err
	}

	blocknum := r.Get("block_header.raw_data.number").Int()

	blockcount = map[string]int64{
		"blocknum":    blocknum,

	}
	return blockcount, nil
}

/*
/wallet/gettransactioncountbyblocknum(Odyssey-v3.2开始支持)
作用：查询特定block上transaction的个数
demo: curl -X POST  http://127.0.0.1:8090/wallet/gettransactioncountbyblocknum -d '{"num" : 100}'
参数说明：num是块的高度.
返回值e：transaction的个数.
{"count": 0}
 */

func (c * Client) GetTransactionCountByBlocknum(blocknum int)(tx_num map[string]int64, err error) {

	params := req.Param{
		"num":    blocknum,
	}

	r, err := c.Call("/wallet/gettransactioncountbyblocknum", params)
	if err != nil {
		return nil, err
	}

	transaction_num := r.Get("count").Int()

	tx_num = map[string]int64{
		"num":    transaction_num,

	}
	return tx_num, nil
}


func main() {

/*	sum1 :=  384 +511 +457 +500 +406 +438 +460 +428 +436 +507 +396 +430 +457 +439 +464 +431 +451 +460 +486 +352 +473 +423 +434 +498 +480 +461 +525 +398 +482 +405 +493 +492 +464 +520 +554 +432 +468 +450 +516 +425 +549 +517 +530 +483 +464 +467 +503 +474 +478 +540 +418 +462 +488 +437 +504 +588 +409 +462 +387 +525 +474 +487 +526 +473 +478 +410 +412 +487 +468 +575 +387 +534 +421 +463 +513 +434 +452 +400 +438 +434 +483 +398 +572 +448 +462 +570 +606 +526 +557 +500 +577 +416 +509 +389 +482 +507 +490 +490 +449 +440 +428 +482 +506 +482 +567 +433 +482 +420 +434 +486 +455 +513 +424 +499 +362 +555 +397 +508 +469 +496 +474 +535 +446 +461 +442 +477 +469 +488 +485 +479 +555 +392 +492 +551 +478 +417 +705 +430 +359 +369 +350 +358 +408 +400 +398 +401 +419 +368 +448 +370 +344 +535 +374 +380 +443 +344 +374 +379 +397 +345 +338 +381 +443 +335 +431 +403 +392 +402 +342 +337 +339 +407 +345 +361 +399 +396 +354 +370 +421 +400 +388 +347 +346 +408 +382 +335 +619 +480 +335 +337 +436 +348 +387 +377 +333 +380 +443 +368 +350 +406 +401 +335 +406 +411 +491 +427 +337 +437 +524 +394 +388 +389 +340 +411 +434 +335 +416 +402 +385 +368 +477 +404 +425 +405 +447 +414 +367 +406 +362 +358 +388 +496 +462 +374 +468 +454 +453 +453 +501 +415 +404 +480 +386 +462 +544 +337 +509 +453 +396 +489 +393 +468 +446 +434 +447 +381 +502 +336 +409 +375 +414 +384 +400 +422 +425 +349 +386 +430 +492 +506 +485 +403 +729 +522 +492 +431 +360 +624 +358 +534 +500 +650 +352 +440 +467 +383 +671 +400 +383 +364 +435 +589 +388 +439 +694 +356 +363 +414 +394 +701 +372 +433 +371 +428 +458 +345 +394 +438 +422 +410 +394 +361 +441 +389 +410 +479 +439 +571 +428 +427 +380 +381 +474 +636 +441 +695 +494 +479 +465 +368 +490 +357 +386 +526 +500 +524 +442 +346 +418 +423 +473 +424 +408 +364 +427 +364 +461 +380 +338 +412 +408 +545 +386 +452 +366 +404 +467 +372 +393 +343 +384 +684 +473 +436 +444 +436 +466 +420 +690 +379 +397 +341 +447 +376 +365 +439 +377 +681 +364

	pjz1 := sum1 / 379

	fmt.Println("pjz1=",pjz1)

	sum2 := 705 + 729+ 671+694+701+695+684+690+681

	pjz2 := sum2 / 9

	fmt.Println("pjz2=",pjz2)*/




c := NewClient("http://127.0.0.1:16667","",true)

	now_blockheight,_ := c.GetNowBlock()


	for i := 5000; i <= int(now_blockheight["blocknum"]); i++ {

		r, err := c.GetTransactionCountByBlocknum(i);
		if err != nil {
			//t.Errorf("BroadcastTransaction failed: %v\n", err)
		} else {
			if r["num"] > 1000 && r["num"] < 2000 {

				fmt.Printf(" GetTransactionCountByBlocknum ,height=%d, in range[1000~2000],contain transaction num= %+v ,TPS=%d \n", i, r["num"],r["num"]/3)

			}else if r["num"] > 2000 {
				fmt.Printf(" GetTransactionCountByBlocknum ,height=%d, in range[    >2000],contain transaction num= %+v ,TPS=%d \n", i, r["num"],r["num"]/3)
			}
		}
	}
}


