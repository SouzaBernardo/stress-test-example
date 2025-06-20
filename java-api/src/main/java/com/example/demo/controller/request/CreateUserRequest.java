package com.example.demo.controller.request;

import com.example.demo.domain.UserStatus;

public record CreateUserRequest(
        String name,
        UserStatus status
) {
}
