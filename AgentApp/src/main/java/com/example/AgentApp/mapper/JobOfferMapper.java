package com.example.AgentApp.mapper;

import com.example.AgentApp.dto.*;
import com.example.AgentApp.model.*;
import org.springframework.stereotype.*;

import java.text.SimpleDateFormat;
import java.util.*;

@Component
public class JobOfferMapper {
    public static List<JobOfferResponseDto> mapToDtos(Set<JobOffer> offers) {
        List<JobOfferResponseDto> jobOfferResponseDtos = new ArrayList<JobOfferResponseDto>();
        for (JobOffer o: offers) {
            jobOfferResponseDtos.add(mapToDto(o));
        }
        return jobOfferResponseDtos;
    }

    public static JobOfferResponseDto mapToDto(JobOffer offer){
        JobOfferResponseDto jobOfferResponseDto = new JobOfferResponseDto();
        jobOfferResponseDto.offerId = offer.getId();
        jobOfferResponseDto.requirements = offer.getRequirements();
        jobOfferResponseDto.jobDescription = offer.getJobDescription();
        jobOfferResponseDto.position = offer.getPosition();
        jobOfferResponseDto.dateCreated = offer.getDateCreated().toString();
        jobOfferResponseDto.dueDate = offer.getDueDate().toString();
        return jobOfferResponseDto;
    }

    public static List<JobOfferWithCompanyResponseDto> mapToDtosWithCompany(List<JobOffer> offers) {
        List<JobOfferWithCompanyResponseDto> jobOfferResponseDtos = new ArrayList<JobOfferWithCompanyResponseDto>();
        for (JobOffer o: offers) {
            jobOfferResponseDtos.add(mapToDtoWithCompany(o));
        }
        return jobOfferResponseDtos;
    }

    private static JobOfferWithCompanyResponseDto mapToDtoWithCompany(JobOffer offer) {
        JobOfferWithCompanyResponseDto jobOfferResponseDto = new JobOfferWithCompanyResponseDto();
        jobOfferResponseDto.offerId = offer.getId();
        jobOfferResponseDto.position = offer.getPosition();
        jobOfferResponseDto.jobDescription = offer.getJobDescription();
        jobOfferResponseDto.requirements = offer.getRequirements();
        jobOfferResponseDto.companyId = offer.getCompany().getId();
        jobOfferResponseDto.companyPolicy = offer.getCompany().getCompanyPolicy();
        jobOfferResponseDto.companyName = offer.getCompany().getCompanyInfo().getName();
        jobOfferResponseDto.companyWebsite = offer.getCompany().getCompanyInfo().getWebsite();
        jobOfferResponseDto.headquarters = offer.getCompany().getCompanyInfo().getHeadquarters();
        jobOfferResponseDto.industry = offer.getCompany().getCompanyInfo().getIndustry();
        jobOfferResponseDto.founded = offer.getCompany().getCompanyInfo().getFounded();
        jobOfferResponseDto.noOfEmpl = offer.getCompany().getCompanyInfo().getNoOfEmpl();
        jobOfferResponseDto.countryOfOrigin = offer.getCompany().getCompanyInfo().getCountryOfOrigin();
        jobOfferResponseDto.offices = offer.getCompany().getCompanyInfo().getOffices();
        jobOfferResponseDto.dateCreated = offer.getDateCreated().toString();
        jobOfferResponseDto.dueDate = offer.getDueDate().toString();
        return jobOfferResponseDto;
    }
}
