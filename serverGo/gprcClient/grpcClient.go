package gprcClient

import (
	"google.golang.org/grpc"
	"rudnWebApp/pb"
)

func NewGrpCClient() (pb.FileGeneratorClient, error) {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewFileGeneratorClient(conn)
	return client, nil
}
