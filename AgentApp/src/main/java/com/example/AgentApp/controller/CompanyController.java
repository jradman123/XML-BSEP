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

    private final CompanyService companyService;
    private final CompanyMapper companyMapper;
    private final CommentMapper commentMapper;
    private final CommentService commentService;
    private final UserService userService;
    private final SalaryCommentMapper salaryCommentMapper;
    private final SalaryCommentService salaryCommentService;
    private final InterviewService interviewService;
    private final InterviewMapper interviewMapper;

    public CompanyController(CommentMapper commentMapper, CompanyMapper companyMapper,
                             CompanyService companyService, CommentService commentService,
                             UserService userService, SalaryCommentMapper salaryCommentMapper,
                             SalaryCommentService salaryCommentService,InterviewService interviewService,
                             InterviewMapper interviewMapper) {
        this.commentMapper = commentMapper;
        this.companyMapper = companyMapper;
        this.companyService = companyService;
        this.commentService = commentService;
        this.userService = userService;
        this.salaryCommentMapper = salaryCommentMapper;
        this.salaryCommentService = salaryCommentService;
        this.interviewService = interviewService;
        this.interviewMapper = interviewMapper;
    }

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

    @PostMapping("/comment")
    public ResponseEntity<?> leaveAComment(@RequestBody CommentDto commentDto){
        Company company = companyService.getById(commentDto.getCompanyId());
        Comment comment = commentMapper.toEntity(commentDto);
        User user = userService.findByUsername(commentDto.userUsername);
        comment.setCompany(company);
        comment.setUser(user);
        Comment savedComment = commentService.create(comment);
        Set<Comment> allCommentsForCompany = commentService.getAllForCompany(commentDto.getCompanyId());
        if (savedComment != null && allCommentsForCompany != null){
            return new ResponseEntity<List<CommentResponseDto>>(commentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add comment for company!", HttpStatus.CONFLICT);
    }

    @GetMapping("{id}/comments")
    public ResponseEntity<?> allComments(@PathVariable Long id){
        Set<Comment> allCommentsForCompany = commentService.getAllForCompany(id);
        if (allCommentsForCompany != null){
            return new ResponseEntity<List<CommentResponseDto>>(commentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all comments for company!", HttpStatus.CONFLICT);
    }

    @PostMapping("/salaryComment")
    public  ResponseEntity<?> leaveSalaryComment(@RequestBody SalaryCommentRequestDto commentDto){
        Company company = companyService.getById(commentDto.companyID);
        SalaryComment comment = salaryCommentMapper.mapToEntity(commentDto);
        SalaryComment savedComment = salaryCommentService.create(comment);
        Set<SalaryComment> allCommentsForCompany = salaryCommentService.getAllForCompany(commentDto.companyID);
        if (savedComment != null && allCommentsForCompany != null){
            return new ResponseEntity<List<SalaryCommentResponseDto>>(salaryCommentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add salary comment for company!", HttpStatus.CONFLICT);
    }

    @GetMapping("{id}/salaryComments")
    public ResponseEntity<?> allSalaryComments(@PathVariable Long id){
        Set<SalaryComment> allCommentsForCompany = salaryCommentService.getAllForCompany(id);
        if ( allCommentsForCompany != null){
            return new ResponseEntity<List<SalaryCommentResponseDto>>(salaryCommentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all salary comments for company!", HttpStatus.CONFLICT);
    }

    @PostMapping("/interview")
    public  ResponseEntity<?> leaveInterviewComment(@RequestBody InterviewRequestDto commentDto){
        Company company = companyService.getById(commentDto.companyID);
        Interview interview = interviewMapper.mapToEntity(commentDto);
        Interview savedInterview = interviewService.create(interview);
        Set<Interview> allInterviewsCompany = interviewService.getAllForCompany(commentDto.companyID);
        if (savedInterview != null && allInterviewsCompany != null){
            return new ResponseEntity<List<InterviewResponseDto>>(interviewMapper.mapToDtos(allInterviewsCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add interview for company!", HttpStatus.CONFLICT);
    }

    @GetMapping("{id}/interviews")
    public ResponseEntity<?> allInterviews(@PathVariable Long id){
        Set<Interview> allInterviewsCompany = interviewService.getAllForCompany(id);
        if ( allInterviewsCompany != null){
            return new ResponseEntity<List<InterviewResponseDto>>(interviewMapper.mapToDtos(allInterviewsCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all interviews for company!", HttpStatus.CONFLICT);
    }

}
