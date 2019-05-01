package main

import (
		"fmt"
		"net/http"
		"os"

		"github.com/luthfiswees/sonar/handler"
		"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	http.HandleFunc("/data", handler.Handler)

	fmt.Println("Now serving on " + os.Getenv("SONAR_PORT"))
	http.ListenAndServe(":" + os.Getenv("SONAR_PORT"), nil)
}