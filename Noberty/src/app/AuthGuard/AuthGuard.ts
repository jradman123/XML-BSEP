import { Injectable } from "@angular/core";
import { CanActivate, Router } from "@angular/router";
import { UserService } from "../services/UserService/user.service";


@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
    constructor(
        private router: Router,
        private userService: UserService
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