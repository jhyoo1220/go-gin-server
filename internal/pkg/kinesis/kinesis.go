package kinesis

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"log"
	"strconv"
	"time"
)

var (
	k = kinesis.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-northeast-2")})))
)

func PutRecord(kinesisStream string, data []byte) error {
	output, err := k.PutRecord(&kinesis.PutRecordInput{
		Data:         data,
		PartitionKey: aws.String(strconv.FormatInt(time.Now().UnixNano(), 10)),
		StreamName:   aws.String(kinesisStream),
	})
	if err == nil {
		log.Println(output.GoString())
	}

	return err
}
