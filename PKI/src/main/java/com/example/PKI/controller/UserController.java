package com.example.PKI.controller;

import com.example.PKI.dto.*;
import com.example.PKI.model.Role;
import com.example.PKI.model.User;
import com.example.PKI.model.UserDetails;
import com.example.PKI.model.CustomToken;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.security.TokenUtils;
import com.example.PKI.service.LoggerService;
import com.example.PKI.service.UserService;
import com.example.PKI.service.CustomTokenService;
import com.example.PKI.service.impl.LoggerServiceImpl;
import com.github.nbaars.pwnedpasswords4j.client.PwnedPasswordChecker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.net.URI;

@CrossOrigin(origins = "http://localhost:4200")
@RestController
@RequestMapping(value = "/api")
public class UserController {

    private UserService userService;

    private PwnedPasswordChecker checker;

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private CustomTokenService customTokenService;

    private final LoggerService loggerService;

    @Autowired
    public UserController(UserService userService) {
        this.checker = PwnedPasswordChecker.standalone("My user agent");
        this.userService = userService;
        this.loggerService = new LoggerServiceImpl(this.getClass());
    }

    @PostMapping("/login")
    public ResponseEntity<LoggedUserDto> login(@Valid
            @RequestBody JwtAuthenticationRequest authenticationRequest, HttpServletResponse response,HttpServletRequest request) {
        try {
            Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                    authenticationRequest.getEmail(), authenticationRequest.getPassword()));
            SecurityContextHolder.getContext().setAuthentication(authentication);
            Role role = userService.findByEmail(authenticationRequest.getEmail()).getRole();
            UserDetails user = (UserDetails) authentication.getPrincipal();
            String jwt = tokenUtils.generateToken(user.getUser().getEmail());
            int expiresIn = tokenUtils.getExpiredIn();
            LoggedUserDto loggedUserDto = new LoggedUserDto(authenticationRequest.getEmail(), role.toString(), new UserTokenState(jwt, expiresIn));
            loggerService.loginSuccess(authenticationRequest.getEmail());
            return ResponseEntity.ok(loggedUserDto);
        }catch (Exception ex) {
            loggerService.loginFailed(authenticationRequest.getEmail(),request.getRemoteAddr());
            return ResponseEntity.badRequest().build();
        }
    }

    @PostMapping("/createSubject")
    public ResponseEntity<UserDto> createUser(@Valid @RequestBody UserDto userDto,HttpServletRequest request) {
        UserDto newUser = userService.createUser(userDto);
        if(newUser == null){
            loggerService.signUpFailed(request.getRemoteAddr());
            return new ResponseEntity<UserDto>((UserDto) null,HttpStatus.BAD_REQUEST);
        }
        newUser.isPawned = checker.check(userDto.getPassword());
        loggerService.signUpSuccess(newUser.getEmail(), request.getRemoteAddr());
        return new ResponseEntity<UserDto>(newUser, HttpStatus.OK);
    }

    @PreAuthorize("hasAuthority('ADMIN') || hasAuthority('USER_ROOT') || hasAuthority('USER_INTERMEDIATE')")
    @GetMapping("/users")
    public ResponseEntity<?> getAll() {
        return new ResponseEntity<ArrayList<User>>((ArrayList<User>) userRepository.findAll(), HttpStatus.OK);
    }

    @GetMapping("/confirmAccount/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token,HttpServletRequest request) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        User user = verificationToken.getUser();
        if (verificationToken.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            loggerService.confirmAccountFailed(user.getEmail(),request.getRemoteAddr());
            return new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if (activated.isActivated()) {
            loggerService.confirmAccountSuccess(user.getEmail(),request.getRemoteAddr());
            return ResponseEntity.status(HttpStatus.FOUND)
                    .location(URI.create("http://localhost:4200/")).build();
        } else {
            loggerService.confirmAccountFailed(user.getEmail(),request.getRemoteAddr());
            return new ResponseEntity<>("Error happened!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @CrossOrigin(origins = "http://localhost:4200")
    //ovo moze da radi bilo koji ulogovani korisnik
    @PutMapping(value = "/changePassword")
    public ResponseEntity<HttpStatus> changePassword(@Valid @RequestBody ChangePasswordDto changePasswordDto, HttpServletRequest request) {
        String token = tokenUtils.getToken(request);
        userService.changePassword(tokenUtils.getEmailFromToken(token), changePasswordDto);
        loggerService.changePasswordSuccess(tokenUtils.getEmailFromToken(token),request.getRemoteAddr());
        return ResponseEntity.noContent().build();
    }

    @PostMapping(value = "/sendCode")
    public ResponseEntity<?> sendCode(@RequestBody String email) {
        User user = userService.findByEmail(email);
        if (user == null) {
            loggerService.sendCodeFailed(user.getEmail());
            return ResponseEntity.notFound().build();
        }
        loggerService.sendCodeSuccess(user.getEmail());
        customTokenService.sendResetPasswordToken(user);
        return ResponseEntity.accepted().build();
    }

    //ovoj metodi mogu svi da pristupe
    @PostMapping(value = "/checkRecoveryEmail")
    public ResponseEntity<String> checkRecoveryEmail(@Valid @RequestBody RequestCheckDto request,HttpServletRequest httpServletRequest) {
        User user = userService.findByEmail(request.getEmail());
        if (user.getRecoveryEmail().equals(request.getRecoveryEmail())) {
            customTokenService.sendResetPasswordToken(user);
            loggerService.checkRecoveryEmailSuccess(user.getEmail(),httpServletRequest.getRemoteAddr());
            return new ResponseEntity<>("Check your email.", HttpStatus.OK);
        }
        loggerService.checkRecoveryEmailFailed(user.getEmail(),httpServletRequest.getRemoteAddr());
        return new ResponseEntity<>("Entered recovery email is not valid!", HttpStatus.BAD_REQUEST);
    }

    @PostMapping(value = "/checkCode")
    public ResponseEntity<String> checkCode(@Valid @RequestBody CheckCodeDto checkCodeDto,HttpServletRequest request) {
        User user = userService.findByEmail(checkCodeDto.getEmail());
        CustomToken token = customTokenService.findByUser(user);
        if (customTokenService.checkResetPasswordCode(checkCodeDto.getCode(), token.getToken())) {
            if (token.getExpiryDate().isBefore(LocalDateTime.now())) {
                customTokenService.deleteById(token.getId());
                customTokenService.sendVerificationToken(user);
                loggerService.checkCodeFailed(user.getEmail(),request.getRemoteAddr());
                return new ResponseEntity<>("Reset password code is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
            }

            loggerService.checkCodeSuccess(user.getEmail(),request.getRemoteAddr());
            customTokenService.deleteById(token.getId());
            return new ResponseEntity<>("\"Success!\"", HttpStatus.OK);
        }

        loggerService.checkCodeFailed(user.getEmail(),request.getRemoteAddr());
        return new ResponseEntity<>("Entered code is not valid!", HttpStatus.BAD_REQUEST);
    }


    @PostMapping(value = "/resetPassword")
    public ResponseEntity<String> resetPassword(@Valid @RequestBody ResetPasswordDto resetPasswordDto,HttpServletRequest request) {
        userService.resetPassword(resetPasswordDto.getEmail(), resetPasswordDto.getNewPassword());
        loggerService.resetPasswordSuccess(resetPasswordDto.getEmail(),request.getRemoteAddr());
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

}
