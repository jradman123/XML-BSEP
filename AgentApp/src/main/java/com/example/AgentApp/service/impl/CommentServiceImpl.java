package com.example.AgentApp.service.impl;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.Comment;
import com.example.AgentApp.repository.CommentRepository;
import com.example.AgentApp.service.CommentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class CommentServiceImpl implements CommentService {

    private final CommentRepository commentRepository;

    public CommentServiceImpl(CommentRepository commentRepository) {
        this.commentRepository = commentRepository;
    }

    @Override
    public Comment create(Comment newComment) {
        return commentRepository.save(newComment);
    }

    @Override
    public Set<Comment> getAllForCompany(Long companyId) {
        return commentRepository.findAllByCompanyId(companyId);
    }
}
