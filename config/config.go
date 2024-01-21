package config

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	RPC_HTTP_BASE_URL = "https://sepolia.infura.io/v3"
	RPC_WS_BASE_URL   = "wss://sepolia.infura.io/ws/v3"
	RPC_API_KEY       string
	GANACHE_HOST      = "192.168.100.40"
	GANACHE_PORT      = "7545"
	GANACHE_HTTP      string
	GANACHE_WS        string
	RPC_HTTP_URL      string
	RPC_WS_URL        string
	DEPLOYER_SK       *ecdsa.PrivateKey
	DEPLOYER_ADDR     common.Address
	CONTRACT_ADDR           = "0x3d31De8Ecdd75f02dE09108c48af6CC219AEa3dC"
	SEPOLIA_CHAIN_ID  int64 = 11155111
)

func SetContractAddress(_contractAddress string) {
	CONTRACT_ADDR = _contractAddress
}

func SetDeployerPrivateKey(_deployerPrivateKey string) {
	var deployerSecretKeyHex = _deployerPrivateKey
	DEPLOYER_SK, _ = crypto.HexToECDSA(deployerSecretKeyHex)
	var deployerPublicKey = DEPLOYER_SK.Public()
	var deployerPublicKeyECDSA, _ = deployerPublicKey.(*ecdsa.PublicKey)
	DEPLOYER_ADDR = crypto.PubkeyToAddress(*deployerPublicKeyECDSA)
}

func SetRPCEndpoint(_rpcAPIKey string) {
	RPC_API_KEY = _rpcAPIKey
	RPC_HTTP_URL = fmt.Sprintf("%s/%s", RPC_HTTP_BASE_URL, _rpcAPIKey)
	RPC_WS_URL = fmt.Sprintf("%s/%s", RPC_WS_BASE_URL, _rpcAPIKey)
}
