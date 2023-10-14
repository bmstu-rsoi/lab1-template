package com.rsoi.lab1personservice.dto.request;

import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;

public class PersonRequest {
    @NotBlank
    private String name;
    private String work;
    @Min(value = 1)
    private Integer age;
    private String address;

    public PersonRequest(String name, String work, Integer age, String address) {
        this.name = name;
        this.work = work;
        this.age = age;
        this.address = address;
    }

    public PersonRequest() {
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getWork() {
        return work;
    }

    public void setWork(String work) {
        this.work = work;
    }

    public Integer getAge() {
        return age;
    }

    public void setAge(Integer age) {
        this.age = age;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }
}
