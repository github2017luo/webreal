package main

import (
	"log"
	"time"
	"github.com/myxtype/webreal"
)

type testBusiness struct {
	sh *webreal.SubscriptionHub
}

func (t *testBusiness) OnConnect(client *webreal.Client) {
	log.Printf("new client %d, query: %s\n", client.Id(), client.Query().Get("token"))
	// 订阅一个
	client.Subscribe(t.sh, "test")
}

func (t *testBusiness) OnMessage(client *webreal.Client, msg *webreal.Message) {
	log.Printf("client %d new message %v", client.Id(), msg.Data)
}

func (t *testBusiness) OnClose(client *webreal.Client) error {
	log.Printf("close client %d\n", client.Id())
	// 一定得退订，否则有不可预知的错误
	defer client.UnsubscribeAll(t.sh)
	return nil
}

func main() {
	var (
		sh = webreal.NewSubscriptionHub()
		bs = testBusiness{sh: sh}
	)
	// 其他业务来主动推送
	go func() {
		for {
			sh.Publish("test", []byte("test test"))
			time.Sleep(time.Second)
		}
	}()
	webreal.NewServer(&bs).Run("127.0.0.1:8060", "/ws")
}
