package com.example.AgentApp.model;

import lombok.Data;

import javax.persistence.*;

@Entity
@Data
public class SalaryComment {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "user_id")
    private User user;

    @ManyToOne
    @JoinColumn(name = "company_id")
    private Company company;

    @Column(nullable = false)
    private String position;

    @Column(nullable = false)
    private String salary;
}

