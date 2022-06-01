package com.example.AgentApp.repository;

import com.example.AgentApp.model.Comment;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.*;

public interface CommentRepository extends JpaRepository<Comment, Long> {
    Set<Comment> findAllByCompanyId(Long companyId);
}
