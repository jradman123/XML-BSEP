package com.example.AgentApp.controller;

import com.example.AgentApp.dto.JwtAuthenticationRequest;
import com.example.AgentApp.dto.LoggedUserDto;
import com.example.AgentApp.dto.UserTokenState;
import com.example.AgentApp.enums.UserRole;
import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.UserService;
import com.github.nbaars.pwnedpasswords4j.client.PwnedPasswordChecker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletResponse;

@RestController
@RequestMapping(value = "/api")
public class AuthenticationController{

    private UserService userService;

    private PwnedPasswordChecker checker;

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;



    @Autowired
    public AuthenticationController(UserService userService) {
        this.checker = PwnedPasswordChecker.standalone("My user agent");
        this.userService = userService;
    }

    @PostMapping("/login")
    public ResponseEntity<LoggedUserDto> login(
            @RequestBody JwtAuthenticationRequest authenticationRequest, HttpServletResponse response) {

        Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                authenticationRequest.getUsername(), authenticationRequest.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);
        UserRole role = userService.findByUsername(authenticationRequest.getUsername()).getRole();
        UserDetails user = (UserDetails) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUser().getUsername());
        int expiresIn = tokenUtils.getExpiredIn();
        LoggedUserDto loggedUserDto = new LoggedUserDto(authenticationRequest.getUsername(), role.toString(), new UserTokenState(jwt, expiresIn));
        return ResponseEntity.ok(loggedUserDto);
    }
}
