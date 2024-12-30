package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/lian-yang/steam"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	timeTip, err := steam.GetTimeTip()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Time tip: %#v\n", timeTip)

	timeDiff := time.Duration(timeTip.Time - time.Now().Unix())
	session := steam.NewSession(&http.Client{}, "")
	if err := session.Login(os.Getenv("steamAccount"), os.Getenv("steamPassword"), os.Getenv("steamSharedSecret"), timeDiff); err != nil {
		log.Fatal(err)
	}
	log.Print("Login successful")

	key, err := session.GetWebAPIKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Key: ", key)

	if err = session.ChatLogin(""); err != nil {
		log.Fatal(err)
	}
	defer session.ChatLogoff()

	tries := 0
	for {
		resp, err := session.ChatPoll("10")
		if err != nil {
			log.Printf("chatpoll failed: %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		for _, msg := range resp.Messages {
			sid := steam.SteamID(0)
			sid.ParseDefaults(msg.Partner)

			log.Printf("Message from %d type %s\n", sid, msg.Type)
			if msg.Type == steam.MessageTypeSayText {
				log.Printf("\tText: %s\n", msg.Text)
				if err := session.ChatSendMessage(sid, msg.Text, msg.Type); err != nil {
					log.Printf("Failed to send identical message: %v\n", err)
				}
			}

			if friendState, err := session.ChatFriendState(sid); err != nil {
				log.Printf("failed to get friend state for %d: %v\n", sid, err)
			} else {
				log.Printf("%d: friend state: %#v\n", sid, friendState)
			}
		}

		tries++
		if tries > 10 {
			break
		}

		time.Sleep(time.Second * 2)
	}

	log.Print("Bye")
}
