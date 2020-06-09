## 使用 abigen 生成 multisig.go

```
cd contracts && abigen --bin=multisig.bin  --abi=multisig.abi --pkg=Contracts --out=multisig.go
```
- bin multisig.bin :指定bin文件来源
- abi multisig.abi :指定abi文件来源
- pkg Contracts    :指定输出文件的包名
- out multisig.go  :指定输出合约交互文件名称

## deploy.go 部署合约

## transfer.go 发起转账交易

## confirm.go 确认转账交易

