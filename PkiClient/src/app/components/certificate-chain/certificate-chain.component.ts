import { FocusMonitorDetectionMode } from '@angular/cdk/a11y';
import { NestedTreeControl } from '@angular/cdk/tree';
import { NgLocalization } from '@angular/common';
import { templateSourceUrl } from '@angular/compiler';
import { Component } from '@angular/core';
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { ChildrenOutletContexts } from '@angular/router';
import { CertificateView } from 'src/app/interfaces/certificate-view';
import { CertificateService } from 'src/app/services/CertificateService/certificate.service';


export interface FoodNode {
    name: string;
    children?: FoodNode[];
}



@Component({
    selector: 'app-certificate-chain',
    templateUrl: './certificate-chain.component.html',
    styleUrls: ['./certificate-chain.component.css']
})
export class CertificateChainComponent {
    c!:any;
    a!:CertificateView[];

    treeControl = new NestedTreeControl<FoodNode>(node => node.children);
    dataSource = new MatTreeNestedDataSource<FoodNode>();

    constructor(service :CertificateService) {
       
         service.getUsersChainCertificates("eamail").subscribe(
            (res) => {
             this.c= res;
            
             for (let i = 0; i < this.c.length; i++) {
               
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
            this.dataSource.data = items;
            for (let i = 0; i < res.length; i++) {
                let q = res[i] as CertificateView[][];
                for (let j = 0; j < res[i].length; j++) {
                   
                    let w = q[j] as CertificateView[]; 
                
                   
                    // for (let k = w.length-1 ; k >= 0; k--){
                    //     console.log(w[k]);
                    //     let children: FoodNode[];
                    //     if(w[k-1]!=null){
                    //     childr                        }

                    // } 
                    //     let node : FoodNode = {
                    //         name:w[k].serialNumber.toString(),
                    //         children =
                    //     }
                    // }

                    
                    
                }
            }
             }
        );
        }
    hasChild = (_: number, node: FoodNode) => !!node.children && node.children.length > 0;

    ngOnInit(): void {
    }
}
