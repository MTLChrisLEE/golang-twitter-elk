package main

import (
	"log"
	"net/http"
	"os"
	"twitter-to-kafka/packages/handlers"
	// twitterstream "github.com/fallenstedt/twitter-stream"
)

func main() {

	l := log.New(os.Stdout, "data-api", log.LstdFlags)
	newHandler := handlers.NewData(l)

	sm := http.NewServeMux()
	sm.Handle("/", newHandler)

	http.ListenAndServe(":9090", sm)

	// config, err := config.LoadConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(config.Twitter.API_KEY)
	// 	fmt.Println(config.Twitter.SECRET)
	// }

	// tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret("key", "secret").RequestBearerToken()

}
