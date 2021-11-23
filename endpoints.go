package space_exp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"log"
	"net/http"

	"github.com/dontogbe/space_exp/svc"
)

type serviceEndpoints struct {
	GetLocationEndpoint endpoint.Endpoint
}

func NewServiceEndpoints(svc svc.Svc) *serviceEndpoints {
	return &serviceEndpoints{
		GetLocationEndpoint: makeGetLocationEndpoint(svc),
	}
}

func makeGetLocationEndpoint(svc svc.Svc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(locationRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		log.Print(req)
		l, err := svc.GetLocation(req.X, req.Y, req.Z, req.Vel)
		if err != nil {
			return locationResponse{Location: fmt.Sprintf("%.2f", l)}, err
		}
		return locationResponse{Location: fmt.Sprintf("%.2f", l)}, nil
	}
}

type locationRequest struct {
	X   float64 `json:"x" `
	Y   float64 `json:"y"`
	Z   float64 `json:"z"`
	Vel float64 `json:"vel"`
}

type locationResponse struct {
	Location string `json:"loc"`
}

func DecodeLocationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request locationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err)

		return nil, err
	}
	return request, nil
}
func EncodeLocationResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
