import { Component, OnInit } from '@angular/core';
import { CompanyResponseDto } from 'src/app/interfaces/company-response-dto';
import { LoggedUserDto } from 'src/app/interfaces/logged-user-dto';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-my-companies-list',
  templateUrl: './my-companies-list.component.html',
  styleUrls: ['./my-companies-list.component.css']
})
export class MyCompaniesListComponent implements OnInit {
  items!: CompanyResponseDto[];
  currentUser: LoggedUserDto;
  constructor(
    private companyService: CompanyService,
    private userService: UserServiceService
  ) {
    this.currentUser = {} as LoggedUserDto;
  }

  ngOnInit(): void {
    this.currentUser = this.userService.getUserValue();
    this.companyService.getAllUsersCompanies(this.currentUser.username).subscribe(res => {
      this.items = res;
      console.log(res);

    });
  }


}
