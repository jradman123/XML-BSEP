import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from 'src/app/services/user-service/user.service';

@Component({
  selector: 'app-unauthenticated-header',
  templateUrl: './unauthenticated-header.component.html',
  styleUrls: ['./unauthenticated-header.component.css']
})
export class UnauthenticatedHeaderComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
