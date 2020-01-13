# vm_mgr
kvm management tool :computer:

This is KVM Management Tool.
I am aiming to manage KVM reasonably.

### 対処法
**authentication unavailable: no polkit agent available to authenticate action 'org.libvirt.unix.manage''**
```
usermod --append --groups libvirt `username`
```
