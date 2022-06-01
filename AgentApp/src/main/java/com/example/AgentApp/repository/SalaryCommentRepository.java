package com.example.AgentApp.repository;

import com.example.AgentApp.model.*;
import org.springframework.data.jpa.repository.*;

import java.util.*;

public interface SalaryCommentRepository extends JpaRepository<SalaryComment, Long> {
    Set<SalaryComment> findAllByCompanyId(Long companyID);
}
