module github.com/qiluge/filecoin-golang-sdk

go 1.14

require (
	github.com/filecoin-project/go-address v0.0.2-0.20200504173055-8b6f2fb2b3ef
	github.com/filecoin-project/lotus v0.4.2
	github.com/filecoin-project/specs-actors v0.8.1-0.20200723200253-a3c01bc62f99
	github.com/whyrusleeping/cbor-gen v0.0.0-20200723182808-cb5de1c427f5 // indirect
)

replace github.com/filecoin-project/lotus v0.4.2 => ../../filecoin-project/lotus
