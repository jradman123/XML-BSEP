import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'jobsSearch'
})
export class JobsSearchPipe implements PipeTransform {

  transform(searchList : Array<any>, search : string): any {

    if (searchList && search)
      return searchList.filter(
        (d) =>
          d.Position.toLowerCase()
        .indexOf(search.toLowerCase()) > -1 ||

          d.JobDescription.toLowerCase()
          .indexOf(search.toLowerCase()) > -1
      );


    return searchList;
  }

}
