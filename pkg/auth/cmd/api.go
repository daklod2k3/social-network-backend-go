package cmd

import "auth/internal"

func StartAPI() {
	server := internal.NewServer()
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
