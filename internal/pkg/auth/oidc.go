package auth

import "github.com/rogalni/cng-hello-backend/config"

var Secret string

func Setup() {
	js := config.App.JwtSecret
	if js == "" {
		// _, err := http.Get(config.App.JwtCertUrl)
		// if err != nil {
		// 	log.Warn().Msg("Could not fetch JWT Certificate")
		// }

		// parse secert
		// secret = r.Body.Read()
	} else {
		Secret = js
	}
}
