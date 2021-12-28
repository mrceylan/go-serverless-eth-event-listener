package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EVENT_NAME = os.Getenv("EVENT_NAME")
var SSM_KEY = os.Getenv("SSM_KEY")
var KINESIS_STREAM_NAME = os.Getenv("KINESIS_STREAM_NAME")
var ETH_NODE_URL = os.Getenv("ETH_NODE_URL")
var CONTRACT_ADDRESS = os.Getenv("CONTRACT_ADDRESS")

var awsSession *session.Session

type TransferEvent struct {
	TxHash      string
	BlockHash   string
	BlockNumber uint64
	Val1        string
	Val2        *big.Int
	Val3        bool
}

func main() {
	var err error
	awsSession, err = session.NewSession()
	if err != nil {
		panic(err)
	}
	lambda.Start(getAndPushContractEvents)
	/* callFunc() */
}

func getAndPushContractEvents() {
	ethClient, err := ethclient.Dial(ETH_NODE_URL)
	if err != nil {
		log.Fatal(err)
	}

	lastBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contractInterface, err := abi.JSON(strings.NewReader(string(MainABI)))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	filterLogsQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(getLatestBlockNumber()),
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{contractInterface.Events[EVENT_NAME].ID},
		},
	}
	logs, err := ethClient.FilterLogs(context.Background(), filterLogsQuery)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		var event MainExampleEvent
		err := contractInterface.UnpackIntoInterface(&event, EVENT_NAME, vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		contractEvent := TransferEvent{
			TxHash:      vLog.TxHash.Hex(),
			BlockHash:   vLog.BlockHash.Hex(),
			BlockNumber: vLog.BlockNumber,
			Val1:        event.Val1,
			Val2:        event.Val2,
			Val3:        event.Val3,
		}

		putDataToKinesis(contractEvent)
	}

	setLatestBlockNumber(lastBlockNumber)

}

func getLatestBlockNumber() int64 {
	ssmClient := ssm.New(awsSession)

	param, err := ssmClient.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(SSM_KEY),
		WithDecryption: aws.Bool(false),
	})
	if err != nil {
		log.Fatal(err)
	}

	val, err := strconv.ParseInt(*param.Parameter.Value, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return val
}

func setLatestBlockNumber(val uint64) {
	ssmClient := ssm.New(awsSession)

	_, err := ssmClient.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(SSM_KEY),
		Value:     aws.String(strconv.FormatUint(val, 10)),
		Type:      aws.String(ssm.ParameterTypeString),
		Overwrite: aws.Bool(true),
	})

	if err != nil {
		log.Fatal(err)
	}
}

func putDataToKinesis(data TransferEvent) {
	svc := kinesis.New(awsSession)
	params := &kinesis.PutRecordInput{
		Data:         []byte(EncodeToBytes(data)),
		PartitionKey: aws.String("PartitionKey"),
		StreamName:   aws.String(KINESIS_STREAM_NAME),
	}
	_, err := svc.PutRecord(params)
	if err != nil {
		log.Fatal(err)
	}
}

func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
