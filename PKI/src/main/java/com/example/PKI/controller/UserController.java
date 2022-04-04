package com.example.PKI.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.PKI.dto.LoginDTO;
import com.example.PKI.model.User;
import com.example.PKI.service.UserService;

@RestController
@RequestMapping(value = "/api")
public class UserController {
	
	private UserService userService;
	
	@Autowired
	public UserController(UserService userService) {
		this.userService = userService;
	}
	
	@PostMapping("/login")
    public ResponseEntity<String> login(@RequestBody LoginDTO loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		if(userService.login(loginDTO)) {
			if(user.isAdmin()) {
				return new ResponseEntity<String>("admin", HttpStatus.OK);
			}else {
				return new ResponseEntity<String>("user", HttpStatus.OK);
			}
		}
        return new ResponseEntity<String>("Invalid login", HttpStatus.BAD_REQUEST);
        
    }
	
	@PostMapping("/logout")
    public ResponseEntity<String> logout(@RequestBody LoginDTO loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		userService.logout(user.getEmail());
        return new ResponseEntity<String>("Success logout", HttpStatus.OK);  
    }

}
