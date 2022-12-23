package saver

import (
	"os"

	"github.com/pkg/errors"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
	saver "gitlab.com/rarimo/savers/saver-grpc-lib/grpc"
	"google.golang.org/grpc"
)

const UrlSuffix = "_ADDR_ENV"

// Storage for saver clients connections
var connections = make(map[string]*grpc.ClientConn)

func Set(networks []*types.NetworkParams) {
	for _, con := range connections {
		con.Close()
	}

	connections = make(map[string]*grpc.ClientConn, len(networks))

	for _, network := range networks {
		addr := os.Getenv(network.Name + UrlSuffix)
		if addr == "" {
			panic(errors.New("address for " + network.Name + " saver not found."))
		}

		con, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		connections[network.Name] = con
	}
}

func GetClient(network string) (saver.SaverClient, error) {
	if con, ok := connections[network]; ok {
		return saver.NewSaverClient(con), nil
	}

	con, err := getClient(network)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get saver connection")
	}

	connections[network] = con
	return saver.NewSaverClient(con), nil
}

func getClient(network string) (*grpc.ClientConn, error) {
	addr := os.Getenv(network + UrlSuffix)
	if addr == "" {
		panic(errors.New("address for " + network + " saver not found."))
	}

	return grpc.Dial(addr, grpc.WithInsecure())
}
