package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

import javax.annotation.security.*;
import java.util.*;

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

    //korisnik
    @PostMapping("/new")
    public ResponseEntity<?> createCompanyRequest(@RequestBody NewCompanyRequestDto requestDto){
        Company newCompany = companyService.createCompany(requestDto);
        if (newCompany != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(newCompany), HttpStatus.CREATED);
        }
        return new ResponseEntity<>("Failed to create company registration request!", HttpStatus.CONFLICT);
    }

    //admin
    @GetMapping("approve/{id}")
    public ResponseEntity<?> approveCompany(@PathVariable Long id) {
        Company company = companyService.approveCompany(id,true);
        if (company != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(company), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to approve company!", HttpStatus.CONFLICT);

    }
    //admin
    @GetMapping("reject/{id}")
    public ResponseEntity<?> rejectCompany(@PathVariable Long id) {
        Company company = companyService.approveCompany(id,false);
        if (company != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(company), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to reject company!", HttpStatus.CONFLICT);
    }

    //owner
    @PutMapping("edit/{id}")
    public ResponseEntity<?> editCompany(@PathVariable Long id, @RequestBody EditCompanyRequestDto requestDto){
        Company company = companyService.editCompany(requestDto,id);
        if (company != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(company), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to edit company!", HttpStatus.CONFLICT);
    }

    //owner
    @PostMapping("createOffer")
    public ResponseEntity<?> crateJobOffer(@RequestBody CreateJobOfferRequestDto requestDto){
        Company company = companyService.addJobOffer(requestDto);
        if (company != null){
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(company), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }

    //admin
    @GetMapping("pending")
    public ResponseEntity<?> getAllPendingCompanies(){
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.PENDING);
        if (companies != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos( companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }

    //svi
    @GetMapping("approved")
    public ResponseEntity<?> getAllApprovedCompanies(){
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.APPROVED);
        if (companies != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos(companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }


}
