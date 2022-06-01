package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public class SalaryCommentServiceImpl implements SalaryCommentService {

    @Autowired
    private SalaryCommentRepository salaryCommentRepository;

    @Override
    public SalaryComment create(SalaryComment comment) {
        return salaryCommentRepository.save(comment);
    }

    @Override
    public Set<SalaryComment> getAllForCompany(Long companyID) {
        return salaryCommentRepository.findAllByCompanyId(companyID);
    }
}
