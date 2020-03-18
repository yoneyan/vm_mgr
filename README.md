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
ClientはコマンドをControllerやNodeに命令を送る部分  
Controllerはユーザやグループやノード管理などを担う部分  
ggateはRestAPIを提供  
imaconはImageを提供  
Nodeはqemuを使って実際にVMを動かしている部分    

## 使用Port
|機能|ポート|
|---|---|
|Controller|50200/tcp|
|Node| 50100/tcp|

**実装予定はGithubのProjectに載せています。**

## 実行
コマンドはWikiに乗せています。(一部を除く)