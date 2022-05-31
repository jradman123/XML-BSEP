package com.example.AgentApp.service.impl;

import com.example.AgentApp.model.Comment;
import com.example.AgentApp.repository.CommentRepository;
import com.example.AgentApp.service.CommentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

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
}
