package com.example.PKI.repository;

import com.example.PKI.model.CustomToken;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CustomTokenRepository extends JpaRepository<CustomToken, Long> {
    CustomToken findByToken(String confirmationToken);
}
