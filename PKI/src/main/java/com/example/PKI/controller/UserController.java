package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.Role;
import com.example.PKI.model.User;
import com.example.PKI.model.UserDetails;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.security.TokenUtils;
import com.example.PKI.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletResponse;
import java.lang.reflect.Array;
import java.util.ArrayList;

@RestController
@RequestMapping(value = "/api")
public class UserController {

    private UserService userService;

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    public UserController(UserService userService) {
        this.userService = userService;
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping("/login")
    public ResponseEntity<LoggedUserDto> login(
            @RequestBody JwtAuthenticationRequest authenticationRequest, HttpServletResponse response) {

        Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                authenticationRequest.getEmail(), authenticationRequest.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);
        Role role = userService.findByEmail(authenticationRequest.getEmail()).getRole();
        UserDetails user = (UserDetails) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUser().getEmail());
        int expiresIn = tokenUtils.getExpiredIn();
        LoggedUserDto loggedUserDto = new LoggedUserDto(authenticationRequest.getEmail(),role.toString(),new UserTokenState(jwt,expiresIn));
        return ResponseEntity.ok(loggedUserDto);
    }


    //@CrossOrigin(origins = "*")
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
