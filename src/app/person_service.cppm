module;

#include <optional>
#include <string>
#include <utility>
#include <vector>

export module app.person_service;

import domain.person;
import app.person_repository;

export namespace app {

    struct PersonCreate {
        std::string name;
        std::optional<int> age;
        std::optional<std::string> address;
        std::optional<std::string> work;
    };

    struct PersonPatch {
        std::optional<std::string> name;
        std::optional<int> age;
        std::optional<std::string> address;
        std::optional<std::string> work;
    };

    class PersonService {
    public:
        explicit PersonService(PersonRepository& repository)
            : repository_(repository) {}

        std::vector<domain::Person> findAll() {
            return repository_.findAll();
        }

        std::optional<domain::Person> findById(int id) {
            return repository_.findById(id);
        }

        domain::Person create(PersonCreate data) {
            domain::Person person{
                0,
                std::move(data.name),
                data.age,
                std::move(data.address),
                std::move(data.work)
            };

            return repository_.create(std::move(person));
        }

        std::optional<domain::Person> update(int id, const PersonPatch& patch) {
            auto person = repository_.findById(id);

            if (!person) {
                return std::nullopt;
            }

            if (patch.name) person->name = *patch.name;
            if (patch.age) person->age = patch.age;
            if (patch.address) person->address = patch.address;
            if (patch.work) person->work = patch.work;

            return repository_.update(std::move(*person));
        }

        bool remove(int id) {
            return repository_.remove(id);
        }

    private:
        PersonRepository& repository_;
    };

}