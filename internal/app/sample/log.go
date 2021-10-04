package sample

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jhyoo1220/go-gin-server/internal/pkg/kinesis"
	"log"
	"net/http"
	"strings"
	"time"
)

type sampleLog struct {
	Session          *int64  `form:"session" form:"session" json:"session, number"`
	IP               *string `form:"ip" json:"ip, string"`
	UserAgent        *string `form:"user_agent" json:"user_agent, string"`
	Url              *string `form:"url" json:"url, string"`
	Referrer         *string `form:"referrer" json:"referrer, string"`
	ClientAccessTime *int64  `form:"client_access_time" json:"client_access_time, number"`
	ServerAccessTime *int64  `form:"server_access_time" json:"server_access_time, number"`
	Extra            *string `form:"extra" json:"extra, string"`
}

const (
	PIXEL = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\xFF\x00\xFF\xFF\xFF\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B"
)

func updateValueByHeader(c *gin.Context, key string, target **string) {
	if valSlice, exists := c.Request.Header[key]; exists {
		valStr := strings.Join(valSlice[:], ",")
		*target = &valStr
	} else {
		*target = nil
	}
}

func readLogBytes(c *gin.Context) ([]byte, error) {
	var slBytes []byte
	sl := sampleLog{}
	if err := c.Bind(&sl); err != nil {
		return slBytes, err
	}

	updateValueByHeader(c, "User-Agent", &sl.UserAgent)
	// IP can be nil according to deployment environment
	updateValueByHeader(c, "X-Forwarded-For", &sl.IP)

	currTime := time.Now().UnixNano() / 1e6
	sl.ServerAccessTime = &currTime

	slBytes, err := json.Marshal(sl)
	if err != nil {
		return slBytes, err
	}

	log.Println(string(slBytes))

	return slBytes, nil
}

func ReadLogAndDoNothing(c *gin.Context) {
	_, err := readLogBytes(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Data(200, "image/gif", []byte(PIXEL))
}

func SendToKinesis(c *gin.Context) {
	kinesisStream := c.Param("stream_name")

	slBytes, err := readLogBytes(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

    err = kinesis.PutRecord(kinesisStream, slBytes)
    if err != nil {
		errStr := fmt.Sprintf("Failed to put record to Kinesis: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": errStr})
		return
	}

	c.Data(200, "image/gif", []byte(PIXEL))
}
