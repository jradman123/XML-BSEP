package com.example.PKI.controller;

import com.example.PKI.dto.LoggedUserDto;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import com.example.PKI.dto.LoginDto;
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

	@CrossOrigin(origins = "http://localhost:4200")
	@PostMapping("/login")
    public ResponseEntity<LoggedUserDto> login(@RequestBody LoginDto loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		if(userService.login(loginDTO)) {
			if(user.isAdmin()) {
				LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(),"admin");
				return new ResponseEntity<LoggedUserDto>(logedUser, HttpStatus.OK);
			}else {
				LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(),"user");
				return new ResponseEntity<LoggedUserDto>(logedUser, HttpStatus.OK);
			}
		}
		LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(),"error");
        return new ResponseEntity<LoggedUserDto>(logedUser,HttpStatus.BAD_REQUEST);
        
    }

	@CrossOrigin(origins = "http://localhost:4200")
	@PostMapping("/logout")
    public ResponseEntity<String> logout(@RequestBody LoginDto loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		userService.logout(user.getEmail());
        return new ResponseEntity<String>("Success logout", HttpStatus.OK);  
    }

}
