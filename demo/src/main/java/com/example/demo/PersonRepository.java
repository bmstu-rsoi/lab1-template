package com.example.demo;

import org.springframework.data.repository.CrudRepository;

import java.util.List;

public interface PersonRepository extends CrudRepository<PersonEntity, Long> {

    List<PersonEntity> findByName(String lastName);

}
