package main

import (
	"fmt"
	"net/http"

	"github.com/paulgureghian/Go_Webservice/controllers"
)

func main() {

	fmt.Println("Web service starting")

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	fmt.Println("Web service started on localhost:3000")
}
