package com.example.AgentApp.service;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public interface CompanyService {
    Company createCompany(NewCompanyRequestDto companyDto);

    Company approveCompany(Long id,boolean approve);

    Company editCompany(EditCompanyRequestDto requestDto, Long id);

    Company addJobOffer(CreateJobOfferRequestDto requestDto);

    List<Company> getAllCompaniesWithStatus(CompanyStatus status);
    Company getById(Long id);

    List<Company> getAllApprovedCompaniesExceptOwners(User user);

    List<Company> getAllUsersCompanies(Long userId);

    List<JobOfferResponseDto> getAllJobOffers(Long companyId);
}
