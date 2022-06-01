package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public class CompanyServiceImpl implements CompanyService {

    @Autowired
    private CompanyRepository companyRepository;
    @Autowired
    private UserRepository userRepository;
    @Autowired
    private JobOfferRepository jobOfferRepository;
    @Autowired
    private CompanyMapper companyMapper;

    @Override
    public Company createCompany(NewCompanyRequestDto companyDto) {
        Company company = companyMapper.mapToCompany(companyDto);
        companyRepository.save(company);
        System.out.println("sacuvana kompanija pod id-jem:" + company.getId());
        return company;
    }

    @Override
    public Company approveCompany(Long id, boolean approve) {
        if (approve) {
            Optional<Company> company = companyRepository.findById(id);
            company.get().setCompanyStatus(CompanyStatus.APPROVED);
            User owner = company.get().getOwner();
            owner.setRole(UserRole.OWNER);
            company.get().setOwner(owner);
            companyRepository.save(company.get());
            userRepository.save(owner);
            return company.get();
//            Set<JobOffer> offers = (Set<JobOffer>) company.get().getJobOffers();
//            Set<Interview> interviews = (Set<Interview>) company.get().getInterviews(); //ovo nece baca error
        }else {
            Optional<Company> company = companyRepository.findById(id);
            company.get().setCompanyStatus(CompanyStatus.DECLINED);
            companyRepository.save(company.get());
            return company.get();
        }
    }

    @Override
    public Company editCompany(EditCompanyRequestDto requestDto, Long id) {
        Optional<Company> company = companyRepository.findById(id);
        Company editedCompany = companyMapper.editCompany(requestDto,company.get());
        companyRepository.save(editedCompany);
        return editedCompany;
    }

    @Override
    public Company addJobOffer(CreateJobOfferRequestDto requestDto) {
        Optional<Company> company = companyRepository.findById(requestDto.companyId);
        JobOffer jobOffer = new JobOffer();
        jobOffer.setCompany(company.get());
        jobOffer.setRequirements(requestDto.requirements);
        jobOffer.setOtherRequirements(requestDto.otherRequirements);
        jobOfferRepository.save(jobOffer);
        Optional<Company> companyWithOffer = companyRepository.findById(requestDto.companyId);
        return companyWithOffer.get();
    }

    @Override
    public List<Company> getAllCompaniesWithStatus(CompanyStatus status) {
        List<Company> companies = companyRepository.findAllByCompanyStatus(status);
        return companies;
    }
    public Company getById(Long id) {
        return companyRepository.findById(id).orElseThrow(() -> new RuntimeException("There is no company with id " + id));
    }

}
