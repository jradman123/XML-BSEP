package com.example.AgentApp.dto;

import java.util.*;

public class CreateJobOfferRequestDto {
    public Long companyId;
    public String name;
    public Set<String> requirements;

    public String position;
    public String jobDescription;

}
