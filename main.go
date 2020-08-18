package main

import (
		"net/http"

		"github.com/paulgureghian/Go_Webservice/controllers"
)

func main() {

		controllers.RegisterControllers()
		http.ListenAndServe(":3000", nil)

}
