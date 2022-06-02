package com.example.AgentApp.repository;

import com.example.AgentApp.model.*;
import org.springframework.data.jpa.repository.*;
import org.springframework.stereotype.*;

import java.util.*;

@Repository
public interface InterviewRepository extends JpaRepository<Interview, Long> {
    Set<Interview> findAllByCompanyId(Long companyID);
}
