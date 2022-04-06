package com.example.PKI.repository;

import com.example.PKI.model.*;
import org.springframework.data.jpa.repository.support.*;

import java.util.*;

public interface CertificateRepository extends JpaRepositoryImplementation<Certificate, Integer> {
    Collection<Certificate> findAllByTypeAndValid(CertificateType ct, boolean b);
    Certificate findBySerialNumber(String serial);
}
