module;

#include <optional>
#include <vector>

export module app.person_repository;

import domain.person;

export namespace app {
    
    class PersonRepository {
    public:
        virtual ~PersonRepository() = default;

        virtual std::vector<domain::Person> findAll() = 0;
        virtual std::optional<domain::Person> findById(int id) = 0;
        virtual domain::Person create(domain::Person person) = 0;
        virtual domain::Person update(domain::Person person) = 0;
        virtual bool remove(int id) = 0;
    };

}