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
@RequestMapping(value = "/api/users")
public class UserController {

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private UserService userService;

    @Autowired
    private UserMapper userMapper;

    @CrossOrigin(origins = "http://localhost:4200")
    @PreAuthorize("hasAuthority('ADMIN') or hasAuthority('REGISTERED_USER') or hasAuthority('OWNER')")
    @PutMapping(value = "/changePassword")
    public ResponseEntity<HttpStatus> changePassword(@Valid @RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        userService.changePassword(tokenUtils.getUsernameFromToken(token), changePasswordDto);
        return ResponseEntity.noContent().build();
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PreAuthorize("hasAuthority('ADMIN') or hasAuthority('REGISTERED_USER') or hasAuthority('OWNER')")
    @GetMapping (value = "/getUserInformation")
    public ResponseEntity<?> getUserInformation(HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        String username=tokenUtils.getUsernameFromToken(token);
        User user = userService.findByUsername(username);
        return new ResponseEntity<UserInformationResponseDto>(userMapper.mapToDto(user),HttpStatus.OK);
    }
}
