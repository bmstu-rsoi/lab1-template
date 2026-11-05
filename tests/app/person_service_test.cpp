#include <gtest/gtest.h>

#include <optional>
#include <string>
#include <vector>
#include <algorithm>
#include <stdexcept>
#include <utility>

import domain.person;
import app.person_repository;
import app.person_service;

namespace {

class FakePersonRepository final : public app::PersonRepository {
public:
    std::vector<domain::Person> persons;
    int nextId = 1;

    std::vector<domain::Person> findAll() override {
        return persons;
    }

    std::optional<domain::Person> findById(int id) override {
        auto it = std::find_if(
            persons.begin(),
            persons.end(),
            [id](const auto& person) {
                return person.id == id;
            }
        );

        if (it == persons.end()) {
            return std::nullopt;
        }

        return *it;
    }

    domain::Person create(domain::Person person) override {
        person.id = nextId++;
        persons.push_back(person);
        return person;
    }

    domain::Person update(domain::Person person) override {
        auto it = std::find_if(
            persons.begin(),
            persons.end(),
            [&person](const auto& current) {
                return current.id == person.id;
            }
        );

        if (it == persons.end()) {
            throw std::runtime_error("Person not found");
        }

        *it = person;
        return person;
    }

    bool remove(int id) override {
        const auto oldSize = persons.size();

        std::erase_if(
            persons,
            [id](const auto& person) {
                return person.id == id;
            }
        );

        return persons.size() != oldSize;
    }
};

class PersonServiceTest : public ::testing::Test {
protected:
    FakePersonRepository repository;
    app::PersonService service{repository};

    void SetUp() override {
        repository.persons.push_back({
            1,
            "Ivan Ivanov",
            25,
            "Moscow",
            "Developer"
        });

        repository.nextId = 2;
    }
};

TEST_F(PersonServiceTest, FindAllReturnsAllPersons) {
    const auto persons = service.findAll();

    ASSERT_EQ(persons.size(), 1);
    EXPECT_EQ(persons[0].name, "Ivan Ivanov");
}

TEST_F(PersonServiceTest, FindByIdReturnsPerson) {
    const auto person = service.findById(1);

    ASSERT_TRUE(person.has_value());
    EXPECT_EQ(person->name, "Ivan Ivanov");
}

TEST_F(PersonServiceTest, FindByIdReturnsNulloptForMissingPerson) {
    EXPECT_FALSE(service.findById(999).has_value());
}

TEST_F(PersonServiceTest, CreateReturnsCreatedPerson) {
    app::PersonCreate data{
        "Petr Petrov",
        30,
        "Kazan",
        "Engineer"
    };

    const auto person = service.create(std::move(data));

    EXPECT_EQ(person.id, 2);
    EXPECT_EQ(person.name, "Petr Petrov");
    EXPECT_EQ(repository.persons.size(), 2);
}

TEST_F(PersonServiceTest, UpdateChangesSpecifiedFields) {
    app::PersonPatch patch;
    patch.name = "Petr Petrov";
    patch.age = 30;

    const auto person = service.update(1, patch);

    ASSERT_TRUE(person.has_value());
    EXPECT_EQ(person->name, "Petr Petrov");
    EXPECT_EQ(person->age, 30);
    EXPECT_EQ(person->address, "Moscow");
}

TEST_F(PersonServiceTest, UpdateClearsNullableField) {
    app::PersonPatch patch;
    patch.address = std::optional<std::string>{};

    const auto person = service.update(1, patch);

    ASSERT_TRUE(person.has_value());
    EXPECT_FALSE(person->address.has_value());
}

TEST_F(PersonServiceTest, UpdateReturnsNulloptForMissingPerson) {
    app::PersonPatch patch;
    patch.name = "Petr";

    EXPECT_FALSE(service.update(999, patch).has_value());
}

TEST_F(PersonServiceTest, RemoveDeletesExistingPerson) {
    EXPECT_TRUE(service.remove(1));
    EXPECT_TRUE(repository.persons.empty());
}

TEST_F(PersonServiceTest, RemoveReturnsFalseForMissingPerson) {
    EXPECT_FALSE(service.remove(999));
}

}