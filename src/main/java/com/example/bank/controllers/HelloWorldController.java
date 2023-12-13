package com.example.bank.controllers;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/hello")
public class HelloWorldController {

    @GetMapping
    public ResponseEntity<HelloWorldResponse> getHelloWorld() {
        final HelloWorldResponse response = new HelloWorldResponse("Hello World!");
        return ResponseEntity.ok(response);
    }
}
