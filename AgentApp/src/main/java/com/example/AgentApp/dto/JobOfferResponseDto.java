package com.example.AgentApp.dto;

import java.time.LocalDate;
import java.util.*;

public class JobOfferResponseDto {
    public Long offerId;
    public Set<String> requirements;
    public String position;
    public String jobDescription;
    public String dateCreated;
    public String dueDate;
}
