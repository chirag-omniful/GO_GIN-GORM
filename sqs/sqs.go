package sqs

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

var sqsSvc *sqs.SQS
var queueUrl string

func SqsSession() {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	if err != nil {
		fmt.Println("Failed to initialize session")
		return
	}

	sqsSvc = sqs.New(sess)

	CreateQueue("taskQueue")

	temp, err := json.Marshal("Chirag_is_working_on_sqs")
	if err != nil {
		log.Fatal("Error in marshalling data")
		return
	}

	err = SendMessageToSqs(temp)
	if err != nil {
		log.Fatal("Error sending message")
	}

	RecieveMessageToSqs()
}

func CreateQueue(name string) {
	result, err := sqsSvc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: aws.String(name),
	})

	if err != nil {
		log.Fatal("Queue not created")
	}

	queueUrl = aws.StringValue(result.QueueUrl)
	log.Println(queueUrl)
}

func SendMessageToSqs(msg []byte) (err error) {
	_, err = sqsSvc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(msg)),
		QueueUrl:    &queueUrl,
	})
	if err != nil {
		log.Println("Could not send message")
		return
	}
	return
}

func RecieveMessageToSqs() {
	msgResult, err := sqsSvc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(1),
	})

	if err != nil {
		log.Fatal("Error receiving message")
	}

	fmt.Println("Message Body: " + *msgResult.Messages[0].Body)

}
