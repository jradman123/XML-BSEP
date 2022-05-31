package com.example.AgentApp.service;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

@Service
public interface CompanyService {
    Company createCompany(NewCompanyRequestDto companyDto);
}
