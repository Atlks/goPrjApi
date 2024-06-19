package lib

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strings"
	"time"
)

const (
	rpcURL               = "https://arb1.arbitrum.io/rpc"               // Arbitrum RPC URL
	privateKeyHex        = "your_private_key_hex"                       // 替换为你的私钥
	usdtContractAddress  = "0xdAC17F958D2ee523a2206206994597C13D831ec7" // USDT 合约地址
	uniswapRouterAddress = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" // Uniswap V2 Router 地址
	amountInUSD          = 3000                                         // 要购买的金额（美元）
	usdtDecimals         = 6
)

var routerABI = `[... Uniswap V2 Router ABI ...]`

func BuyEth() {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(42161)) // Arbitrum Chain ID
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	usdtAmount := decimal.NewFromFloat(amountInUSD).Mul(decimal.New(1, usdtDecimals)).BigInt()

	contractAddress := common.HexToAddress(uniswapRouterAddress)
	parsed, err := abi.JSON(strings.NewReader(routerABI))
	if err != nil {
		log.Fatalf("Failed to parse router ABI: %v", err)
	}

	input, err := parsed.Pack("swapExactTokensForETH", usdtAmount, big.NewInt(0), []common.Address{
		common.HexToAddress(usdtContractAddress),
		common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"), // WETH
	}, fromAddress, big.NewInt(time.Now().Add(time.Minute*15).Unix()))
	if err != nil {
		log.Fatalf("Failed to pack input: %v", err)
	}

	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), auth.GasLimit, auth.GasPrice, input)

	signedTx, err := auth.Signer(auth.From, tx)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
