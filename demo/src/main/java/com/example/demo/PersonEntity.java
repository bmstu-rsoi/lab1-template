package com.example.demo;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class PersonEntity {

    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    public long id;
    public String name;
    public int age;
    public String address;
    public String work;

    public PersonEntity(){

    };

    public PersonEntity(String name, int age, String address, String work) {
        this.name = name;
        this.age = age;
        this.address = address;
        this.work = work;
    }
}

