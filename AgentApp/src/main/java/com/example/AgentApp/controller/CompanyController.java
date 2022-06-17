package com.example.AgentApp.controller;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.security.TokenUtils;
import com.example.AgentApp.service.*;
import com.example.AgentApp.service.impl.LoggerServiceImpl;
import org.springframework.http.*;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Valid;
import java.text.ParseException;
import java.util.*;

@CrossOrigin(origins = "https://localhost:4200")
@RequestMapping("/api/company")
@RestController
public class CompanyController {

    private final JobOfferService jobOfferService;
    private final CompanyService companyService;
    private final CompanyMapper companyMapper;
    private final CommentMapper commentMapper;
    private final CommentService commentService;
    private final UserService userService;
    private final SalaryCommentMapper salaryCommentMapper;
    private final SalaryCommentService salaryCommentService;
    private final InterviewService interviewService;
    private final InterviewMapper interviewMapper;
    private final TokenUtils tokenUtils;
    private final LoggerService loggerService;

    public CompanyController(CommentMapper commentMapper, CompanyMapper companyMapper,
                             CompanyService companyService, JobOfferService jobOfferService, CommentService commentService,
                             UserService userService, SalaryCommentMapper salaryCommentMapper,
                             SalaryCommentService salaryCommentService, InterviewService interviewService,
                             InterviewMapper interviewMapper, TokenUtils tokenUtils) {
        this.commentMapper = commentMapper;
        this.companyMapper = companyMapper;
        this.companyService = companyService;
        this.jobOfferService = jobOfferService;
        this.commentService = commentService;
        this.userService = userService;
        this.salaryCommentMapper = salaryCommentMapper;
        this.salaryCommentService = salaryCommentService;
        this.interviewService = interviewService;
        this.interviewMapper = interviewMapper;
        this.tokenUtils = tokenUtils;
        this.loggerService = new LoggerServiceImpl(this.getClass());
    }
    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("/{id}")
    public ResponseEntity<Object> getById(@PathVariable Long id){
        return ResponseEntity.ok(companyMapper.mapToDto(companyService.getById(id)));
    }

    @PreAuthorize("hasAnyAuthority('OWNER', 'REGISTERED_USER')")
    @PostMapping("/new")
    public ResponseEntity<?> createCompanyRequest(@RequestBody NewCompanyRequestDto requestDto){
        Company newCompany = companyService.createCompany(requestDto);
        if (newCompany != null){
            loggerService.createCompanyRequestSuccess(newCompany.getId().toString());
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(newCompany), HttpStatus.CREATED);
        }

