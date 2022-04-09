package com.example.PKI.controller;

import com.example.PKI.dto.LoggedUserDto;
import com.example.PKI.dto.LoginDto;
import com.example.PKI.model.User;
import com.example.PKI.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

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
    public ResponseEntity<LoggedUserDto> login(@RequestBody LoginDto loginDto) {
        String mail = loginDto.getEmail();
        User user = userService.findByEmail(loginDto.getEmail());
        if (userService.login(loginDto)) {
            LoggedUserDto loggedUserDto;
            if (user.isAdmin()) {
                loggedUserDto = new LoggedUserDto(loginDto.getEmail(), "longedUser");
            } else {
                loggedUserDto = new LoggedUserDto(loginDto.getEmail(), "user");
            }
            return new ResponseEntity<LoggedUserDto>(loggedUserDto, HttpStatus.OK);
        }
        LoggedUserDto loggedUser = new LoggedUserDto(loginDto.getEmail(), "error");
        return new ResponseEntity<LoggedUserDto>(loggedUser, HttpStatus.BAD_REQUEST);

    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping("/logout")
    public ResponseEntity<String> logout(@RequestBody LoginDto loginDTO) {
        User user = userService.findByEmail(loginDTO.getEmail());
        userService.logout(user.getEmail());
        return new ResponseEntity<String>("Success logout", HttpStatus.OK);
    }

}
