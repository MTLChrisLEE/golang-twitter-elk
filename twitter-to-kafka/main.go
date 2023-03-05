package main

import (
	"io/ioutil"
	"log"
	"net/http"
	// twitterstream "github.com/fallenstedt/twitter-stream"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("HELLO")

		d, _ := ioutil.ReadAll(r.Body)

		log.Printf("Data %s\n", d)

	})

	http.ListenAndServe(":9090", nil)

	// config, err := config.LoadConfig()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(config.Twitter.API_KEY)
	// 	fmt.Println(config.Twitter.SECRET)
	// }

	// tok, err := twitterstream.NewTokenGenerator().SetApiKeyAndSecret("key", "secret").RequestBearerToken()

}
