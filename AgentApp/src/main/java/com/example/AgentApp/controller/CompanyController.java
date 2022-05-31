package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.http.*;
import org.springframework.web.bind.annotation.*;

import javax.annotation.security.*;

@RequestMapping("/company")
@RestController
public class CompanyController {

    private final CompanyService companyService;
    private final CompanyMapper companyMapper;
    private final CommentMapper commentMapper;
    private final CommentService commentService;

    public CompanyController(CommentMapper commentMapper, CompanyMapper companyMapper, CompanyService companyService, CommentService commentService) {
        this.commentMapper = commentMapper;
        this.companyMapper = companyMapper;
        this.companyService = companyService;
        this.commentService = commentService;
    }

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

    @PostMapping("/comment")
    public ResponseEntity<Object> leaveAComment(@RequestBody CommentDto commentDto){
        Company company = companyService.getById(commentDto.getCompanyId());
        Comment comment = commentMapper.toEntity(commentDto);
        comment.setCompany(company);
        return ResponseEntity.ok(commentService.create(comment));
    }

}
