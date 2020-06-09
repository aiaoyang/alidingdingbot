package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"
)

const webHookURL = "https://oapi.dingtalk.com/robot/send"
const botSecret = ""
const accesstoken = ""

func runCMDJenkinsBuild() bool {
	cmd := exec.Command("sh", "-c", "cd /root/jenkinsjar && java -jar jenkins-cli.jar -s http://10.10.0.15:8080 -auth Jephy:`cat blhx_autopacktoken` build Auto_Pack -v -s")

	log.Println(cmd.Path)
	buf := &bytes.Buffer{}
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		log.Println(err)
		log.Println(buf.String())
		return false
	}
	log.Println(buf.String())
	return true
}

func notifyBot(ctx context.Context, msg string) {
	if runCMDJenkinsBuild() {
		msg = "构建成功"
	} else {
		msg = "构建失败"
	}
	m := strings.NewReader(newMsg(msg))
	req, err := http.NewRequest("POST", webHookURL, m)
	req.URL.RawQuery = genReqValues().Encode()
	req.Header.Set("Content-Type", "application/json")
	log.Printf(req.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	defer client.CloseIdleConnections()
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(res.Body)
	log.Println(buf.String())
}

func genReqValues() url.Values {
	timestamp := time.Now().Unix() * 1000
	src := fmt.Sprintf("%d\n%s", timestamp, botSecret)
	sign := genSign(src, botSecret)
	// signBase64 := base64.StdEncoding.EncodeToString([]byte(sign1))
	value := url.Values{}
	log.Println(sign)
	value.Add("sign", sign)
	value.Add("timestamp", fmt.Sprintf("%d", timestamp))
	value.Add("access_token", accesstoken)
	log.Println(value.Encode())
	return value
}

func newMsg(msg string) string {
	src := AliBotResponse{
		MsgType: "text",
		Text: Content{
			Msg: "构建结果: " + msg,
		},
	}
	jsonStr, err := json.Marshal(src)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", jsonStr)
	return fmt.Sprintf("%s", jsonStr)
}
