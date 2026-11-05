module;

#include <optional>
#include <string>

export module domain.person;

export namespace domain {

    struct Person {
        int id;
        std::string name;
        std::optional<int> age;
        std::optional<std::string> address;
        std::optional<std::string> work;
    };

}
