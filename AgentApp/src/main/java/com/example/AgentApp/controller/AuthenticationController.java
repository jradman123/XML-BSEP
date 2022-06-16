package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.UserRole;
import com.example.AgentApp.exception.ResourceConflictException;
import com.example.AgentApp.model.CustomToken;
import com.example.AgentApp.model.User;
import com.example.AgentApp.model.UserDetails;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.CustomTokenService;
import com.example.AgentApp.service.LoggerService;
import com.example.AgentApp.service.UserService;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import de.taimos.totp.TOTP;
import org.apache.commons.codec.binary.Base32;
import org.apache.commons.codec.binary.Hex;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.UriComponentsBuilder;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.validation.Valid;
import java.net.URI;
import java.net.UnknownHostException;
import java.text.ParseException;
import java.time.LocalDateTime;

@RestController
@CrossOrigin(origins = "https://localhost:4200")
@RequestMapping(value = "api/auth")
public class AuthenticationController {

    private final UserService userService;

    private final TokenUtils tokenUtils;

    private final AuthenticationManager authenticationManager;

    private final CustomTokenService customTokenService;

    private final LoggerService loggerService;


    @Autowired
    public AuthenticationController(UserService userService, TokenUtils tokenUtils, CustomTokenService customTokenService, AuthenticationManager authenticationManager) {
        this.userService = userService;
        this.tokenUtils = tokenUtils;
        this.customTokenService = customTokenService;
        this.authenticationManager = authenticationManager;
        this.loggerService = new LoggerServiceImpl(this.getClass());
    }

