package gprcClient

import (
	"google.golang.org/grpc"
	"rudnWebApp/pb"
	"rudnWebApp/util"
)

func NewGrpCClient(config util.Config) (pb.FileGeneratorClient, error) {
	conn, err := grpc.Dial(config.GrpcServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewFileGeneratorClient(conn)
	return client, nil
}
