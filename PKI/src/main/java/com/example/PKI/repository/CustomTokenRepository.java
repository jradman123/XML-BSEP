package com.example.PKI.repository;

import com.example.PKI.model.CustomToken;
import com.example.PKI.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CustomTokenRepository extends JpaRepository<CustomToken, Long> {
    CustomToken findByToken(String confirmationToken);
    CustomToken findByUser(User user);
}