        loggerService.createCompanyRequestFailed(newCompany.getId().toString());
        return new ResponseEntity<>("Failed to create company registration request!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAuthority('ADMIN')")
    @GetMapping("approve/{id}")
    public ResponseEntity<?> approveCompany(@PathVariable Long id,HttpServletRequest request) {
        Company company = companyService.approveCompany(id,true);
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.PENDING);
        if (company != null){
            loggerService.approveCompanyRequestSuccess(id.toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos( companies), HttpStatus.OK);
        }
        loggerService.approveCompanyRequestFailed(id.toString());
        return new ResponseEntity<>("Failed to approve company!", HttpStatus.CONFLICT);

    }

    @PreAuthorize("hasAuthority('ADMIN')")
    @GetMapping("reject/{id}")
    public ResponseEntity<?> rejectCompany(@PathVariable Long id,HttpServletRequest request) {
        Company company = companyService.approveCompany(id,false);
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.PENDING);
        if (company != null){
            loggerService.rejectCompanyRequestSuccess(id.toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos( companies), HttpStatus.OK);
        }
        loggerService.rejectCompanyRequestFailed(id.toString());
        return new ResponseEntity<>("Failed to reject company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAuthority('OWNER')")
    @PutMapping("edit/{id}")
    public ResponseEntity<?> editCompany(@PathVariable Long id, @RequestBody EditCompanyRequestDto requestDto,HttpServletRequest request){
        Company company = companyService.editCompany(requestDto,id);
        if (company != null){
            loggerService.editCompanySuccess(id.toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<NewCompanyResponseDto>(companyMapper.mapToCompanyCreateResponse(company), HttpStatus.OK);
        }
        loggerService.editCompanyFailed(id.toString());
        return new ResponseEntity<>("Failed to edit company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAuthority('OWNER')")
    @PostMapping("create-offer")
    public ResponseEntity<?> crateJobOffer(@RequestBody CreateJobOfferRequestDto requestDto,HttpServletRequest request) throws ParseException {
        Company company = companyService.addJobOffer(requestDto);
        Set<JobOffer> allOffers = jobOfferService.getAllOffersForCompany(requestDto.companyId);
        if (company != null){
            loggerService.createJobOfferSuccess(company.getId().toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<JobOfferResponseDto>>(JobOfferMapper.mapToDtos(allOffers),
                    HttpStatus.OK);
        }
        loggerService.createJobOfferFailed(company.getId().toString());
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAuthority('ADMIN')")
    @GetMapping("pending")
    public ResponseEntity<?> getAllPendingCompanies(){
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.PENDING);
        if (companies != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos( companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("approved")
    public ResponseEntity<?> getAllApprovedCompanies(){
        List<Company> companies = companyService.getAllCompaniesWithStatus(CompanyStatus.APPROVED);
        if (companies != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos(companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to add job offer to company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('OWNER', 'REGISTERED_USER')")
    @PostMapping("/comment")
    public ResponseEntity<?> leaveAComment(@Valid @RequestBody CommentDto commentDto,HttpServletRequest request){
        Company company = companyService.getById(commentDto.getCompanyId());
        Comment comment = commentMapper.toEntity(commentDto);
        Comment savedComment = commentService.create(comment);
        Set<Comment> allCommentsForCompany = commentService.getAllForCompany(commentDto.getCompanyId());
        if (savedComment != null && allCommentsForCompany != null){
            loggerService.leaveCommentSuccess(company.getId().toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<CommentResponseDto>>(commentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        loggerService.leaveCommentFailed(company.getId().toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
        return new ResponseEntity<>("Failed to add comment for company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("{id}/comments")
    public ResponseEntity<?> allComments(@PathVariable Long id){
        Set<Comment> allCommentsForCompany = commentService.getAllForCompany(id);
        if (allCommentsForCompany != null){
            return new ResponseEntity<List<CommentResponseDto>>(commentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all comments for company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('OWNER', 'REGISTERED_USER')")
    @PostMapping("/salary-comment")
    public  ResponseEntity<?> leaveSalaryComment(@RequestBody SalaryCommentRequestDto commentDto,HttpServletRequest request){
        Company company = companyService.getById(commentDto.companyID);
        SalaryComment comment = salaryCommentMapper.mapToEntity(commentDto);
        SalaryComment savedComment = salaryCommentService.create(comment);
        Set<SalaryComment> allCommentsForCompany = salaryCommentService.getAllForCompany(commentDto.companyID);
        if (savedComment != null && allCommentsForCompany != null){
            loggerService.leaveInterviewCommentSuccess(company.getId().toString(), tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<SalaryCommentResponseDto>>(salaryCommentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        loggerService.leaveSalaryCommentFailed(company.getId().toString(),tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
        return new ResponseEntity<>("Failed to add salary comment for company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("{id}/salary-comments")
    public ResponseEntity<?> allSalaryComments(@PathVariable Long id){
        Set<SalaryComment> allCommentsForCompany = salaryCommentService.getAllForCompany(id);
        if ( allCommentsForCompany != null){
            return new ResponseEntity<List<SalaryCommentResponseDto>>(salaryCommentMapper.mapToDtos(allCommentsForCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all salary comments for company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('OWNER', 'REGISTERED_USER')")
    @PostMapping("/interview")
    public  ResponseEntity<?> leaveInterviewComment(@RequestBody InterviewRequestDto commentDto,HttpServletRequest request){
        Company company = companyService.getById(commentDto.companyID);
        Interview interview = interviewMapper.mapToEntity(commentDto);
        Interview savedInterview = interviewService.create(interview);
        Set<Interview> allInterviewsCompany = interviewService.getAllForCompany(commentDto.companyID);
        if (savedInterview != null && allInterviewsCompany != null){
            loggerService.leaveInterviewCommentSuccess(company.getId().toString(), tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
            return new ResponseEntity<List<InterviewResponseDto>>(interviewMapper.mapToDtos(allInterviewsCompany), HttpStatus.OK);
        }
        loggerService.leaveInterviewCommentFailed(company.getId().toString(), tokenUtils.getUsernameFromToken(tokenUtils.getToken(request)));
        return new ResponseEntity<>("Failed to add interview for company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("{id}/interviews")
    public ResponseEntity<?> allInterviews(@PathVariable Long id){
        Set<Interview> allInterviewsCompany = interviewService.getAllForCompany(id);
        if ( allInterviewsCompany != null){
            return new ResponseEntity<List<InterviewResponseDto>>(interviewMapper.mapToDtos(allInterviewsCompany), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to get all interviews for company!", HttpStatus.CONFLICT);
    }

    //@PreAuthorize("hasAuthority('OWNER')")
    @GetMapping("/users-company/{username}")
    public ResponseEntity<List<CompanyResponseDto>> getUsersCompany(@PathVariable String username){
        Long id = userService.getByUsername(username);
        return ResponseEntity.ok(companyMapper.mapToDtos(companyService.getAllUsersCompanies(id)));
    }


    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("/search-companies")
    public ResponseEntity<?> getAllForUser(HttpServletRequest request){
        String username = tokenUtils.getUsernameFromToken(tokenUtils.getToken(request));
        User user = userService.findByUsername(username);
        List<Company> companies;
        if(user.getRole().equals(UserRole.OWNER)){
            companies = companyService.getAllApprovedCompaniesExceptOwners(user);
        }else {
            companies = companyService.getAllCompaniesWithStatus(CompanyStatus.APPROVED);
        }
        if (companies != null){
            return new ResponseEntity<List<CompanyResponseDto>>(companyMapper.mapToDtos(companies), HttpStatus.OK);
        }
        return new ResponseEntity<>("Failed to return any company!", HttpStatus.CONFLICT);
    }

    @PreAuthorize("hasAnyAuthority('ADMIN', 'OWNER', 'REGISTERED_USER')")
    @GetMapping("/isUsersCompany/{id}")
    public ResponseEntity<IsUsersCompanyDto> isUsersCompany(@PathVariable Long id, HttpServletRequest request){
        String username = tokenUtils.getUsernameFromToken(tokenUtils.getToken(request));
        User user = userService.findByUsername(username);
        Company company = companyService.getById(id);
        IsUsersCompanyDto isUsersCompanyDto = new IsUsersCompanyDto();
        if(company.getOwner().getUsername().equals(user.getUsername())){
            isUsersCompanyDto.setMessage("TRUE");
        }else {
            isUsersCompanyDto.setMessage("FALSE");
        }
        return new ResponseEntity<IsUsersCompanyDto>(isUsersCompanyDto, HttpStatus.OK);
    }

}
