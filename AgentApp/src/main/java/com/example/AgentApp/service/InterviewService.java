package com.example.AgentApp.service;

import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.util.*;

@Service
public interface InterviewService {
    Interview create(Interview interview);

    Set<Interview> getAllForCompany(Long companyID);
}
