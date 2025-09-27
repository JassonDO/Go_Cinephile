package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"
	"voting-golang/blockchain/abis/voting"
	"voting-golang/config"

	"voting-golang/model"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	POOL_CREATED = "0x1d9ed257ba18f079fa73a2be0a204f1883bf6fbea149574f90e7853dda2d13e7"
	POOL_VOTED   = "0x27e4e9f06dee6980758f4150556c677abc1c6ba9832c4be48ce08e0e850e12f4"
)

func main() {

	db := config.MongoClient.Database("voting")

	contractAddress := common.HexToAddress("0x14851Ed2089F3ca46e42195A345D024eBaE48C27")

	latestBlock, err := config.EthClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(5 * time.Second)
	var locker sync.Mutex
	go func() {
		for range ticker.C {
			_latestBlock, err := config.EthClient.BlockByNumber(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}
			locker.Lock()
			latestBlock = _latestBlock
			fmt.Println("Latest block: ", latestBlock.Number().Int64())
			locker.Unlock()
		}
	}()

	var blockNumber = big.NewInt(getCurrentBlock())
	for {
		if blockNumber.Int64() >= latestBlock.Number().Int64()-15 {
			continue
		}

		nextBlock := big.NewInt(blockNumber.Int64())

		if blockNumber.Int64()+4000 < latestBlock.Number().Int64() {
			nextBlock = nextBlock.Add(nextBlock, big.NewInt(4000))
		} else {
			nextBlock = nextBlock.Add(nextBlock, big.NewInt(10))
		}

		fmt.Println("Current block: ", blockNumber.Int64())
		fmt.Println("Next block: ", nextBlock.Int64())

		query := ethereum.FilterQuery{
			FromBlock: blockNumber,
			ToBlock:   nextBlock,
			Addresses: []common.Address{
				contractAddress,
			},
		}

		logs, err := config.EthClient.FilterLogs(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}

		f, err := voting.NewVotingFilterer(
			contractAddress,
			config.EthClient,
		)

		if err != nil {
			log.Fatal(err)
		}

		for _, vLog := range logs {
			for _, topic := range vLog.Topics {
				switch topic.Hex() {
				case POOL_CREATED:
					e, err := f.ParsePoolCreated(vLog)
					if err != nil {
						continue

					}
					poolCreated := model.Pool{
						PoolId:      e.Id.Int64(),
						Name:        e.Name,
						Description: e.Description,
						ClosedAt:    e.ClosedAt.Int64(),
						Timestamp:   e.Timestamp.Int64(),
						Status:      0,
					}

					_, err = db.Collection("pool").InsertOne(context.Background(), poolCreated)
					if err != nil {
						log.Println("Error insert pool created: ", err)
					}
					var candidates []interface{}

					for i, candidate := range e.Cadidates {
						candidates = append(candidates, model.Candidate{
							PoolId:      e.Id.Int64(),
							CandidateId: int64(i),
							Name:        candidate,
						})
					}
					_, err = db.Collection("candidate").InsertMany(context.Background(), candidates)

					if err != nil {
						log.Println("Error insert candidates: ", err)
					}

				case POOL_VOTED:
					e, err := f.ParsePoolVoted(vLog)
					if err != nil {
						continue
					}

					var poolVoted = model.Vote{
						PoolId:     e.Id.Int64(),
						CadidateId: e.CadidateId.Int64(),
						Timestamp:  e.Timestamp.Int64(),
						Name:       e.Name,
						Voter:      strings.ToLower(e.Voter.Hex()),
					}

					_, err = db.Collection("vote").InsertOne(context.Background(), poolVoted)

					if err != nil {
						log.Println("Error insert pool voted: ", err)
					}

					_, err = db.Collection("candidate").UpdateOne(context.Background(),
						bson.M{
							"pool_id":      e.Id.Int64(),
							"candidate_id": e.CadidateId.Int64(),
						}, bson.M{"$inc": bson.M{"vote_count": 1}})
					if err != nil {
						log.Println("Error update candidate: ", err)
					}
				}

			}

		}

		err = saveCurrentBlock(nextBlock.Int64())
		if err != nil {
			log.Fatal(err)
		}

		blockNumber = nextBlock

	}

}

func getCurrentBlock() int64 {
	var metadata model.Metadata
	err := config.MongoClient.Database("voting").Collection("metadata").FindOne(context.Background(), bson.M{}).Decode(&metadata)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			var metadata = model.Metadata{
				FirstBlock:   27916311,
				CurrentBlock: 27916311,
			}
			_, err := config.MongoClient.Database("voting").Collection("metadata").InsertOne(context.Background(), metadata)

			if err != nil {
				log.Fatal(err)
			}

			return metadata.CurrentBlock

		}

		log.Fatal(err)
	}

	return metadata.CurrentBlock
}

func saveCurrentBlock(
	currentBlock int64,
) error {
	_, err := config.MongoClient.Database("voting").Collection("metadata").UpdateOne(context.Background(), bson.M{}, bson.M{"$set": bson.M{"current_block": currentBlock}})
	if err != nil {
		return err
	}

	return nil
}
