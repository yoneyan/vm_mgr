# vm_mgr
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-controller/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg)
![https://github.com/yoneyan/vm_mgr/workflows/Go-node/badge.svg](https://github.com/yoneyan/vm_mgr/workflows/Go-client/badge.svg)  

kvm management tool :computer:

vm_mgrはvm+managerを合わせた意味になっています。   

## 状況
|機能|状況|
|---|---|
|controller|NG|
|node|NG|
|client|NG|

### 現時点（2020/2/26 3:10）
実装予定の内容が多すぎるため、完了したタスクのみを挙げています。  
**完了タスク**
* gRPCによるデータのやり取り client -> node controller -> nodeはまだダメ
* cliライブラリをurfave/cliからspf13/cobraへ移行
* VM停止、開始
* VM死活監視
* node側DBの完成

## 実行テスト
### Node

`go run . start`
### Client
**VM作成**  
`go run . vm create -n test -c 1 -m 1024 -p /home/yoneyan/test.qcow2 -s 1024 -N br100 -v 200 -C /home/yoneyan/Downloads/ubuntu-18.04.4-live-server-amd64.iso -M false`  
**VM削除(dbからも消し去る)**  
`go run . vm delete 1`  
**VM起動**  
`go run . vm start 1`  
**VM停止**  
`go run . vm stop 1`  
**VM取得(name)**  
`go run . vm get name test`  
**VM取得(id)**  
`go run . vm get id 1`  
**VM取得(all)**  
`go run . vm get all`  


## 構想
* Controller
* Node
* Client  

上記の3つの構造を基本とします。  
ControllerとNode間はgRPCにてデータのやり取りを行います。
変更として、LibvirtのAPIを使わず直接QEMUを叩く設計に変えました。

**実装予定はGithubのProjectに載せています。**

#### 対処法
**authentication unavailable: no polkit agent available to authenticate action 'org.libvirt.unix.manage''**
```
usermod --append --groups libvirt `username`
```

### gRPCで使用するポート
Port: 50100/tcp  