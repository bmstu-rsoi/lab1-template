package com.rsoi.lab1personservice.repo;

import com.rsoi.lab1personservice.dao.Person;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;


@Repository
public interface PersonRepo extends JpaRepository<Person, Integer> {
}
