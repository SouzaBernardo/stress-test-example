package com.example.demo.service;

import com.example.demo.controller.request.CreateUserRequest;
import com.example.demo.controller.response.UserResponse;
import com.example.demo.domain.User;
import com.example.demo.repository.UserRepository;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.function.Function;

@Service
public class UserService {

    private final UserRepository userRepository;

    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public Mono<UserResponse> create(CreateUserRequest request) {
        User entity = new User(request.name(), request.status());
        return userRepository.save(entity)
                .flatMap(mapper());
    }

    public Flux<UserResponse> list() {
        return userRepository.findAll()
                .flatMap(mapper());
    }

    public Mono<Long> count() {
        return userRepository.count();
    }

    private static Function<User, Mono<? extends UserResponse>> mapper() {
        return user -> {
            UserResponse response = new UserResponse(user.getName(), user.getStatus().toString());
            return Mono.just(response);
        };
    }
}
