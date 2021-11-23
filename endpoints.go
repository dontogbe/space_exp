package space_exp

import (
	"context"
	"encoding/json"
	"errors"
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
		l, err := svc.GetLocation(req.X, req.Y, req.Z, req.Vel)
		if err != nil {
			return locationResponse{Location: l}, err
		}
		return locationResponse{Location: l}, nil
	}
}

type locationRequest struct {
	X   float64 `json:"x,omitempty" `
	Y   float64 `json:"y,omitempty"`
	Z   float64 `json:"z,omitempty"`
	Vel float64 `json:"vel"`
}

type locationResponse struct {
	Location float64 `json:"loc"`
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
