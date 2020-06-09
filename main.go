package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const appSecret = ""
const jenkinsURL = ""

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"token", "Authorization", "Content-Type", "Origin", "Content-Length", "Access-Control-Allow-Origin"},
		ExposeHeaders:   []string{"token", "Access-Control-Allow-Origin", "Authorization"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/api", OnAliBotPOST)
	r.Run(":9099")
}

// OnAliBotPOST ali bot post message
func OnAliBotPOST(c *gin.Context) {
	localTimestamp := time.Now()
	// nowTimeStampString := strconv.FormatInt(nowTimestamp, 10)

	remoteSign := c.Request.Header.Get("sign")
	remoteTimestamp, err := strconv.ParseInt(c.Request.Header.Get("timestamp"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, AliBotResponse{
			MsgType: "text",
			Text: Content{
				Msg: "time format",
			},
		})
		return
	}

	remoteTime := time.Unix(remoteTimestamp, 0)
	genString := fmt.Sprintf("%d\n%s", remoteTimestamp, appSecret)
	localSign := genSign(genString, appSecret)

	log.Printf("localTime: %d, localSign: %s\nremoteTime: %d, remoteSign: %s\n", localTimestamp.Unix(), localSign, remoteTimestamp, remoteSign)

	if time.Duration(int64(math.Abs(float64(remoteTime.Unix()-localTimestamp.Unix())))) > time.Hour {
		fmt.Println("time error")
		c.AbortWithStatusJSON(http.StatusBadRequest, AliBotResponse{
			MsgType: "text",
			Text: Content{
				Msg: "request error",
			},
		})
		return
	}
	if remoteSign == "" || localSign != remoteSign {
		c.AbortWithStatusJSON(http.StatusBadRequest, AliBotResponse{
			MsgType: "text",
			Text: Content{
				Msg: "request error",
			},
		})
		return
	}
	c.JSON(http.StatusOK, AliBotResponse{
		MsgType: "text",
		Text: Content{
			Msg: "开始构建",
		},
	})
	go notifyBot(context.TODO(), "")
	return
}
func genSign(msg string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(msg))
	// sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
