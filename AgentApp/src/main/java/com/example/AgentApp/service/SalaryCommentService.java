package com.example.AgentApp.service;

import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public interface SalaryCommentService {
    SalaryComment create(SalaryComment comment);

    Set<SalaryComment> getAllForCompany(Long companyID);
}
