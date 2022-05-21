package com.example.PKI.controller;

import com.example.PKI.dto.LoggedUserDto;
import com.example.PKI.dto.LoginDto;
import com.example.PKI.dto.UserDto;
import com.example.PKI.model.User;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.lang.reflect.Array;
import java.util.ArrayList;

@RestController
@RequestMapping(value = "/api")
public class UserController {

    private UserService userService;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping("/login")
    public ResponseEntity<LoggedUserDto> login(@RequestBody LoginDto loginDTO) {
        User user = userService.findByEmail(loginDTO.getEmail());
        if (userService.login(loginDTO)) {
           // if (user.isAdmin()) {
                LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(), "admin");
                return new ResponseEntity<LoggedUserDto>(logedUser, HttpStatus.OK);
            /*} else {
                LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(), "user");
                return new ResponseEntity<LoggedUserDto>(logedUser, HttpStatus.OK);
            }*/
        }
        LoggedUserDto logedUser = new LoggedUserDto(loginDTO.getEmail(), "error");
        return new ResponseEntity<LoggedUserDto>(logedUser, HttpStatus.BAD_REQUEST);

    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping("/createSubject")
    public ResponseEntity<UserDto> createUser(@RequestBody UserDto userDto) {

        UserDto newUser = userService.createUser(userDto);
        return new ResponseEntity<UserDto>(newUser, HttpStatus.OK);
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @GetMapping("/users")
    public ResponseEntity<?> getAll() {
        return new ResponseEntity<ArrayList<User>>((ArrayList<User>) userRepository.findAll(), HttpStatus.OK);
    }
}
