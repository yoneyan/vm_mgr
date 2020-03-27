import {Component, OnInit} from '@angular/core';
import {ImageService} from "../../service/image.service";
import {Location} from "@angular/common";
import {GroupService} from "../../service/group.service";
import {NodeService} from "../../service/node.service";
import {SpecService} from "../../service/spec.service";
import {VmService} from "../../service/vm.service";
import {FormBuilder} from "@angular/forms";


@Component({
  selector: 'app-create-vm',
  templateUrl: './create-vm.component.html',
  styleUrls: ['./create-vm.component.css']
})
export class CreateVmComponent implements OnInit {

  constructor(private ImageService: ImageService,
              private location: Location,
              private GroupService: GroupService,
              private NodeService: NodeService,
              private SpecService: SpecService,
              public VMService: VmService,
              private formBuilder: FormBuilder,
  ) {
  }

  public groups: GroupDataStruct[];
  public nodes: NodeDataStruct[];
  public images: ImageDataStruct[];
  public cpus: CPUDataStruct[];
  public memorys: MemoryDataStruct[];
  public storages: StorageDataStruct[];
  public storagesizes: StorageSizeDataStruct[];
  public model: SpecDataStruct
    = new class implements SpecDataStruct {
    cpu: number;
    group: string;
    image: string;
    memory: number;
    name: string;
    nodeid: number;
    storage: number;
    storagetype: number;
  }
  message:string;

  ngOnInit(): void {
    this.getImage()
    this.getNode()
    this.getGroup()
    this.getCPU()
    this.getMemory()
    this.getStorage()
    this.getStorageSize()
  }

  getImage(): void {
    this.ImageService.getImage()
      .then(d => {
        // let data:
        this.images = d
        console.log(d)
      })
  }

  getNode(): void {
    this.NodeService.getNode()
      .then(d => {
        // let data:
        this.nodes = d
        console.log(d)
      })
  }

  getGroup(): void {
    this.GroupService.getGroup()
      .then(d => {
        // let data:
        this.groups = d
        console.log(d)
      })
  }

  getCPU(): void {
    this.SpecService.getCPU()
      .then(d => {
        // let data:
        this.cpus = d
        console.log(d)
      })
  }

  getMemory(): void {
    this.SpecService.getMemory()
      .then(d => {
        // let data:
        this.memorys = d
        console.log(d)
      })
  }

  getStorage(): void {
    this.SpecService.getStorage()
      .then(d => {
        // let data:
        this.storages = d
        console.log(d)
      })
  }

  getStorageSize(): void {
    this.SpecService.getStorageSize()
      .then(d => {
        // let data:
        this.storagesizes = d
        console.log(d)
      })
  }

  goBack(): void {
    this.location.back();
  }

  onSubmit(data) {
    // Process checkout data here

    // this.VMService.createVM()
    //   .then(r => [
    //     console.log(r)
    //   ])

    var result = this.model.image.split(':')
    var cpu = this.model.cpu

    console.dir(data);

    this.VMService.createVM(data.nodeid, data.name, data.group, cpu, data.memory, data.storage, data.storagetype, result[0], result[1])
      .then(d => {
        // let data:
        this.storagesizes = d
        console.log(d)
      })
    this.message = 'Create Process...';
  }
}

interface ImageDataStruct {
  name: string
  tag: string
}

interface GroupDataStruct {
  name: string
}

interface NodeDataStruct {
  id: string
  hostname: string
}

interface CPUDataStruct {
  spec: number
}

interface MemoryDataStruct {
  size: number
}

interface StorageDataStruct {
  type: string
  name: string
}

interface StorageSizeDataStruct {
  size: number
}

interface SpecDataStruct {
  nodeid: number,
  name: string,
  group: string,
  cpu: number,
  memory: number,
  storagetype: number,
  storage: number,
  image: string
}
