package com.example.AgentApp.repository;

import com.example.AgentApp.model.*;
import org.springframework.data.jpa.repository.*;
import org.springframework.stereotype.*;

@Repository
public interface CompanyRepository extends JpaRepository<Company,Long> {
}
