package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
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
 * 写入智能合约-合约币交易
 */

func main() {
	priKey := "5cc1a2676080fe6a3ae0b107967fdcae3f3c671d89d0241828e2a137effacd81"         // 发起方私钥 地址0x2074d05c2d8C52a892E5A1dF0685378b89Ccc420 的私钥
	contractAddress := common.HexToAddress("0x956B3669D8914BFcaf6815f67CbC3299C27c58b8") // 多签合约地址
	tokenAddress := common.HexToAddress("0x3C7E3Ffad7CB26fC9E51F49D277aCFE09Ae73eA2")    // 代币合约地址
	destAddress := common.HexToAddress("0x9Af40dce2Ebc76F42Ea74e2cAe460181eFb27167")     // 转账对象地址
	amount := big.NewInt(4000)                                                           // 代币交易数额

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/5329b08a37c048d3a3370ca8d53ed609")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(priKey) //
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

	instance, err := Contracts.NewContracts(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	methodID, err := hex.DecodeString("a9059cbb")
	paddedAddress := common.LeftPadBytes(destAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx, err := instance.SubmitTransaction(auth, tokenAddress, big.NewInt(0), data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s \n", tx.Hash().Hex()) // tx sent: 0x3bddc671b15533f96b16586aa6d39f7db581c86f937e7ad2e0a8c15567fd647d

}
