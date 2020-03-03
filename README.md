# vm_mgr
Master->
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-controller/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-client/badge.svg)  

kvm management tool :computer:

VMを管理するという意味を込めてvm_mgrとしています。   

## 状況
|機能|状況|
|---|---|
|controller|NG|
|node|OK(一部NG)|
|client|OK(一部NG)|

## 特徴
* ユーザ認証が可能（未実装）
* gRPCを使用
* cliライブラリとしてspf13/cobraの使用

## 使用Port
|機能|ポート|
|---|---|
|Controller|50200/tcp|
|Node| 50100/tcp|

## 構想
* Controller
* Node
* Client  

上記の3つの構造を基本とします。  
ControllerとNode間はgRPCにてデータのやり取りを行います。

**実装予定はGithubのProjectに載せています。**


## 実行テスト
### Node

**起動**
`go build .`  
`sudo ./node start`  
*qemuコマンドを使用するため、root権限必須*

### Client -> Node
現状ではホストIPなしでもアクセス可能であるが、最終的にはホストIP必須になる予定  
**VM作成**  
例として teという名前のcore:1,memory:1024のVMを作成  
`go run . vm create -n te -c 1 -m 1024 -P /home/yoneyan -s 10240 -N br0 -v 200 -C /home/yoneyan/Downloads/ubuntu-18.04.4-live-server-amd64.iso -a false -H 127.0.0.1`  
**VM削除(dbからも消し去る)**  
`go run . vm delete 1 -H 127.0.0.1`  
**VM起動**  
`go run . vm start 1 -H 127.0.0.1`  
**VM停止**  
`go run . vm stop 1 -H 127.0.0.1`  
**VM取得(name)**  
`go run . vm get name test -H 127.0.0.1`  
**VM取得(id)**  
`go run . vm get id 1 -H 127.0.0.1`  
**VM取得(all)**  
`go run . vm get all -H 127.0.0.1`  
**Node停止**
`go run . node stop -H 127.0.0.1:50100`

### Client -> Controller
**ユーザ作成**
aというユーザにaというパスワード  
`go run . user add a a -H 127.0.0.1:50200 -u test -p test`
**ユーザ削除**
`go run . user remove a -H 127.0.0.1:50200 -u test -p test`