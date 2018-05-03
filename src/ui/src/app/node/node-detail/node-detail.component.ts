import { ChangeDetectorRef, Component } from '@angular/core';
import { INodeDetail, NodeService } from "../node.service";
import { MessageService } from "../../shared/message-service/message.service";
import "rxjs/add/operator/zip"
import "rxjs/add/operator/do"

@Component({
  selector: 'node-detail',
  templateUrl: './node-detail.component.html'
})
export class NodeDetailComponent {
  nodeDetailOpened: boolean;
  nodeDetail: INodeDetail;
  nodeGroups: string;

  constructor(private nodeService: NodeService,
              private changeDetectorRef: ChangeDetectorRef,
              private messageService: MessageService) {
    this.changeDetectorRef.detach();
  }

  openNodeDetailModal(nodeName: string): void {
    this.changeDetectorRef.detach();
    this.nodeDetailOpened = true;
    this.nodeGroups = "";
    let obs1 = this.nodeService.getNodeByName(nodeName)
      .do((nodeDetail: INodeDetail) => this.nodeDetail = nodeDetail);
    let obs2 = this.nodeService.getNodeGroupsOfOneNode(nodeName)
      .do((res: Array<string>) => res.forEach(value => this.nodeGroups = this.nodeGroups.concat(`${value};`)));
    obs1.zip(obs2)
      .subscribe(
        () => this.changeDetectorRef.reattach(),
        (err) => {
          this.nodeDetailOpened = false;
          this.messageService.dispatchError(err);
        });
  }

  toPercentage(num: number) {
    return Math.round(num * 100) / 100 + '%';
  }

  toGigaBytes(num: string, baseUnit?: string) {
    let denominator = 1024 * 1024 * 1024;
    if (baseUnit === 'KiB') {
      denominator = 1024 * 1024;
    }
    return Math.round(Number.parseInt(num) / denominator) + 'GB';
  }
}