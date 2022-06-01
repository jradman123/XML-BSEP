package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Component
public class CompanyMapper {

    @Autowired
    private UserRepository userRepository;

    public Company mapToCompany(NewCompanyRequestDto dto){
        User owner = userRepository.findByUsername(dto.ownerUsername);
        Company company = new Company();
        company.setCompanyStatus(CompanyStatus.PENDING);
        company.setCompanyInfo(new CompanyInfo());
        company.getCompanyInfo().setCountryOfOrigin(dto.countryOfOrigin);
        company.getCompanyInfo().setFounded(dto.founded);
        company.getCompanyInfo().setHeadquarters(dto.headquarters);
        company.getCompanyInfo().setIndustry(dto.industry);
        company.getCompanyInfo().setName(dto.companyName);
        company.getCompanyInfo().setOffices(dto.offices);
        company.getCompanyInfo().setWebsite(dto.companyWebsite);
        company.getCompanyInfo().setNoOfEmpl(dto.noOfEmpl);
        company.setCompanyPolicy(dto.companyPolicy);
        company.setOwner(owner);
        return company;
    }

    public NewCompanyResponseDto mapToCompanyCreateResponse(Company company){
        NewCompanyResponseDto response = new NewCompanyResponseDto();
        response.companyPolicy = company.getCompanyPolicy();
       // response.contactInfo = company.getContactInfo();
        response.companyStatus = company.getCompanyStatus();
        response.message = "Request for creating company is created!";
        return response;
    }

    public Company editCompany(EditCompanyRequestDto requestDto, Company company) {
        company.setCompanyPolicy(requestDto.companyPolicy);
        company.getCompanyInfo().setName(requestDto.companyName);
        company.getCompanyInfo().setWebsite(requestDto.companyWebsite);
        company.getCompanyInfo().setHeadquarters(requestDto.headquarters);
        company.getCompanyInfo().setIndustry(requestDto.industry);
        company.getCompanyInfo().setFounded(requestDto.founded);
        company.getCompanyInfo().setNoOfEmpl(requestDto.noOfEmpl);
        company.getCompanyInfo().setCountryOfOrigin(requestDto.countryOfOrigin);
        company.getCompanyInfo().setOffices(requestDto.offices);
        return company;
    }

    public CompanyResponseDto mapToDto(Company company){
        CompanyResponseDto companyResponseDto = new CompanyResponseDto();
        companyResponseDto.companyId = company.getId();
        companyResponseDto.companyPolicy = company.getCompanyPolicy();
        companyResponseDto.companyName = company.getCompanyInfo().getName();
        companyResponseDto.companyWebsite = company.getCompanyInfo().getWebsite();
        companyResponseDto.headquarters = company.getCompanyInfo().getHeadquarters();
        companyResponseDto.industry = company.getCompanyInfo().getIndustry();
        companyResponseDto.founded = company.getCompanyInfo().getFounded();
        companyResponseDto.noOfEmpl = company.getCompanyInfo().getNoOfEmpl();
        companyResponseDto.countryOfOrigin = company.getCompanyInfo().getCountryOfOrigin();
        companyResponseDto.offices = company.getCompanyInfo().getOffices();
        return companyResponseDto;
    }

    public List<CompanyResponseDto> mapToDtos(List<Company> companies){
        List<CompanyResponseDto> companyDtos = new ArrayList<CompanyResponseDto>();
        for (Company company: companies) {
            companyDtos.add(mapToDto(company));
        }
        return companyDtos;
    }
}
