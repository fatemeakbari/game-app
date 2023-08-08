package presenceadapter

import (
	"context"
	"gameapp/contract/goproto/presence"
	presenceentity "gameapp/entity/presence"
	"gameapp/pkg"
	"google.golang.org/grpc"
)

type Client struct {
	address string
}

func New(address string) Client {
	return Client{
		address: address,
	}
}

func (c Client) GetPresence(ctx context.Context, request presenceentity.GetPresenceRequest) (*presenceentity.GetPresenceResponse, error) {

	conn, err := grpc.Dial(c.address, grpc.WithInsecure())

	if err != nil {
		return &presenceentity.GetPresenceResponse{}, err
	}

	client := presence.NewPresenceClient(conn)

	req := pkg.MapToProGetPresenceRequest(request)

	res, err := client.GetPresence(ctx, &req)

	if err != nil {
		return &presenceentity.GetPresenceResponse{}, err
	}

	return pkg.MapToGetPresenceResponse(res), nil

}
