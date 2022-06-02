package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.enums.*;
import com.example.AgentApp.model.*;
import com.example.AgentApp.service.*;
import org.springframework.beans.factory.annotation.*;
import org.springframework.stereotype.*;

import java.util.*;

@Component
public class InterviewMapper {

    @Autowired
    private UserService userService;
    @Autowired
    private CompanyService companyService;

    public Interview mapToEntity(InterviewRequestDto commentDto) {
        Interview interview = new Interview();
        User user = userService.findByUsername(commentDto.userUsername);
        Company company = companyService.getById(commentDto.companyID);
        interview.setComment(commentDto.comment);
        interview.setRating(commentDto.rating);
        if (commentDto.difficulty.equals("HARD")) interview.setDifficulty(InterviewDifficulty.HARD);
        else if (commentDto.difficulty.equals("INTERMEDIATE")) interview.setDifficulty(InterviewDifficulty.INTERMEDIATE);
        else interview.setDifficulty(InterviewDifficulty.EASY);
        interview.setUser(user);
        interview.setCompany(company);
        return interview;
    }

    public List<InterviewResponseDto> mapToDtos(Set<Interview> allInterviewsCompany) {
        List<InterviewResponseDto> interviewDtos = new ArrayList<InterviewResponseDto>();
        for (Interview interview: allInterviewsCompany) {
            interviewDtos.add(mapToDto(interview));
        }
        return interviewDtos;
    }

    private InterviewResponseDto mapToDto(Interview interview) {
        InterviewResponseDto interviewResponseDto = new InterviewResponseDto();
        interviewResponseDto.comment = interview.getComment();
        interviewResponseDto.difficulty = interview.getDifficulty().toString();
        interviewResponseDto.rating = interview.getRating();
        interviewResponseDto.companyID = interview.getCompany().getId();
        interviewResponseDto.userUsername = interview.getUser().getUsername();
        return interviewResponseDto;
    }
}
