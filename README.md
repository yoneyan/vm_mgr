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
|controller|OK(一部NG)|
|node|OK(一部NG)|
|client|OK(一部NG)|
現時点では基本的な機能は動作できるようになっています。  

## 特徴
* ユーザ認証、グループ認証ffが可能
* tokenを用いた認証が可能
* gRPCを使用
* cliライブラリとしてspf13/cobraの使用

## 仕組み
![vm_mgr](https://user-images.githubusercontent.com/40447529/76900943-a6940100-68dd-11ea-9d1c-801bdecbb7f1.png)  

|名称|中身|
|---|---|
|Client|コマンドによる操作|
|Controller|ユーザやグループやノード管理など|
|ggate|RestAPIの提供|  
|imacon|Imageの提供|  
|Node|VMホスト|

## 使用Port(gRPC)
|機能|ポート|
|---|---|
|imacon|50300/tcp|
|Controller|50200/tcp|
|Node| 50100/tcp|

**実装予定はGithubのProjectに載せています。**

## 実行
コマンドはWikiに乗せています。(一部を除く)