    @PostMapping("/login")
    public ResponseEntity<Object> login(
            @RequestBody JwtAuthenticationRequest authenticationRequest, HttpServletResponse response,HttpServletRequest request) {

        try
            { Authentication authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(
                    authenticationRequest.getUsername(), authenticationRequest.getPassword()));

            User u = userService.findByUsername(authenticationRequest.getUsername());
            String code = authenticationRequest.getCode();
            if (u.isUsing2FA() && (code == null || !code.equals(getTOTPCode(u.getSecret())))) {
                return ResponseEntity.badRequest().body("Code invalid");
            }
            SecurityContextHolder.getContext().setAuthentication(authentication);
            UserRole role = userService.findByUsername(authenticationRequest.getUsername()).getRole();
            UserDetails user = (UserDetails) authentication.getPrincipal();
            String jwt = tokenUtils.generateToken(user.getUser().getUsername());
            int expiresIn = tokenUtils.getExpiredIn();
            LoggedUserDto loggedUserDto = new LoggedUserDto(authenticationRequest.getUsername(), role.toString(), new UserTokenState(jwt, expiresIn));
            loggerService.loginSuccess(authenticationRequest.getUsername());
            return ResponseEntity.ok(loggedUserDto);
            }
        catch (Exception ex) {
            loggerService.loginFailed(authenticationRequest.getUsername(),request.getRemoteAddr());
            return ResponseEntity.badRequest().build();
        }
    }

    public static String getTOTPCode(String secretKey) {
        Base32 base32 = new Base32();
        byte[] bytes = base32.decode(secretKey);
        String hexKey = Hex.encodeHexString(bytes);
        return TOTP.getOTP(hexKey);
    }

    @PostMapping("/signup")
    public ResponseEntity<String> addUser(@Valid @RequestBody RegistrationRequestDto userRequest, UriComponentsBuilder ucBuilder, HttpServletRequest request) throws UnknownHostException, ParseException {
        User existUser = this.userService.findByUsername(userRequest.getUsername());

        if (existUser != null) {
            loggerService.signUpFailed(request.getRemoteAddr());
            throw new ResourceConflictException(userRequest.getUsername(), "Username already exists");
        }

        User savedUser = userService.addUser(userRequest);
        if (savedUser != null) {
            loggerService.signUpSuccess(userRequest.getUsername(), request.getRemoteAddr());
            return new ResponseEntity<>("SUCCESS!", HttpStatus.CREATED);
        }
        loggerService.signUpFailed(request.getRemoteAddr());
        return new ResponseEntity<>("ERROR!", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @GetMapping("/confirm-account/{token}")
    public ResponseEntity<String> confirmAccount(@PathVariable String token,HttpServletRequest request) {
        CustomToken verificationToken = customTokenService.findByToken(token);
        User user = verificationToken.getUser();
        if (verificationToken.getExpiryDate().isBefore(LocalDateTime.now())) {
            customTokenService.deleteById(verificationToken.getId());
            customTokenService.sendVerificationToken(user);
            loggerService.confirmAccountFailed(user.getUsername(),request.getRemoteAddr());
            return new ResponseEntity<>("Confirmation link is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
        }
        User activated = userService.activateAccount(user);
        customTokenService.deleteById(verificationToken.getId());
        if (activated.isConfirmed()) {
            loggerService.confirmAccountSuccess(user.getUsername(),request.getRemoteAddr());
            return ResponseEntity.status(HttpStatus.FOUND)
                    .location(URI.create("http://localhost:4200/login")).build();

        } else {
            loggerService.confirmAccountFailed(user.getUsername(),request.getRemoteAddr());
            return new ResponseEntity<>("Error happened!", HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @PostMapping(value = "/send-code")
    public ResponseEntity<?> sendCode(@RequestBody String username) {
        User user = userService.findByUsername(username);
        if (user == null) {
            loggerService.sendCodeFailed(user.getEmail());
            return ResponseEntity.notFound().build();
        }
        loggerService.sendCodeSuccess(user.getEmail());
        customTokenService.sendResetPasswordToken(user);
        return ResponseEntity.accepted().build();
    }

    @PostMapping(value = "/check-code")
    public ResponseEntity<String> checkCode(@Valid @RequestBody CheckCodeDto checkCodeDto,HttpServletRequest request) {
        User user = userService.findByUsername(checkCodeDto.getUsername());
        CustomToken token = customTokenService.findByUser(user);
        if (customTokenService.checkResetPasswordCode(checkCodeDto.getCode(), token.getToken())) {
            if (token.getExpiryDate().isBefore(LocalDateTime.now())) {
                customTokenService.deleteById(token.getId());
                customTokenService.sendVerificationToken(user);
                loggerService.checkCodeFailed(user.getUsername(),request.getRemoteAddr());
                return new ResponseEntity<>("Reset password code is expired,we sent you new one.Please check you mail box.", HttpStatus.BAD_REQUEST);
            }

            loggerService.checkCodeSuccess(user.getUsername(),request.getRemoteAddr());
            customTokenService.deleteById(token.getId());
            return new ResponseEntity<>("\"Success!\"", HttpStatus.OK);
        }
        loggerService.checkCodeFailed(user.getUsername(),request.getRemoteAddr());
        return new ResponseEntity<>("Entered code is not valid!", HttpStatus.BAD_REQUEST);
    }

    @PostMapping(value = "/reset-password")
    public ResponseEntity<String> resetPassword(@Valid @RequestBody ResetPasswordDto resetPasswordDto,HttpServletRequest request) {
        userService.resetPassword(resetPasswordDto.getUsername(), resetPasswordDto.getNewPassword());
        loggerService.resetPasswordSuccess(resetPasswordDto.getUsername(),request.getRemoteAddr());
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

    @PutMapping(value = "/two-factor-auth")
    public ResponseEntity<SecretDto> change2FAStatus(@RequestBody Change2FAStatusDto dto) {
        String secret = userService.change2FAStatus(dto.username, dto.status);
        return ResponseEntity.ok(new SecretDto(secret));
    }
    @GetMapping(value= "/two-factor-auth-status/{username}")
    public ResponseEntity<Boolean> check2FAStatus(@PathVariable String username, HttpServletRequest request) {
            boolean twoFAEnabled = userService.check2FAStatus(username);
            return ResponseEntity.ok(twoFAEnabled);
    }
}