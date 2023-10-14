package com.rsoi.lab1personservice.dto.response;

public class PersonResponse {
    private String name;
    private String work;
    private Integer age;
    private String address;
    private Integer id;

    public PersonResponse(Integer id, String work, Integer age, String address, String name) {
        this.id = id;
        this.work = work;
        this.age = age;
        this.address = address;
        this.name = name;
    }

    public PersonResponse() {
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

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }
}
