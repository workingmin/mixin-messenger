package main

import (
	"api/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/MixinNetwork/bot-api-go-client"
)

func main() {
	ctx := context.Background()

	f, err := os.Open("keystore.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, _ := io.ReadAll(f)
	var keystore utils.Keystore
	json.Unmarshal(b, &keystore)

	accessToken, err := bot.SignAuthenticationTokenWithoutBody(keystore.ClientId, keystore.SessionId, keystore.PrivateKey, "GET", "/me")
	if err != nil {
		panic(err)
	}
	user, err := bot.UserMe(ctx, accessToken)
	if err != nil {
		panic(err)
	}

	b, _ = json.Marshal(user)
	fmt.Println(string(b))
}
