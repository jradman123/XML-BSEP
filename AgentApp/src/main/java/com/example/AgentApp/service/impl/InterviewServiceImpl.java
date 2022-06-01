package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public class InterviewServiceImpl implements InterviewService {
    @Autowired
    private InterviewRepository interviewRepository;

    @Override
    public Interview create(Interview interview) {
        return interviewRepository.save(interview);
    }

    @Override
    public Set<Interview> getAllForCompany(Long companyID) {
        return interviewRepository.findAllByCompanyId(companyID);
    }
}
