package main

import (
	"context"
	"crypto/ecdsa"
	Contracts "eth_multisig/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

)

/**
 * 写入智能合约
 */

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/5329b08a37c048d3a3370ca8d53ed609")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("5cc1a2676080fe6a3ae0b107967fdcae3f3c671d89d0241828e2a137effacd81") //地址 0x2074d05c2d8C52a892E5A1dF0685378b89Ccc420 的私钥
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
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x956B3669D8914BFcaf6815f67CbC3299C27c58b8")
	instance, err := Contracts.NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}

	destAddress := common.HexToAddress("0x9Af40dce2Ebc76F42Ea74e2cAe460181eFb27167")
	val := big.NewInt(1000000000000000000) // in wei
	data := []byte("liyaojian")

	tx, err := instance.SubmitTransaction(auth,destAddress,val,data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s \n", tx.Hash().Hex()) // tx sent: 0xf57a7f6cd6375cdf15cbe527b2d3be32f4f256cfbf981f30f6c199988df1ae51

}
