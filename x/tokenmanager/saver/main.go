package saver

import (
	"errors"
	"os"

	"gitlab.com/rarify-protocol/rarimo-core/x/tokenmanager/types"
	saver "gitlab.com/rarify-protocol/saver-grpc-lib/grpc"
	"google.golang.org/grpc"
)

const UrlSuffix = ".ADDR.ENV"

// Storage for saver clients connections
var connections = make(map[string]*grpc.ClientConn)

func Set(networks map[string]*types.ChainParams) {
	for _, con := range connections {
		con.Close()
	}

	connections = make(map[string]*grpc.ClientConn, len(networks))

	for network, _ := range networks {
		addr := os.Getenv(network + UrlSuffix)
		if addr == "" {
			panic(errors.New("address for " + network + " saver not found."))
		}

		con, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		connections[network] = con
	}
}

func GetClient(network string) saver.SaverClient {
	if con, ok := connections[network]; ok {
		return saver.NewSaverClient(con)
	}

	return nil
}
