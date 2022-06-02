package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.repository.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Component
public class SalaryCommentMapper {

    @Autowired
    private UserService userService;
    @Autowired
    private CompanyService companyService;

    public SalaryComment mapToEntity(SalaryCommentRequestDto commentDto) {
        SalaryComment salaryComment = new SalaryComment();
        User user = userService.findByUsername(commentDto.userUsername);
        Company company = companyService.getById(commentDto.companyID);
        salaryComment.setSalary(commentDto.salary);
        salaryComment.setPosition(commentDto.position);
        salaryComment.setUser(user);
        salaryComment.setCompany(company);
        return  salaryComment;
    }

    public List<SalaryCommentResponseDto> mapToDtos(Set<SalaryComment> allCommentsForCompany) {
        List<SalaryCommentResponseDto> commentDtos = new ArrayList<SalaryCommentResponseDto>();
        for (SalaryComment comment: allCommentsForCompany) {
            commentDtos.add(mapToDto(comment));
        }
        return commentDtos;
    }

    private SalaryCommentResponseDto mapToDto(SalaryComment comment) {
        SalaryCommentResponseDto salaryCommentResponseDto = new SalaryCommentResponseDto();
        salaryCommentResponseDto.companyId = comment.getCompany().getId();
        salaryCommentResponseDto.userUsername = comment.getUser().getUsername();
        salaryCommentResponseDto.salary = comment.getSalary();
        salaryCommentResponseDto.position = comment.getPosition();
        return salaryCommentResponseDto;
    }

}
