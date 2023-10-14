package com.rsoi.lab1personservice.service;

import com.rsoi.lab1personservice.dao.Person;
import com.rsoi.lab1personservice.dto.request.PersonRequest;
import com.rsoi.lab1personservice.dto.response.PersonResponse;
import com.rsoi.lab1personservice.repo.PersonRepo;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import java.time.LocalDateTime;
import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class PersonServiceTest {
    @Mock
    private PersonRepo personRepo;

    @Test
    void getPersonByIdShouldReturnPersonByExistingId() {
        Person person = new Person("name", "work", 21, "address");
        PersonService personService = new PersonService(personRepo);
        when(personRepo.findById(123))
                .thenReturn(java.util.Optional.of(person));
        PersonResponse personResponse1 = personService.getPersonById(123);

        assertAll("Verify person properties",
                () -> assertEquals(personResponse1.getWork(), person.getWork()),
                () -> assertEquals(personResponse1.getName(), person.getName()),
                () -> assertEquals(personResponse1.getAddress(), person.getAddress()),
                () -> assertEquals(personResponse1.getAge(), person.getAge())
        );
        verify(personRepo).findById(123);
    }

    @Test
    void createPersonShouldReturnPersonId() {
        PersonRequest personDto = new PersonRequest("name", "work", 21, "address");
        Person person = new Person("name", "work", 21, "address");
        PersonService personService = new PersonService(personRepo);

        when(personRepo.save(any())).thenReturn(person);

        assertEquals(personService.createPerson(personDto), person.getId());
        verify(personRepo, times(1)).save(any(Person.class));
    }

    @Test
    void updatePerson() {
        LocalDateTime createdDttm = LocalDateTime.now();
        PersonRequest personDto = new PersonRequest("nameNew", "work", 21, "addressNew");
        Person person = new Person( "name", "work", 21, "address");
        Person personUpdated = new Person( "nameNew", "work", 21,"addressNew");
        PersonService personService = new PersonService(personRepo);

        when(personRepo.findById(123)).thenReturn(java.util.Optional.of(person));
        when(personRepo.save(any(Person.class))).thenReturn(any(Person.class));

        Person returnedPerson = personService.updatePerson(123, personDto);
        assertAll("Verify person properties",
                () -> assertEquals(personUpdated.getWork(), returnedPerson.getWork()),
                () -> assertEquals(personUpdated.getName(), returnedPerson.getName()),
                () -> assertEquals(personUpdated.getAddress(), returnedPerson.getAddress()),
                () -> assertEquals(personUpdated.getAge(), returnedPerson.getAge())
        );
    }

    @Test
    void deletePerson() {
        PersonService personService = new PersonService(personRepo);
        doNothing().when(personRepo).deleteById(123);

        personService.deletePerson(123);

        verify(personRepo, times(1)).deleteById(123);
    }
}