import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'sortByDate'
})
export class SortByDatePipe implements PipeTransform {

  transform(value: any, args?: any): any {
    const sortedValues = value.sort((a:any, b:any) => new Date(b.createdOn).getTime() - new Date(a.createdOn).getTime());
    return sortedValues;
  }

}
