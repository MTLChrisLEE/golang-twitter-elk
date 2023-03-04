package main

import (
	"fmt"
	"twitter-to-kafka/packages/config"
)

func main() {

	// vp := viper.New()

	// vp.SetConfigFile("./test.json")
	// err := vp.ReadInConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(vp.GetString("foo"))

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(config.Twitter.API_KEY)
		fmt.Println(config.Twitter.SECRET)
	}

	// tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret("key", "secret").RequestBearerToken()

}
