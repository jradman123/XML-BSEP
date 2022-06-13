package com.example.AgentApp.dto;

import java.time.LocalDate;
import java.util.*;

public class CreateJobOfferRequestDto {
    public Long companyId;
    public Set<String> requirements;
    public String position;
    public String jobDescription;
    public String dueDate;
}
