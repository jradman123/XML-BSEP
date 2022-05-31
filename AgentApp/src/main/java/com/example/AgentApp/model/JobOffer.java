package com.example.AgentApp.model;

import lombok.Data;
import lombok.ToString;

import javax.persistence.*;
import java.util.List;

@Entity
@Data
public class JobOffer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;

    @Column
    @ElementCollection(targetClass = String.class)
    private List<String> requirements;

    @Column
    @ElementCollection(targetClass = String.class)
    private List<String> otherRequirements;



}
