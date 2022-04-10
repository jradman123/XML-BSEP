import { NestedTreeControl } from '@angular/cdk/tree';
import { Component, Input, OnInit } from '@angular/core';
import { MatTreeNestedDataSource } from '@angular/material/tree';


export interface FoodNode {
    name: string;
    children?: FoodNode[];
}

const items: FoodNode[] = [
    {
        name: 'Vegetables',
        children: [
            {
                name: 'Green',
                children:
                    [
                        { name: 'Broccoli' },
                        { name: 'Brussels sprouts' }],
            }
        ],
    },
];
@Component({
    selector: 'app-certificate-chain',
    templateUrl: './certificate-chain.component.html',
    styleUrls: ['./certificate-chain.component.css']
})
export class CertificateChainComponent {


    treeControl = new NestedTreeControl<FoodNode>(node => node.children);
    dataSource = new MatTreeNestedDataSource<FoodNode>();

    constructor() {
        this.dataSource.data = items;
    }

    hasChild = (_: number, node: FoodNode) => !!node.children && node.children.length > 0;

    ngOnInit(): void {
    }
}
