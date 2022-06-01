package com.example.AgentApp.repository;

import com.example.AgentApp.model.*;
import org.springframework.data.jpa.repository.*;
import org.springframework.stereotype.*;

import java.util.*;

@Repository
public interface JobOfferRepository extends JpaRepository<JobOffer, Long> {
    
    List<JobOffer> findAllByCompanyId(Long companyId);
}
