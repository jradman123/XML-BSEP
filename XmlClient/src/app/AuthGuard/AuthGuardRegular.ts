import { Injectable } from '@angular/core';
import { Router, CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { UserService } from 'src/app/services/user-service/user.service';

@Injectable({ providedIn: 'root' })
export class AuthGuardRegular implements CanActivate {
    constructor(
        private router: Router,
        private authenticationService: UserService
    ) { }

    canActivate() {
        const loggedIn = localStorage.getItem("token")
        const role = localStorage.getItem('role');
        if (loggedIn != null && role === 'Regular') {
            // logged in so return true
            return true;
        }
        // not logged in so redirect to login page with the return url
        this.router.navigate(['/login']);
        return false;
    }
}