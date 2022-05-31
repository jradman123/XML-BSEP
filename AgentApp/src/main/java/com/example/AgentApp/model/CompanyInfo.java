package com.example.AgentApp.model;

import lombok.*;

import javax.persistence.*;

@Data
@Entity
public class CompanyInfo {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @Column
    private String name;
    @Column
    private String website;
    @Column
    private String headquarters;
    @Column
    private String industry;
    @Column
    private String founded;
    @Column
    private String noOfEmpl;
    @Column
    private String countryOfOrigin;
    @Column
    private String offices;

}
