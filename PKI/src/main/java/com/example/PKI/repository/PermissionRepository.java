package com.example.PKI.repository;

import com.example.PKI.model.Permission;
import org.springframework.data.jpa.repository.JpaRepository;

public interface PermissionRepository extends JpaRepository<Permission, Long> {
    public Permission findByName(String name);
}
