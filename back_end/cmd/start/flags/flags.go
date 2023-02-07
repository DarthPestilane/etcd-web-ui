package flags

import (
	"github.com/spf13/pflag"
)

var (
	Endpoint  string // etcd endpoint to connect to
	Username  string // etcd username
	Password  string // etcd password
	KeyPrefix string // etcd key prefix
)

func Init(flagSet *pflag.FlagSet) {
	flagSet.SortFlags = false
	flagSet.StringVar(&Endpoint, "endpoint", "127.0.0.1:2379", "the endpoint to connect to")
	flagSet.StringVar(&Username, "username", "", "etcd auth username")
	flagSet.StringVar(&Password, "password", "", "etcd auth password")
	flagSet.StringVar(&KeyPrefix, "key-prefix", "/", "etcd key prefix")
}
