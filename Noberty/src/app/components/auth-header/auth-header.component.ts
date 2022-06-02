import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoggedUserDto } from 'src/app/interfaces/logged-user-dto';
import { CompanyService } from 'src/app/services/company-service/company.service';
import { UserServiceService } from 'src/app/services/UserService/user-service.service';

@Component({
  selector: 'app-auth-header',
  templateUrl: './auth-header.component.html',
  styleUrls: ['./auth-header.component.css']
})
export class AuthHeaderComponent implements OnInit {

  currentUser : LoggedUserDto;
  role! : string ;
  id! : number;
  constructor(private userService : UserServiceService, private companyService : CompanyService, private router : Router) {
    this.currentUser = {} as LoggedUserDto;
   }

  ngOnInit(): void {
   this.currentUser = this.userService.getUserValue();
   this.role = this.currentUser.role;
   console.log(this.role);

   // get companies of this user and navigate to /company/id
  }

  neznamdrugacije() {
    this.companyService.getAllUsersCompanies(this.currentUser.username).subscribe(res =>
      {
        this.id = res[0].companyId;

        console.log(res);
        this.router.navigate(['/company/' + this.id ]);

      });
  }

  logout() : void {
    this.userService.logout();
  }

}
