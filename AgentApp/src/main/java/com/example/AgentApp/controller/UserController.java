package com.example.AgentApp.controller;


import com.example.AgentApp.dto.ChangePasswordDto;
import com.example.AgentApp.dto.UserInformationResponseDto;
import com.example.AgentApp.mapper.UserMapper;
import com.example.AgentApp.model.User;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;

@RestController
@CrossOrigin(origins = "http://localhost:4200")
@RequestMapping(value = "/api/user")
public class UserController {
    private final TokenUtils tokenUtils;
    private final UserService userService;
    
    public UserController(TokenUtils tokenUtils, UserService userService) {
        this.tokenUtils = tokenUtils;
        this.userService = userService;
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @PutMapping(value = "/change-password")
    public ResponseEntity<HttpStatus> changePassword(@Valid @RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        userService.changePassword(tokenUtils.getUsernameFromToken(token), changePasswordDto);
        return ResponseEntity.noContent().build();
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping(value = "/user-info")
    public ResponseEntity<?> getUserInformation(HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        String username = tokenUtils.getUsernameFromToken(token);
        User user = userService.findByUsername(username);
        return new ResponseEntity<UserInformationResponseDto>(UserMapper.mapToDto(user), HttpStatus.OK);
    }
}
