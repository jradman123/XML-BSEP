package com.example.PKI.controller;

import com.example.PKI.dto.LogedUserDTO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

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

	@CrossOrigin(origins = "http://localhost:4200")
	@PostMapping("/login")
    public ResponseEntity<LogedUserDTO> login(@RequestBody LoginDTO loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		if(userService.login(loginDTO)) {
			if(user.isAdmin()) {
				LogedUserDTO logedUser = new LogedUserDTO(loginDTO.getEmail(),"admin");
				return new ResponseEntity<LogedUserDTO>(logedUser, HttpStatus.OK);
			}else {
				LogedUserDTO logedUser = new LogedUserDTO(loginDTO.getEmail(),"user");
				return new ResponseEntity<LogedUserDTO>(logedUser, HttpStatus.OK);
			}
		}
		LogedUserDTO logedUser = new LogedUserDTO(loginDTO.getEmail(),"error");
        return new ResponseEntity<LogedUserDTO>(logedUser,HttpStatus.BAD_REQUEST);
        
    }

	@CrossOrigin(origins = "http://localhost:4200")
	@PostMapping("/logout")
    public ResponseEntity<String> logout(@RequestBody LoginDTO loginDTO)  {
		User user = userService.findByEmail(loginDTO.getEmail());
		userService.logout(user.getEmail());
        return new ResponseEntity<String>("Success logout", HttpStatus.OK);  
    }

}
