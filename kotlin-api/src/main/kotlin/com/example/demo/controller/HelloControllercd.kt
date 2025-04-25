package com.example.demo.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class HelloControllercd {
    @GetMapping("/hello")
    suspend fun hello() = mapOf("message" to "hello")
}