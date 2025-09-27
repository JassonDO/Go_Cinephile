package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"voting-golang/config"
	"voting-golang/handler"
	"voting-golang/model"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	log.Println("Scheduler is running...")
	for {
		log.Println("Checking for closed pools...")

		db := config.MongoClient.Database("voting")
		poolCollection := db.Collection("pool")

		var pools []model.Pool

		current := time.Now().Unix()
		cur, err := poolCollection.Find(context.Background(), bson.M{"status": 0, "closed_at": bson.M{"$lte": current}})

		if err != nil {
			log.Println(err)
			time.Sleep(5 * time.Second)
			continue
		}

		for cur.Next(context.Background()) {
			var pool model.Pool
			err := cur.Decode(&pool)
			if err != nil {
				fmt.Println(err)
				continue
			}
			pools = append(pools, pool)
		}

		for _, pool := range pools {

			var votes []model.Vote
			voteCollection := db.Collection("vote")
			cur, err := voteCollection.Aggregate(context.Background(), []bson.M{
				{"$match": bson.M{"pool_id": pool.PoolId}},
				{"$lookup": bson.M{
					"from":         "user",
					"localField":   "voter",
					"foreignField": "address",
					"as":           "user",
				}},
				{"$unwind": bson.M{
					"path": "$user",
				}},
			})
			if err != nil {
				log.Println(err)
				continue
			}

			for cur.Next(context.Background()) {
				var vote model.Vote
				err := cur.Decode(&vote)
				if err != nil {
					fmt.Println(err)
					continue
				}
				votes = append(votes, vote)
			}

			winner, err := handler.GetWinnerByPoolId(context.Background(), pool.PoolId)

			if err != nil {
				fmt.Println(err)
				continue
			}

			body := ""
			html := ""

			if len(winner) == 0 {
				body = fmt.Sprintf("No one won the pool %s", pool.Name)
			} else if len(winner) == 1 {
				body = fmt.Sprintf("The winner of pool %s is %s", pool.Name, winner[0].Name)
			} else {
				body = fmt.Sprintf("The winners of pool %s are %s", pool.Name, winner[0].Name)
				for i := 1; i < len(winner); i++ {
					body = body + fmt.Sprintf(", %s", winner[i].Name)
				}
			}

			for _, vote := range votes {
				subject := "Result of voting pool " + pool.Name
				html = fmt.Sprintf("<p>Dear %s,</p><p>%s</p><p>Thank you for using our service.</p>", vote.User.Name, body)
				err := SendEmail(vote.User.Email, subject, body, html)
				if err != nil {
					log.Println(err)
					continue
				}
				time.Sleep(1 * time.Second)
			}

			_, err = poolCollection.UpdateOne(context.Background(), bson.M{"pool_id": pool.PoolId}, bson.M{"$set": bson.M{"status": 1}})
			if err != nil {
				log.Println(err)
				continue
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func SendEmail(email string, subject string, body string, html string) error {

	to := mail.NewEmail("", email)

	message := mail.NewSingleEmail(config.From, subject, to, body, html)

	response, err := config.SendGridClient.Send(message)

	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("error sending email: %s", response.Body)
	}

	return nil
}
