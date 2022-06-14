import { Injectable } from "@angular/core";
import { CanActivate, Router } from "@angular/router";
import { UserServiceService } from "../services/UserService/user-service.service";

@Injectable({ providedIn: 'root' })
export class OwnerGuard implements CanActivate {
    constructor(
        private router: Router,
        private userService: UserServiceService
    ) { }

    canActivate() {
        let role = localStorage.getItem('role');
        if (role === 'OWNER') { 
             return true;
        }

        alert('You do not have owner permissions.')
        this.router.navigate(['/user/landing']); 
        return false;
    }
}