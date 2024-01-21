package indexer

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"nft-backend/database"
	"nft-backend/model"

	ps "github.com/etaaa/Golang-Ethereum-Personal-Sign"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RPC_HTTP_BASE_URL               = "https://sepolia.infura.io/v3"
	RPC_WS_BASE_URL                 = "wss://sepolia.infura.io/ws/v3"
	RPC_API_KEY                     = "xxx"
	GANACHE_HOST                    = "192.168.100.40"
	GANACHE_PORT                    = "7545"
	GANACHE_HTTP                    = fmt.Sprintf("http://%s:%s", GANACHE_HOST, GANACHE_PORT)
	GANACHE_WS                      = fmt.Sprintf("ws://%s:%s", GANACHE_HOST, GANACHE_PORT)
	RPC_HTTP_URL                    = fmt.Sprintf("%s/%s", RPC_HTTP_BASE_URL, RPC_API_KEY)
	RPC_WS_URL                      = fmt.Sprintf("%s/%s", RPC_WS_BASE_URL, RPC_API_KEY)
	deployerPrivateKeyHex           = ""
	deployerPrivateKey, _           = crypto.HexToECDSA(deployerPrivateKeyHex)
	deployerPublicKey               = deployerPrivateKey.Public()
	deployerPublicKeyECDSA, _       = deployerPublicKey.(*ecdsa.PublicKey)
	deployerAddress                 = crypto.PubkeyToAddress(*deployerPublicKeyECDSA)
	contractAddress                 = "0x3d31De8Ecdd75f02dE09108c48af6CC219AEa3dC"
	SEPOLIA_CHAIN_ID          int64 = 11155111
)

func ListenForEvents(_contractAddress string) {
	// Connect to Ethereum node using WebSocket
	client, err := ethclient.Dial(RPC_WS_URL)
	if err != nil {
		fmt.Println(err)
	}

	contract, err := NewNft(common.HexToAddress(_contractAddress), client)
	if err != nil {
		fmt.Println(err)
	}

	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel1 := make(chan *NftBlacklisted)

	// Start a goroutine which watches new events
	go func() {
		sub, err := contract.WatchBlacklisted(watchOpts, channel1, nil, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case blackListedEvent := <-channel1:
				handleBlackListedEvent(blackListedEvent)
			}
		}

	}()

	channel2 := make(chan *NftMinted)
	go func() {
		sub, err := contract.WatchMinted(watchOpts, channel2, nil, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case mintedEvent := <-channel2:
				handleMintedEvent(mintedEvent)
			}
		}
	}()

	channel3 := make(chan *NftRegistered)
	go func() {
		sub, err := contract.WatchRegistered(watchOpts, channel3, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case registeredEvent := <-channel3:
				handleRegisteredEvent(registeredEvent)
			}
		}
	}()

	fmt.Println("Waiting for events...")

}

func handleBlackListedEvent(event *NftBlacklisted) {
	fmt.Printf("Blacklisted event - AddedBy: %s, Blacklisted address: %s\n",
		event.AddedBy.Hex(), event.BlacklistedAddress.Hex())

	blacklistedEvent := &model.BlacklistedEvent{
		AddedBy:            event.AddedBy.Hex(),
		BlacklistedAddress: event.BlacklistedAddress.Hex(),
	}
	database.AddEvent(blacklistedEvent)
}

func handleRegisteredEvent(event *NftRegistered) {
	fmt.Printf("Registered event - Address: %s\n",
		event.User.Hex())

	registeredEvent := &model.RegisteredEvent{
		Address: event.User.Hex(),
	}
	database.AddEvent(registeredEvent)
}

func handleMintedEvent(event *NftMinted) {
	fmt.Printf("Minted event - User: %s, TokenId: %s\n",
		event.User.Hex(), event.TokenId)

	blacklistedEvent := &model.MintedEvent{
		Minter:  event.User.Hex(),
		TokenId: event.TokenId.Uint64(),
	}
	database.AddEvent(blacklistedEvent)
}

func DeployNftContract() string {
	client, err := ethclient.Dial(RPC_HTTP_URL)
	if err != nil {
		fmt.Println(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), deployerAddress)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(deployerPrivateKey, big.NewInt(SEPOLIA_CHAIN_ID))
	if err != nil {
		fmt.Println(err)
	}

	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	address, _, _, err := DeployNft(auth, client)
	if err != nil {
		fmt.Println(err)
	}

	return address.Hex()
}

// cant use crypto.Sign() beacuse it calculates ECDSA signature, not Ethereum signature which is received from Metamask
func VerifySignature(message string, signature string) bool {
	hash := crypto.Keccak256Hash([]byte(message))
	// fmt.Printf("Keccak hash: %v\n", hash.Hex())
	// backendSignature, err := crypto.Sign(hash.Bytes(), deployerPrivateKey)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("BackendSignature: %v\n", hexutil.Encode(backendSignature))

	backendSignature, _ := ps.PersonalSign(hash.String(), deployerPrivateKey)
	fmt.Printf("Backend signature: %v\n", backendSignature)

	// publicKeyBytes := crypto.FromECDSAPub(deployerPublicKeyECDSA)

	// signatureBytes := []byte(signature)
	// signatureNoRecoverID := signatureBytes[:len(signatureBytes)-1] // remove recovery ID
	// isValid := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), []byte(newSignature))

	// fmt.Printf("Signature: %v\n", signature)

	isValid := backendSignature == signature

	fmt.Printf("isValid: %v\n", isValid)

	return isValid
}

func AddToBlacklist(message string, signature string, addressToBlacklist string) (bool, error) {
	isValid := VerifySignature(message, signature)

	if !isValid {
		return false, errors.New("invalid signature")
	}

	client, err := ethclient.Dial(RPC_HTTP_URL)
	if err != nil {
		return false, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), deployerAddress)
	if err != nil {
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(deployerPrivateKey, big.NewInt(SEPOLIA_CHAIN_ID))
	if err != nil {
		return false, err
	}

	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	instance, err := NewNft(common.Address(common.FromHex(contractAddress)), client)
	tx, err := instance.AddToBlacklist(auth, common.Address(common.FromHex(addressToBlacklist)))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return false, err
	}

	return receipt.Status == 1, nil
}

func SetContractAddress(_contractAddress string) {
	contractAddress = _contractAddress
}

func SetDeployerPrivateKey(_deployerPrivateKey string) {
	deployerPrivateKeyHex = _deployerPrivateKey
	deployerPrivateKey, _ = crypto.HexToECDSA(deployerPrivateKeyHex)
	deployerPublicKey = deployerPrivateKey.Public()
	deployerPublicKeyECDSA, _ = deployerPublicKey.(*ecdsa.PublicKey)
	deployerAddress = crypto.PubkeyToAddress(*deployerPublicKeyECDSA)
}

func SetRPCEndpoint(_rpcAPIKey string) {
	RPC_API_KEY = _rpcAPIKey
	RPC_HTTP_URL = fmt.Sprintf("%s/%s", RPC_HTTP_BASE_URL, _rpcAPIKey)
	RPC_WS_URL = fmt.Sprintf("%s/%s", RPC_WS_BASE_URL, _rpcAPIKey)
}
