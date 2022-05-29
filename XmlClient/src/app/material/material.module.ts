import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WavesModule, ButtonsModule } from 'angular-bootstrap-md'




@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    WavesModule,
    ButtonsModule
  ],
  exports: [WavesModule,ButtonsModule]
})
export class MaterialModule { }
