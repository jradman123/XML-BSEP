package com.example.AgentApp.repository;

import com.example.AgentApp.enums.*;
import com.example.AgentApp.model.*;
import org.springframework.data.jpa.repository.*;
import org.springframework.stereotype.*;

import java.util.*;

@Repository
public interface CompanyRepository extends JpaRepository<Company,Long> {
    Optional<Company> findById(Long id);
    List<Company> findAllByCompanyStatus(CompanyStatus status);
}
