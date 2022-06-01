import { Injectable } from "@angular/core";
import { CanActivate, Router } from "@angular/router";
import { UserServiceService } from "../services/UserService/user-service.service";


@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
    constructor(
        private router: Router,
        private userService: UserServiceService
    ) { }

    canActivate() {
        const currentUser = this.userService.currentUserValue;
        if (currentUser) { 
             return true;
        }

        this.router.navigate(['/login']); 
        return false;
    }
}