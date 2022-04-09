package com.example.PKI.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.example.PKI.dto.LoginDto;
import com.example.PKI.model.User;
import com.example.PKI.repository.UserRepository;
import com.example.PKI.service.UserService;

@Service
public class UserServiceImpl implements UserService{

	private UserRepository userRepository;
	
	@Autowired
	public UserServiceImpl(UserRepository userRepository) {
		this.userRepository = userRepository;
	}
	
	@Override
	public boolean login(LoginDto user) {
		User foundUser = userRepository.findByEmail(user.getEmail());
		if(foundUser.getPassword().equals(user.getPassword())){
			foundUser.setLoggedIn(true);
			userRepository.save(foundUser);
			return true;
		}
		return false;
	}

	@Override
	public void logout(String email) {
		User foundUser = userRepository.findByEmail(email);
		foundUser.setLoggedIn(false);
		userRepository.save(foundUser);
	}

	@Override
	public User findByEmail(String email) {
		return userRepository.findByEmail(email);
	}
	
	

}
