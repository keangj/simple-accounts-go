package main

import "simple-accounts/cmd"

// @title           Simple Accounts API
// @version         1.0
// @description     This is a sample server for Simple Accounts API.

// @contact.name   Jay
// @contact.url    null
// @contact.email  keangjay@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	cmd.Run()
}
