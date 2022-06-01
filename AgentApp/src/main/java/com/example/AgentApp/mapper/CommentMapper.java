package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.CompanyRepository;
import com.example.AgentApp.repository.UserRepository;
import com.example.AgentApp.service.*;
import org.springframework.stereotype.Component;

import java.util.*;

@Component
public class CommentMapper {

    private final UserRepository userRepository;
    private final CompanyService companyService;

    public CommentMapper(UserRepository userRepository, CompanyService companyService) {
        this.userRepository = userRepository;
        this.companyService = companyService;
    }

    public Comment toEntity(CommentDto dto){
        Comment comment = new Comment();
        User u = userRepository.findByUsername(dto.getUserUsername());
        Company company = companyService.getById(dto.companyId);
        comment.setCompany(company);
        comment.setUser(u);
        comment.setComment(dto.getComment());
        return comment;
    }

    public List<CommentResponseDto> mapToDtos(Set<Comment> allCommentsForCompany) {
        List<CommentResponseDto> commentDtos = new ArrayList<CommentResponseDto>();
        for (Comment comment: allCommentsForCompany) {
            commentDtos.add(mapToDto(comment));
        }
        return commentDtos;
    }

    private CommentResponseDto mapToDto(Comment comment) {
        CommentResponseDto commentResponseDto = new CommentResponseDto();
        commentResponseDto.comment = comment.getComment();
        commentResponseDto.userUsername = comment.getUser().getUsername();
        commentResponseDto.companyId = comment.getCompany().getId();
        return commentResponseDto;
    }
}
