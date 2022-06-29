import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'profileSearch'
})
export class ProfileSearchPipe implements PipeTransform {

  transform(searchList : Array<any>, search : string): any {

    if (searchList && search)
      return searchList.filter(
        (d) =>
          d.firstName.toLowerCase()
        .indexOf(search.toLowerCase()) > -1 ||

          d.lastName.toLowerCase()
          .indexOf(search.toLowerCase()) > -1
      );


    return searchList;
  }


}
