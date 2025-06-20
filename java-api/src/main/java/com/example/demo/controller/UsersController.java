package com.example.demo.controller;

import com.example.demo.controller.request.CreateUserRequest;
import com.example.demo.controller.response.UserResponse;
import com.example.demo.service.UserService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/users")
public class UsersController {

    private final UserService userService;

    public UsersController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping
    public Mono<UserResponse> create(@RequestBody CreateUserRequest request) {
        return userService.create(request);
    }

    @GetMapping
    public Flux<UserResponse> findAll() {
        return userService.list();
    }

    @GetMapping("/count")
    public Mono<Long> count() {
        return userService.count();
    }

}
