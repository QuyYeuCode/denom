module github.com/QuyYeuCode/denom

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.3
	github.com/cosmos/ibc-go/v2 v2.0.3
	github.com/ignite-hq/cli v0.20.0
	github.com/kr/pretty v0.3.1 // indirect
	github.com/spf13/cast v1.4.1
	github.com/stretchr/testify v1.8.4
	github.com/tendermint/spn v0.1.1-0.20220407154406-5cfd1bf28150
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.6
	golang.org/x/net v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/grpc v1.61.0 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)