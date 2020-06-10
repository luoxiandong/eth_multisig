
1.  使用 abigen 生成 multisig.go  (已生成，此步骤不需重复)
    
    ```
    cd contracts && abigen --bin=multisig.bin  --abi=multisig.abi --pkg=Contracts --out=multisig.go
    ```
    - bin=multisig.bin :指定bin文件来源
    - abi=multisig.abi :指定abi文件来源
    - pkg=Contracts    :指定输出文件的包名
    - out=multisig.go  :指定输出合约交互文件名称

2.  deploy.go  部署多签合约

3.  transfer.go 发起主币转账交易

4.  transfer_token.go 发起合约币转账交易

5.  confirm.go (其他人)确认转账交易，达到required自动转出

