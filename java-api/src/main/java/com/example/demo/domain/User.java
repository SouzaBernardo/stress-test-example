package com.example.demo.domain;

import org.springframework.data.annotation.Id;
import org.springframework.data.relational.core.mapping.Table;

@Table("users")
public class User {
    @Id
    private Integer id;
    private String name;
    private UserStatus status;

    public User(String name, UserStatus status) {
        this.name = name;
        this.status = status;
    }

    public String getName() {
        return name;
    }

    public UserStatus getStatus() {
        return status;
    }
}
