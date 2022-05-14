package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
)

var (
	clientId  = flag.String("clientId", "", "discord client id")
	guildId   = flag.String("guildId", "", "discord guild id")
	token     = flag.String("token", "TOKEN", "")
	slashData = derefString(flag.String("data", "", "json command data"))
	url       = "https://discord.com/api/v8/applications/" + *clientId + "/guilds/" + *guildId + "/commands"
	headers   = `
	{
		"Authorization": "Bot ` + *token + `"
	}`
	data      = []byte(slashData)
)

func init() {
	flag.Parse()
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

func main() {
	buffedData := bytes.NewBuffer(data)
	resp, err := http.Post(url, headers, buffedData)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}

