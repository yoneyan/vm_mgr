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

### 現時点（2020/2/22 4:51）
実装予定の内容が多すぎるため、完了したタスクのみを挙げています。  
**完了タスク**
* gRPCによるデータのやり取り client -> node controller -> nodeはまだダメ
* cliライブラリをurfave/cliからspf13/cobraへ移行

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