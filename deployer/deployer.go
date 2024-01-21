package deployer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"nft-backend/config"
	"nft-backend/nft"

	ps "github.com/etaaa/Golang-Ethereum-Personal-Sign"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployNftContract() string {
	client, err := ethclient.Dial(config.RPC_HTTP_URL)
	if err != nil {
		fmt.Println(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), config.DEPLOYER_ADDR)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.DEPLOYER_SK, big.NewInt(config.SEPOLIA_CHAIN_ID))
	if err != nil {
		fmt.Println(err)
	}

	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	address, _, _, err := nft.DeployNft(auth, client)
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

	backendSignature, _ := ps.PersonalSign(hash.String(), config.DEPLOYER_SK)
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

	client, err := ethclient.Dial(config.RPC_HTTP_URL)
	if err != nil {
		return false, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), config.DEPLOYER_ADDR)
	if err != nil {
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.DEPLOYER_SK, big.NewInt(config.SEPOLIA_CHAIN_ID))
	if err != nil {
		return false, err
	}

	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))

	instance, err := nft.NewNft(common.Address(common.FromHex(config.CONTRACT_ADDR)), client)
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
