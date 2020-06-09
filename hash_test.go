package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os/exec"
	"testing"
)

func Test_hash(t *testing.T) {
	cmdbuild()
}

func sign(msg string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(msg))
	// sha := hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func cmdbuild() bool {
	cmd := exec.Command("bash", "-c", "pwd")
	// cmd.Path = "/root/"
	log.Println(cmd)
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
