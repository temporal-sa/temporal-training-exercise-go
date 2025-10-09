package com.temporal.training.solution1;

public class GreetingActivityImpl implements GreetingActivity {
    @Override
    public String createGreeting(String name) {
        return "Hello, " + name + "!";
    }
}
