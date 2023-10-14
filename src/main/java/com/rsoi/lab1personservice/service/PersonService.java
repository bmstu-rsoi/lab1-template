package com.rsoi.lab1personservice.service;

import com.rsoi.lab1personservice.dao.Person;
import com.rsoi.lab1personservice.dto.request.PersonRequest;
import com.rsoi.lab1personservice.dto.response.PersonResponse;
import com.rsoi.lab1personservice.repo.PersonRepo;
import org.hibernate.ObjectNotFoundException;
import org.springframework.stereotype.Service;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;


@Service
public class PersonService {

    private final PersonRepo personRepo;

    public PersonService(PersonRepo personRepo) {
        this.personRepo = personRepo;
    }

    public PersonResponse getPersonById(Integer id) {
        Optional<Person> person = personRepo.findById(id);
        return person.isEmpty() ? null : new PersonResponse(
                person.get().getId(),
                person.get().getWork(),
                person.get().getAge(),
                person.get().getAddress(),
                person.get().getName());
    }

    public List<PersonResponse> getPersons() {
        List<Person> persons = personRepo.findAll();
        return persons
                .stream()
                .map(person -> new PersonResponse(
                        person.getId(),
                        person.getWork(),
                        person.getAge(),
                        person.getAddress(),
                        person.getName()))
                .collect(Collectors.toList());
    }

    public Integer createPerson(PersonRequest personRequest) {
        Person personSave = new Person(
                personRequest.getName(),
                personRequest.getWork(),
                personRequest.getAge(),
                personRequest.getAddress()
        );
        return personRepo.save(personSave).getId();
    }

    public Person updatePerson(Integer id, PersonRequest personDto) {
        Person person = personRepo
                .findById(id)
                .orElseThrow(
                        () -> new ObjectNotFoundException((Object) "Person not founded!", "Person with id = " + id + " not founded!")
                );

        if (personDto.getName() != null) {
            person.setName(personDto.getName());
        }

        if (personDto.getAge() != null) {
            person.setAge(personDto.getAge());
        }

        if (personDto.getAddress() != null) {
            person.setAddress(personDto.getAddress());
        }

        if (personDto.getWork() != null) {
            person.setWork(personDto.getWork());
        }

        personRepo.save(person);

        return person;
    }

    public void deletePerson(Integer id) {
        personRepo.deleteById(id);
    }
}
