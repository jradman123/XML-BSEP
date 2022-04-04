package com.example.PKI.repository;

import org.springframework.stereotype.Repository;
import org.springframework.data.jpa.repository.JpaRepository;
import com.example.PKI.model.User;

@Repository
public interface UserRepository extends JpaRepository<User,Integer>{
	User findByEmail(String email);
}
