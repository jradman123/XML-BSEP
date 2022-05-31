package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

@Component
public class CompanyMapper {

    @Autowired
    private UserRepository userRepository;

    public Company mapToCompany(NewCompanyRequestDto dto){
        User owner = userRepository.findByUsername(dto.ownerUsername);
        Company company = new Company();
        company.setCompanyStatus(CompanyStatus.PENDING);
        company.setContactInfo(dto.contactInfo);
        company.setCompanyPolicy(dto.companyPolicy);
        company.setOwner(owner);
        return company;
    }

    public NewCompanyResponseDto mapToCompanyCreateResponse(Company company){
        NewCompanyResponseDto response = new NewCompanyResponseDto();
        response.companyPolicy = company.getCompanyPolicy();
        response.contactInfo = company.getContactInfo();
        response.companyStatus = company.getCompanyStatus();
        response.message = "Request for creating company is created!";
        return response;
    }
}
