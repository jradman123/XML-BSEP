package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.mapper.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

@Service
public class CompanyServiceImpl implements CompanyService {

    @Autowired
    private CompanyRepository companyRepository;
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
    public Company getById(Long id) {
        return companyRepository.findById(id).orElseThrow(() -> new RuntimeException("There is no company with id " + id));
    }

}
