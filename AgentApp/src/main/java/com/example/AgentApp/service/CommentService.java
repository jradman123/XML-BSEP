package com.example.AgentApp.service;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.Comment;

import java.util.*;

public interface CommentService {
    Comment create(Comment newComment);

    Set<Comment> getAllForCompany(Long companyId);
}
