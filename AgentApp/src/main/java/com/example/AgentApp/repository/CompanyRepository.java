package com.example.AgentApp.repository;

import com.example.AgentApp.dto.JobOfferResponseDto;
import com.example.AgentApp.enums.CompanyStatus;
import com.example.AgentApp.model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface CompanyRepository extends JpaRepository<Company, Long> {
    Optional<Company> findById(Long id);

    List<Company> findAllByCompanyStatus(CompanyStatus status);

    @Query(value = "select * from company c where c.owner_id != ?1 and c.company_status = 0", nativeQuery = true)
    List<Company> findAllExceptOwnerId(Long id);

    @Query(value = "select * from company c where c.owner_id = ?1", nativeQuery = true)
    List<Company> getAllUsersCompanies(Long userId);

    @Query(value = "select * from company c where c.id = ?1", nativeQuery = true)
    List<JobOfferResponseDto> getAllJobOffers(Long companyId);
}
