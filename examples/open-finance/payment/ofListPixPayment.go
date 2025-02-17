package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/open_finance"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := open_finance.NewGerencianet(credentials)

	const inicio = "2022-03-01T03:01:35Z"
	const fim = "2022-12-05T22:01:35Z"	

	res, err := gn.OfListPixPayment(inicio, fim) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}