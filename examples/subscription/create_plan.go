package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"name": "My plan",
		"interval": 2,
		"repeats": nil,
	}

	res, err := gn.CreatePlan(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}