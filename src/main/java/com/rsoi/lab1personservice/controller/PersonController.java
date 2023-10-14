package com.rsoi.lab1personservice.controller;

import com.rsoi.lab1personservice.dto.request.PersonRequest;
import com.rsoi.lab1personservice.dto.response.PersonResponse;
import com.rsoi.lab1personservice.service.PersonService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.headers.Header;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.responses.ApiResponses;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.hibernate.ObjectNotFoundException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.util.List;
import java.util.Optional;

@RestController
@Validated
@RequestMapping("/api/v1")
@Tag(name = "Person REST API operations")
public class PersonController {
    @Autowired
    PersonService personService;

    @Operation(summary = "Get person by id", description = "Returns a person with given id")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Successfully retrieved", content = {
                    @Content(schema= @Schema(implementation = PersonResponse.class))}
            ),
            @ApiResponse(responseCode = "404", description = "Not found - The person was not found")
    })
    @GetMapping("/persons/{id}")
    public ResponseEntity<?> getPersonById(@PathVariable(name = "id") Integer id) {
        return Optional.ofNullable(personService
                        .getPersonById(id))
                .map(person -> ResponseEntity.ok().body(person))
                .orElseGet(() -> ResponseEntity.notFound().build());
    }

    @Operation(summary = "Get all persons", description = "Returns all persons")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Successfully retrieved")
    })
    @GetMapping("/persons")
    public List<PersonResponse> getPersons() {
        return personService.getPersons();
    }

    @Operation(summary = "Create person", description = "Returns Location header with created person's id")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "201", description = "Successfully created", headers =
                    {@Header(name = "Location", description = "Path to new person")}
            ),
            @ApiResponse(responseCode = "400", description = "Error while parsing json")
    })
    @PostMapping("/persons")
    public ResponseEntity<?> createPerson(@Valid @RequestBody PersonRequest personRequest) {
        Integer createdPersonId = personService.createPerson(personRequest);
        return ResponseEntity.created(URI.create( "/api/v1/persons/" + createdPersonId)).build();
    }

    @Operation(summary = "Update person by id", description = "Returns an updated person")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "200", description = "Successfully updated", content = {
                    @Content(schema= @Schema(implementation = PersonResponse.class))}),
            @ApiResponse(responseCode = "404", description = "Not found - The person was not found"),
            @ApiResponse(responseCode = "400", description = "Error while parsing json")
    })
    @PatchMapping("/persons/{id}")
    public ResponseEntity<?> updatePerson(@Valid @RequestBody PersonRequest personRequest, @PathVariable("id") Integer id) {
        try {
            return ResponseEntity.ok().body(personService.updatePerson(id, personRequest));
        } catch (ObjectNotFoundException objectNotFoundException) {
            return ResponseEntity.notFound().build();
        }
    }

    @Operation(summary = "Delete person by id", description = "Delete person")
    @ApiResponses(value = {
            @ApiResponse(responseCode = "204", description = "Successfully deleted"),
    })
    @DeleteMapping("/persons/{id}")
    public ResponseEntity<?> deletePerson(@PathVariable(name = "id") Integer personId) {
        personService.deletePerson(personId);
        return ResponseEntity.noContent().build();
    }
}
