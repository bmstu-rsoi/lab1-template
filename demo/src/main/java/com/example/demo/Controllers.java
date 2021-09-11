package com.example.demo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletResponse;
import java.util.List;
import java.util.Optional;

@Controller
@EnableAutoConfiguration
public class Controllers {

    @Autowired
    PersonRepository personRepository;

    @PostMapping(value = "/api/v1/persons")
    @ResponseStatus(HttpStatus.CREATED)
    @ResponseBody
    @ModelAttribute
    public ResponseEntity<PostException> post(@RequestBody PersonEntity pe, HttpServletResponse response) {

        PostException exception = new PostException();
        if(pe.work==null) exception.errors.put("work", "don't be null");
        if(pe.age <= 0) exception.errors.put("age", "don't be equal or lower than 0");
        if(pe.address==null) exception.errors.put("address", "don't be null");
        if(pe.name==null) exception.errors.put("name", "don't be null");
        if(exception.errors.size()>0){
            exception.message = "wrong fields";
            return ResponseEntity.status(400).body(exception);
        }

        try {
            PersonEntity personEntity = personRepository.save(pe);
            response.setHeader("Location","/api/v1/persons/"+personEntity.id);
        } catch (Exception e){
            exception.message = e.getMessage();
            return ResponseEntity.status(400).body(exception);
        }

        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @GetMapping("/api/v1/persons/{id}")
    @ResponseBody
    public Optional<PersonEntity> get(@PathVariable Long id, HttpServletResponse response) {
        Optional<PersonEntity> personEntity = personRepository.findById(id);
        if(personEntity.isEmpty()) response.setStatus(404);
        return personEntity;
    }

    @DeleteMapping("/api/v1/persons/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    @ResponseBody
    public void delete(@PathVariable Long id) {
        personRepository.deleteById(id);
    }

    @PatchMapping("/api/v1/persons/{id}")
    @ResponseStatus(HttpStatus.OK)
    @ResponseBody
    public ResponseEntity patch(@RequestBody PersonEntity pe, @PathVariable Long id, HttpServletResponse response) {
        Optional<PersonEntity> personEntity = personRepository.findById(id);
        if(personEntity.isEmpty()) {
            response.setStatus(404);
            return null;
        }

        if(pe.name != null) personEntity.get().name = pe.name;
        if(pe.address != null) personEntity.get().address = pe.address;
        if(pe.age != 0) personEntity.get().age = pe.age;
        if(pe.work != null) personEntity.get().work = pe.work;
        PersonEntity r = personRepository.save(personEntity.get());

        return ResponseEntity.ok(r);
    }

    @GetMapping("/api/v1/persons")
    @ResponseBody
    public List<PersonEntity> simple3() {
        return (List<PersonEntity>) personRepository.findAll();
    }

}