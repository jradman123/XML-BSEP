package com.example.PKI.repository;

import com.example.PKI.model.*;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.jpa.repository.support.*;

import java.util.*;

public interface CertificateRepository extends JpaRepositoryImplementation<Certificate, Integer> {
    Collection<Certificate> findAllByTypeAndIsRevoked(CertificateType ct, boolean b);

    Certificate findBySerialNumber(String serial);

    @Query(value = "select c.type from certificate c where c.serial_number = ?1", nativeQuery = true)
    CertificateType findTypeBySerialNumber(String serialNumber);


    @Query(value = "select * from certificate c where c.type = 1 and c.subject_email = ?1 and c.valid_from < ?2 and c.valid_to > ?3 ", nativeQuery = true)
    Certificate[] findAllSignersCertByUser(String email,String startDate,String endDate);

    @Query(value = "select * from certificate c where c.type  = 0 or c.type = 1 and c.valid_from < ?1 and c.valid_to > ?2", nativeQuery = true)
    Collection<Certificate> findCertificatesValidForDateRange(String startDate, String endDate);

}
