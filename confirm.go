package main

import (
	"context"
	"crypto/ecdsa"
	Contracts "eth_multisig/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/**
 * 写入智能合约-确认交易
 */

func main() {
	txHash := common.HexToHash("0x6672c1b8cdcbcfd6efdca138e6e4c9e3af3f4ee2e77a4df744a7cb4b615c606e") // transfer交易返回的HashID
	priKey := "B1DA1D9167CDEB85B9FA486A197C67BA78431E9B6A90F2D3CD4A53B46831DD71"                     // 确认方私钥  0xaEAc2c548Eb63F8415308B3c153A58bE6d25278B 的私钥
	contractAddress := common.HexToAddress("0x956B3669D8914BFcaf6815f67CbC3299C27c58b8")             // 多签约地址

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/5329b08a37c048d3a3370ca8d53ed609")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	instance, err := Contracts.NewContracts(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	transactionId := receipt.Logs[0].Topics[1].Big()
	fmt.Println("transactionId: ", transactionId)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	tx, err := instance.ConfirmTransaction(auth, transactionId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xcd0f4ded8ce182e9971430c25937b0e9506ba01c3fa368dec51b503ef057e712
}
