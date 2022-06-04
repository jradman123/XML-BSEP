package com.example.AgentApp.model;

import lombok.*;

import javax.persistence.*;

@Data
@Entity
public class CompanyInfo {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private String website;
    private String headquarters;
    private String industry;
    private String founded;
    private String noOfEmpl;
    private String countryOfOrigin;
    private String offices;

}
