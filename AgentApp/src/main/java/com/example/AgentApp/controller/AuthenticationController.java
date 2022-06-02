package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.UserRole;
import com.example.AgentApp.exception.ResourceConflictException;
import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.UriComponentsBuilder;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;
import java.net.UnknownHostException;
import java.text.ParseException;
import java.time.LocalDateTime;

@RestController
@RequestMapping(value = "/api")
public class AuthenticationController {

    @Autowired
    private UserService userService;

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CustomTokenService customTokenService;

    @CrossOrigin(origins = "http://localhost:4200")
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

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping("/signup")
    public ResponseEntity<String> addUser(@Valid @RequestBody RegistrationRequestDto userRequest, UriComponentsBuilder ucBuilder) throws UnknownHostException, ParseException {
        User existUser = this.userService.findByUsername(userRequest.getUsername());

        if (existUser != null) {
            throw new ResourceConflictException(userRequest.getUsername(), "Username already exists");
        }

        User savedUser = userService.addUser(userRequest);
        if (savedUser != null) {
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @GetMapping("/confirmAccount/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        User user = verificationToken.getUser();
        if (verificationToken.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            return new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if (activated.isConfirmed()) {
            return new ResponseEntity<>("Account is activated.", HttpStatus.OK);
        } else {
            return new ResponseEntity<>("Error happened!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping(value = "/sendCode")
    public ResponseEntity<?> sendCode(@RequestBody String username) {
        User user = userService.findByUsername(username);
        customTokenService.sendResetPasswordToken(user);
        return ResponseEntity.accepted().build();
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping(value = "/checkCode")
    public ResponseEntity<String> checkCode(@RequestBody CheckCodeDto checkCodeDto) {
        User user = userService.findByUsername(checkCodeDto.getUsername());
        CustomToken token = customTokenService.findByUser(user);
        if (customTokenService.checkResetPasswordCode(checkCodeDto.getCode(), token.getToken())) {
            return new ResponseEntity<>("Success!", HttpStatus.OK);
        }

        return new ResponseEntity<>("Entered code is not valid!", HttpStatus.BAD_REQUEST);
    }

    @CrossOrigin(origins = "http://localhost:4200")
    @PostMapping(value = "/resetPassword")
    public ResponseEntity<String> resetPassword(@Valid @RequestBody ResetPasswordDto resetPasswordDto) {
        userService.resetPassword(resetPasswordDto.getUsername(), resetPasswordDto.getNewPassword());
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }
}