package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.util.*;

@Component
public class JobOfferMapper {
    public List<JobOfferResponseDto> mapToDtos(List<JobOffer> offers) {
        List<JobOfferResponseDto> jobOfferResponseDtos = new ArrayList<JobOfferResponseDto>();
        for (JobOffer o: offers) {
            jobOfferResponseDtos.add(mapToDto(o));
        }
        return jobOfferResponseDtos;
    }

    public JobOfferResponseDto mapToDto(JobOffer offer){
        JobOfferResponseDto jobOfferResponseDto = new JobOfferResponseDto();
        jobOfferResponseDto.offerId = offer.getId();
        jobOfferResponseDto.requirements = offer.getRequirements();
        jobOfferResponseDto.otherRequirements = offer.getOtherRequirements();
        return jobOfferResponseDto;
    }
}
