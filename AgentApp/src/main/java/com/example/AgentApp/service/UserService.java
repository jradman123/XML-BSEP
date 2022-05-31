package com.example.AgentApp.service;

import com.example.AgentApp.model.User;

public interface UserService {
    User findByUsername(String username);
}
