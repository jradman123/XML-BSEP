package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.Role;
import com.example.PKI.model.User;
import com.example.PKI.model.UserDetails;
import com.example.PKI.model.CustomToken;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.security.TokenUtils;
import com.example.PKI.service.UserService;
import com.example.PKI.service.CustomTokenService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.time.LocalDate;
import java.time.LocalDateTime;
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
    private CustomTokenService customTokenService;

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

    @CrossOrigin(origins = "*")
    @GetMapping("/confirmAccount/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        User user = verificationToken.getUser();
        if(verificationToken.getExpiryDate().isBefore(LocalDateTime.now())){
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            return  new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.",HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if(activated.isActivated()) {
            return new ResponseEntity<>("Account is activated.", HttpStatus.OK);
        }else{
            return new ResponseEntity<>("Error happened!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @CrossOrigin(origins = "http://localhost:4200")
   //ovo moze da radi bilo koji ulogovani korisnik
    @PutMapping(value = "/changePassword")
    public ResponseEntity<HttpStatus> changePassword(@RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        userService.changePassword(tokenUtils.getEmailFromToken(token), changePasswordDto);
        return ResponseEntity.noContent().build();
    }


    @CrossOrigin(origins = "http://localhost:4200")
    //ovoj metodi mogu svi da pristupe
    @PostMapping(value = "/checkRecoveryEmail")
    public ResponseEntity<String> checkRecoveryEmail(@RequestBody RequestCheckDto request) {
        User user = userService.findByEmail(request.getEmail());
        if(user.getRecoveryEmail().equals(request.getRecoveryEmail())){
            customTokenService.sendResetPasswordToken(user);
            return new ResponseEntity<>("Check your email.",HttpStatus.OK);
        }

        return new ResponseEntity<>("Entered recovery email is not valid!", HttpStatus.BAD_REQUEST);
    }

    @CrossOrigin(origins = "http://localhost:4200")
    //ovoj metodi mogu svi da pristupe
    @PostMapping(value = "/checkCode")
    public ResponseEntity<String> checkCode(@RequestBody CheckCodeDto checkCodeDto) {
        User user = userService.findByEmail(checkCodeDto.getEmail());
        CustomToken token = customTokenService.findByUser(user);
        if(customTokenService.checkResetPasswordCode(checkCodeDto.getCode(),token.getToken())){
            return new ResponseEntity<>("Success!",HttpStatus.OK);
        }

        return new ResponseEntity<>("Entered code is not valid!", HttpStatus.BAD_REQUEST);
    }

    @CrossOrigin(origins = "http://localhost:4200")
    //ovoj metodi mogu svi da pristupe
    @PostMapping(value = "/resetPassword")
    public ResponseEntity<String> resetPassword(@RequestBody ResetPasswordDto resetPasswordDto) {
        userService.resetPassword(resetPasswordDto.getEmail(),resetPasswordDto.getNewPassword());
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

}
