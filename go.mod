module github.com/qiluge/filecoin-golang-sdk

go 1.14

require (
	github.com/filecoin-project/go-address v0.0.2-0.20200504173055-8b6f2fb2b3ef
	github.com/filecoin-project/lotus v0.4.2
	github.com/filecoin-project/specs-actors v0.6.2-0.20200724193152-534b25bdca30
	github.com/whyrusleeping/cbor-gen v0.0.0-20200723182808-cb5de1c427f5 // indirect
)

replace github.com/filecoin-project/lotus v0.4.2 => ../../filecoin-project/lotus

replace github.com/supranational/blst => github.com/supranational/blst v0.1.2-alpha.1

replace github.com/golangci/golangci-lint => github.com/golangci/golangci-lint v1.18.0

replace github.com/filecoin-project/filecoin-ffi => ../../filecoin-project/lotus/extern/filecoin-ffi
