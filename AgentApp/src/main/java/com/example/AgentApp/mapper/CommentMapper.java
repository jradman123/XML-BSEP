package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.CommentDto;
import com.example.AgentApp.model.Comment;
import com.example.AgentApp.model.User;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.UserRepository;
import org.springframework.stereotype.Component;

import java.util.Optional;

@Component
public class CommentMapper {

    private final UserRepository userRepository;
    private final CompanyRepository companyRepository;

    public CommentMapper(UserRepository userRepository, CompanyRepository companyRepository) {
        this.userRepository = userRepository;
        this.companyRepository = companyRepository;
    }

    public Comment toEntity(CommentDto dto){
        Comment comment = new Comment();
        User u = userRepository.findByUsername(dto.getUserUsername());
        comment.setUser(u);
        comment.setComment(dto.getComment());
        return comment;
    }

}
