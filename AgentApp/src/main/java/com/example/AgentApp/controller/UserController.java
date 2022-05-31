package com.example.AgentApp.controller;


import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

@RestController
@RequestMapping(value = "/api/users")
public class UserController {

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private UserService userService;

    @PreAuthorize("hasAuthority('ADMIN') or hasAuthority('REGISTERED_USER') or hasAuthority('OWNER')")
    @PutMapping(value = "/changePassword")
    public ResponseEntity<HttpStatus> changePassword(@Valid @RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        userService.changePassword(tokenUtils.getUsernameFromToken(token), changePasswordDto);
        return ResponseEntity.noContent().build();
    }
}
