# vm_mgr
Master->
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-controller/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-client/badge.svg)  

kvm management tool :computer:

VMを管理するという意味を込めてvm_mgrとしています。   

**実装予定はGithubのProjectに載せています。**
##必須条件
**go v1.11**以上を推奨  

go 1.10はgGate以外であれば可能だが、ggateはgo 1.11以上である必要あり  
使用しているパッケージ(gin)が1.11以上でないと動かないため  

## 状況
|機能|状況|
|---|---|
|controller|OK(おそらく)|
|node|OK(おそらく)|
|client|OK(おそらく)|
|gGate|OK(VM作成やユーザ・グループ追加などを除く)|
|imacon|OK(Joinのみ。SFTP経由による追加は未定)|
現時点では基本的な機能は動作できるようになっています。  

## 特徴
* ユーザ認証、グループ認証が可能
* tokenを用いた認証が可能
* gRPCを使用
* cliライブラリとしてspf13/cobraの使用

## 仕組み
![vm_mgr](https://user-images.githubusercontent.com/40447529/77772317-43e3f780-708b-11ea-8e12-0dea6d622174.png)

|名称|内容|
|---|---|
|Client|コマンドによる操作|
|Controller|ユーザやグループやノード管理など|
|ggate|RestAPIの提供|  
|imacon|Imageの提供|  
|Node|VMホスト|
|Nginx(3rd party)|Webサーバ|


## スケーラビリティ
### 対応状況
* imacon  
* node  
* gGate  
上記の3つのシステムのみ可能  
コントローラは現時点ではできないが、対応予定あり  

## 使用Port
### gRPC
|機能|ポート|
|---|---|
|imacon|50300/tcp|
|Controller|50200/tcp|
|Node| 50100/tcp|
### その他
|機能|ポート|用途|
|---|---|---|
|nginx|80/tcp,443/tcp|Webサーバ|
|gGate|8080/tcp|RestAPI|
|Controller|8081/tcp|WebSocket(VNC)|

## 実行
コマンドはWikiに乗せています。(一部を除く)