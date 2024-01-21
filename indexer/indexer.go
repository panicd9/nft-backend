package indexer

import (
	"context"
	"fmt"
	"nft-backend/config"
	"nft-backend/database"
	"nft-backend/model"
	"nft-backend/nft"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenForEvents(_contractAddress string) {
	// Connect to Ethereum node using WebSocket
	client, err := ethclient.Dial(config.RPC_WS_URL)
	if err != nil {
		fmt.Println(err)
	}

	contract, err := nft.NewNft(common.HexToAddress(_contractAddress), client)
	if err != nil {
		fmt.Println(err)
	}

	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	blacklistedChannel := make(chan *nft.NftBlacklisted)

	// Start a goroutine which watches new events
	go func() {
		sub, err := contract.WatchBlacklisted(watchOpts, blacklistedChannel, nil, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case blackListedEvent := <-blacklistedChannel:
				handleBlackListedEvent(blackListedEvent)
			}
		}

	}()

	mintedChannel := make(chan *nft.NftMinted)
	go func() {
		sub, err := contract.WatchMinted(watchOpts, mintedChannel, nil, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case mintedEvent := <-mintedChannel:
				handleMintedEvent(mintedEvent)
			}
		}
	}()

	registeredChannel := make(chan *nft.NftRegistered)
	go func() {
		sub, err := contract.WatchRegistered(watchOpts, registeredChannel, nil)
		defer sub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}

		for {
			select {
			case registeredEvent := <-registeredChannel:
				handleRegisteredEvent(registeredEvent)
			}
		}
	}()

	fmt.Println("Waiting for events...")

}

func handleBlackListedEvent(event *nft.NftBlacklisted) {
	fmt.Printf("Blacklisted event - AddedBy: %s, Blacklisted address: %s\n",
		event.AddedBy.Hex(), event.BlacklistedAddress.Hex())

	blacklistedEvent := &model.BlacklistedEvent{
		AddedBy:            event.AddedBy.Hex(),
		BlacklistedAddress: event.BlacklistedAddress.Hex(),
	}
	database.AddEvent(blacklistedEvent)
}

func handleRegisteredEvent(event *nft.NftRegistered) {
	fmt.Printf("Registered event - Address: %s\n",
		event.User.Hex())

	registeredEvent := &model.RegisteredEvent{
		Address: event.User.Hex(),
	}
	database.AddEvent(registeredEvent)
}

func handleMintedEvent(event *nft.NftMinted) {
	fmt.Printf("Minted event - User: %s, TokenId: %s\n",
		event.User.Hex(), event.TokenId)

	blacklistedEvent := &model.MintedEvent{
		Minter:  event.User.Hex(),
		TokenId: event.TokenId.Uint64(),
	}
	database.AddEvent(blacklistedEvent)
}
