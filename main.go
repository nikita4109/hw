package main

import (
	"log"
	"math/big"
	"net/http"

	"hw/pkg/userwallet"
	"hw/pkg/walletfactory"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	ethNodeURL             = "https://sepolia.drpc.org"
	factoryContractAddress = "0xb3b936149f569b4ed9fb2c6035113005a9939b12"
)

var client *ethclient.Client
var ks *keystore.KeyStore
var auth *bind.TransactOpts
var walletFactory *walletfactory.Walletfactory

func init() {
	var err error
	client, err = ethclient.Dial(ethNodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("c97de53acd5dab88b30321690f352a8908f73bff2b7ae375e3c5bcebec3f55b3")
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	auth = bind.NewKeyedTransactor(privateKey)
	walletFactory, err = walletfactory.NewWalletfactory(common.HexToAddress(factoryContractAddress), client)
	if err != nil {
		log.Fatalf("Failed to instantiate WalletFactory: %v", err)
	}
}

func registerUser(c *gin.Context) {
	tx, err := walletFactory.CreateWallet(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transaction": tx.Hash().Hex()})
}

func depositFunds(c *gin.Context) {
	var req struct {
		WalletAddress string  `json:"wallet_address"`
		Amount        float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount := big.NewInt(int64(req.Amount * 1e18))
	wallet, err := userwallet.NewUserwallet(common.HexToAddress(req.WalletAddress), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := wallet.Deposit(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  amount,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": tx.Hash().Hex()})
}

func buyAsset(c *gin.Context) {
	var req struct {
		WalletAddress string `json:"wallet_address"`
		Asset         string `json:"asset"`
		Amount        uint64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount := big.NewInt(int64(req.Amount * 1e18))
	wallet, err := userwallet.NewUserwallet(common.HexToAddress(req.WalletAddress), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := wallet.BuyAsset(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, req.Asset, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": tx.Hash().Hex()})
}

func sellAsset(c *gin.Context) {
	var req struct {
		WalletAddress string `json:"wallet_address"`
		Asset         string `json:"asset"`
		Amount        uint64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount := big.NewInt(int64(req.Amount * 1e18))
	wallet, err := userwallet.NewUserwallet(common.HexToAddress(req.WalletAddress), client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx, err := wallet.SellAsset(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, req.Asset, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": tx.Hash().Hex()})
}

func getRates(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		Get("https://www.cbr-xml-daily.ru/daily_json.js")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp.String()})
}

func main() {
	ks = keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	router := gin.Default()

	router.POST("/register", registerUser)
	router.POST("/deposit", depositFunds)
	router.POST("/buy", buyAsset)
	router.POST("/sell", sellAsset)
	router.GET("/rates", getRates)

	router.Run(":8080")
}
