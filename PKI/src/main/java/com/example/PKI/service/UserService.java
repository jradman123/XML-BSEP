package com.example.PKI.service;
import com.example.PKI.dto.LoginDTO;
import com.example.PKI.model.User;

public interface UserService {
	boolean login(LoginDTO loginDTO);
	void logout(String email);
	User findByEmail(String email);

}
