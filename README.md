
### 1.  使用 abigen 生成 multisig.go  (已生成，此步骤不需重复)
    
    ```
    cd contracts && abigen --bin=multisig.bin  --abi=multisig.abi --pkg=Contracts --out=multisig.go
    ```
    - bin=multisig.bin :指定bin文件来源
    - abi=multisig.abi :指定abi文件来源
    - pkg=Contracts    :指定输出文件的包名
    - out=multisig.go  :指定输出合约交互文件名称

### 2.  deploy.go  

部署多签合约

    - 输入
        1. N个地址
        2. 确认数M
        3. 其中一个地址的私钥
    - 输出
        1. 交易Hash
        2. 多签合约地址    

### 3.  transfer.go 

发起主币转账交易

    - 输入
        1. 发起方私钥
        2. 多签合约地址
        3. 交易对象地址
        4. 交易数额
    - 输出
        1. 交易Hash  
    
### 4.  transfer_token.go 

发起合约币转账交易

    - 输入
        1. 发起方私钥
        2. 多签合约地址
        3. 交易对象地址
        4. 代币合约地址
        5. 代币交易数额 
    - 输出
        1. 交易Hash  

### 5.  confirm.go 

确认转账交易，达到required自动转出

    - 输入
        1. 待确认的交易Hash
        2. 确认方私钥
        3. 多签合约地址
    - 输出
        1. 交易Hash  
