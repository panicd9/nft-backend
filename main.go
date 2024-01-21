package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nft-backend/config"
	"nft-backend/database"
	"nft-backend/deployer"
	"nft-backend/indexer"
	"nft-backend/model"
	"nft-backend/utils"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

var contractAddrSepolia = "0x3d31De8Ecdd75f02dE09108c48af6CC219AEa3dC"
var contractAddrGanache = "0xE7436A481248471e88b473188da401325951f670"

func main() {
	godotenv.Load()
	config.SetDeployerPrivateKey(utils.GoDotEnvVariable("DEPLOYER_PRIVATE_KEY_HEX"))
	config.SetRPCEndpoint(utils.GoDotEnvVariable("RPC_API_KEY"))

	db = database.ConnectToDB()

	router := mux.NewRouter()

	// Enable CORS for all routes
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	router.HandleFunc("/events/{address}", GetEventsHandler).Methods("GET")
	router.HandleFunc("/blacklist-user", AddToBlacklistHandler).Methods("POST")
	// contractAddr := indexer.DeployNftContract()
	// fmt.Println(contractAddr)
	config.SetContractAddress("0x3d31De8Ecdd75f02dE09108c48af6CC219AEa3dC")
	go indexer.ListenForEvents(config.CONTRACT_ADDR)

	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router))
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["address"]

	mintEvents := []model.MintedEvent{}
	db.Where("minter = ?", address).Find(&mintEvents)

	registerEvents := []model.RegisteredEvent{}
	db.Where("address = ?", address).Find(&registerEvents)

	blacklistEvents := []model.BlacklistedEvent{}
	db.Where("blacklisted_address = ?", address).Find(&blacklistEvents)

	allEvents := map[string]interface{}{
		"mintEvents":      mintEvents,
		"registerEvents":  registerEvents,
		"blacklistEvents": blacklistEvents,
	}

	jsonResponse, err := json.Marshal(allEvents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func AddToBlacklistHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Address   string `json:"address"`
		Signature string `json:"signature"`
		Timestamp int64  `json:"timestamp"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("Address: %s, Signature: %s, Timestamp: %d\n", requestData.Address, requestData.Signature, requestData.Timestamp)

	// Verify timestamp is not older than 60 seconds
	timeDiff := time.Now().Unix() - requestData.Timestamp
	fmt.Println("Time diff: ", timeDiff)

	if timeDiff > 60 {
		http.Error(w, "Outdated timestamp", http.StatusBadRequest)
		return
	}

	// Verify the signature
	message := fmt.Sprintf("Authorize-%d-%s", requestData.Timestamp, requestData.Address)
	success, err := deployer.AddToBlacklist(message, requestData.Signature, requestData.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !success {
		http.Error(w, "Signature verification failed", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User blacklisted successfully"))
}
