package main

import (
	kitHTTP "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"

	"github.com/dontogbe/space_exp"
	"github.com/dontogbe/space_exp/svc"
)

func main() {
	svc := svc.NewSvc()
	serviceEndpoints := space_exp.NewServiceEndpoints(svc)
	mappingHandler := kitHTTP.NewServer(
		serviceEndpoints.GetLocationEndpoint,
		space_exp.DecodeLocationRequest,
		space_exp.EncodeLocationResponse,
	)
	http.Handle("/map", mappingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
