package main

import (
	"context"
	"crypto/ecdsa"
	Contracts "eth_multisig/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**
 * 部署智能合约
 */
func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/5329b08a37c048d3a3370ca8d53ed609")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("5cc1a2676080fe6a3ae0b107967fdcae3f3c671d89d0241828e2a137effacd81")// 0x2074d05c2d8C52a892E5A1dF0685378b89Ccc420 的私钥
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	owners := []common.Address{
		common.HexToAddress("0x2074d05c2d8C52a892E5A1dF0685378b89Ccc420"),
		common.HexToAddress("0xaEAc2c548Eb63F8415308B3c153A58bE6d25278B"),
		common.HexToAddress("0x44A791a7C6F6F5d249539C7bBe5D0e378a49CfA3"),
	}
	required := big.NewInt(2)
	address, tx, instance, err := Contracts.DeployContracts(auth, client, owners,required)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x956B3669D8914BFcaf6815f67CbC3299C27c58b8
	fmt.Println(tx.Hash().Hex()) // 0x8c337b45413636bbc4453f55387c91bcd40f6b811f94da81ccce3dcaa84d9e34

	_ = instance
}
