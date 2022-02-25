package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
)

var (
	url     string
	headers string
	data    []byte
)

func init() {
	guildId := flag.String("guildId", "", "discord guild id")
	clientId := flag.String("clientId", "", "discord client id")
	token := flag.String("token", "TOKEN", "")
	dataPointer := derefString(flag.String("data", "", "json command data"))
	data = []byte(dataPointer)

	url = "https://discord.com/api/v8/applications/" + *clientId + "/guilds/" + *guildId + "/commands"
	headers = `
	{
		"Authorization": "Bot ` + *token + `"
	}
	`
}

func main() {
	buffedData := bytes.NewBuffer(data)
	resp, err := http.Post(url, headers, buffedData)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}
