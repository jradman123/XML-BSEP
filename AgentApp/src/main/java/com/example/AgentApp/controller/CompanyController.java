package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

import javax.annotation.security.*;

@RequestMapping("/company")
@RestController
public class CompanyController {

    @Autowired
    private CompanyService companyService;
    @Autowired
    private CompanyMapper companyMapper;

    @GetMapping("")
    public ResponseEntity<String> getAll(){
        return  new ResponseEntity<>("ok", HttpStatus.OK);
    }

    @PermitAll
    @PostMapping("/new")
    public ResponseEntity<?> createCompanyRequest(@RequestBody NewCompanyRequestDto requestDto){
        Company newCompany = companyService.createCompany(requestDto);
        if (newCompany != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(newCompany), HttpStatus.CREATED);
        }
        return new ResponseEntity<>("Failed to create company registration request!", HttpStatus.CONFLICT);
    }

}
