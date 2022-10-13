package websocket

import (
	"context"
	"fmt"
	"log"
	"swap_bot/libs"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Define a scaffold
type Websocket struct {
	*ethclient.Client
	query ethereum.FilterQuery
}

// Initialization function
func Init() *Websocket {
	client, err := ethclient.Dial(WS_URL)
	if err != nil {
		log.Fatal(err)
	}

	websocket := &Websocket{Client: client, query: ethereum.FilterQuery{
		Addresses: []common.Address{
			common.HexToAddress(PANCAKE_FACTORY_ADDRESS),
		},
		Topics: [][]common.Hash{{
			common.HexToHash(CREATE_PAIR_TOPIC),
		}},
	}}

	return websocket
}

func (websocket *Websocket) Connect() {
	channel := make(chan types.Log)
	sub, err := websocket.SubscribeFilterLogs(context.Background(), websocket.query, channel)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-channel:
			if DEBUG {
				fmt.Println(vLog)
			}
			wg.Add(1)
			go func() {
				handleEvent(vLog)
				wg.Done()
			}()
		}
	}
}

type PairData struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type JsonResponse struct {
	Result []PairData `json:"result"`
}

func handleEvent(vLog types.Log) {
	token1 := common.HexToAddress(vLog.Topics[1].Hex())
	token2 := common.HexToAddress(vLog.Topics[2].Hex())

	checkToken1BNB := token1.String() == BNB_ADDRESS
	checkToken1BUSD := token1.String() == BUSD_ADDRESS
	checkToken2BNB := token2.String() == BNB_ADDRESS
	checkToken2BUSD := token2.String() == BUSD_ADDRESS

	if DEBUG || checkToken1BNB || checkToken2BNB || checkToken1BUSD || checkToken2BUSD {
		pairAddress := common.HexToAddress(hexutil.Encode(vLog.Data[12:32]))
		blockNumber := hexutil.EncodeUint64(vLog.BlockNumber)
		bodyParams := fmt.Sprintf(`
					{
						"address": ["%s"],
						"fromBlock": "%s",
						"toBlock": "%s",
						"topics": ["%s"]
					}`,
			pairAddress,
			blockNumber,
			blockNumber,
			PAIR_MINT_TOPIC,
		)
		responseResult := GetLogPairData(bodyParams)
		if responseResult != nil && len(responseResult.Result) > 0 {
			amount1 := responseResult.Result[0].Topics[1]
			amount2 := responseResult.Result[0].Data[0:66]

			if DEBUG || checkToken1BNB && amount1 >= BNB_MINIMUM_REQUIRED_AMOUNT ||
				checkToken1BUSD && amount1 >= BUSD_MINIMUM_REQUIRED_AMOUNT ||
				checkToken2BNB && amount2 >= BNB_MINIMUM_REQUIRED_AMOUNT ||
				checkToken2BUSD && amount2 >= BUSD_MINIMUM_REQUIRED_AMOUNT {
				libs.RecordData(token1.String(), amount1, token2.String(), amount2, blockNumber)
			}
		}
	}
}

func GetLogPairData(bodyParams string) *JsonResponse {
	response, err := libs.JsonRpcRequest(
		"eth_getLogs",
		bodyParams,
		&JsonResponse{},
	)

	if err != nil {
		log.Println("JsonResponse")
		log.Println(err)
		return nil
	}

	return response.Result().(*JsonResponse)
}
