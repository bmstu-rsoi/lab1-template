package com.rsoi.lab1personservice.dao;

import jakarta.persistence.*;

@Entity
@Table(name = "person")

public class Person {
    @Id
    @Column(name="id")
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Integer id;

    @Column(name = "name")
    private String name;

    @Column(name = "work")
    private String work;

    @Column(name = "age")
    private Integer age;

    @Column(name = "address")
    private String address;


    public Person() {
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
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

    public Person(String name, String work, Integer age, String address) {
        this.name = name;
        this.work = work;
        this.age = age;
        this.address = address;
    }
}
