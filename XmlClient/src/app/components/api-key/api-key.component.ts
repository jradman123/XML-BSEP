import { Component, OnInit } from '@angular/core';
import { ApiKeyService } from 'src/app/services/api-key-service/api-key.service';

@Component({
  selector: 'app-api-key',
  templateUrl: './api-key.component.html',
  styleUrls: ['./api-key.component.css']
})
export class ApiKeyComponent implements OnInit {

  apiKey! : string;

  constructor(private apiKeyService : ApiKeyService) { 
  }

  ngOnInit(): void {
  }

  generate(){
   this.apiKeyService.generateApiKey(localStorage.getItem('username')!).subscribe(
     res => this.apiKey = res
   )
  }

}
