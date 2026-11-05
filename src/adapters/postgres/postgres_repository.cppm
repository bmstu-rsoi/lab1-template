
module;

#include <libpq-fe.h>

#include <stdexcept>
#include <string>
#include <string_view>
#include <vector>
#include <optional>

export module adapters.postgres.postgres_repository;

import app.person_repository;
import domain.person;

export namespace adapters::postgres {

class PostgresRepository final : public app::PersonRepository {
public:
    explicit PostgresRepository(const std::string& connectionString);
    ~PostgresRepository() override;

    PostgresRepository(const PostgresRepository&) = delete;
    PostgresRepository& operator=(const PostgresRepository&) = delete;

    PostgresRepository(PostgresRepository&&) = delete;
    PostgresRepository& operator=(PostgresRepository&&) = delete;

    std::vector<domain::Person> findAll() override;
    std::optional<domain::Person> findById(int id) override;
    domain::Person create(domain::Person person) override;
    domain::Person update(domain::Person person) override;
    bool remove(int id) override;

private:
    PGconn* connection_;

    void checkConnection();
};

}

namespace {

domain::Person mapPerson(PGresult* result, int row) {
    domain::Person person;

    person.id = std::stoi(PQgetvalue(result, row, 0));
    person.name = PQgetvalue(result, row, 1);

    if (!PQgetisnull(result, row, 2)) {
        person.age = std::stoi(PQgetvalue(result, row, 2));
    }

    if (!PQgetisnull(result, row, 3)) {
        person.address = PQgetvalue(result, row, 3);
    }

    if (!PQgetisnull(result, row, 4)) {
        person.work = PQgetvalue(result, row, 4);
    }

    return person;
}

void checkResult(PGresult* result, ExecStatusType expected) {
    if (!result) {
        throw std::runtime_error("PostgreSQL query failed: no result");
    }

    if (PQresultStatus(result) != expected) {
        const std::string error = PQresultErrorMessage(result);
        PQclear(result);

        throw std::runtime_error(
            "PostgreSQL query failed: " + error
        );
    }
}

}

namespace adapters::postgres {

PostgresRepository::PostgresRepository(
    const std::string& connectionString
)
    : connection_(PQconnectdb(connectionString.c_str())) {

    checkConnection();
}

PostgresRepository::~PostgresRepository() {
    if (connection_) {
        PQfinish(connection_);
    }
}

void PostgresRepository::checkConnection() {
    if (!connection_) {
        throw std::runtime_error(
            "PostgreSQL connection allocation failed"
        );
    }

    if (PQstatus(connection_) != CONNECTION_OK) {
        const std::string error = PQerrorMessage(connection_);

        PQfinish(connection_);
        connection_ = nullptr;

        throw std::runtime_error(
            "PostgreSQL connection failed: " + error
        );
    }
}

std::vector<domain::Person> PostgresRepository::findAll() {
    PGresult* result = PQexec(
        connection_,
        "SELECT id, name, age, address, work FROM persons ORDER BY id"
    );

    checkResult(result, PGRES_TUPLES_OK);

    const int rowCount = PQntuples(result);

    std::vector<domain::Person> persons;
    persons.reserve(rowCount);

    for (int row = 0; row < rowCount; ++row) {
        persons.push_back(mapPerson(result, row));
    }

    PQclear(result);

    return persons;
}

std::optional<domain::Person> PostgresRepository::findById(int id) {
    const std::string idValue = std::to_string(id);
    const char* params[] = {idValue.c_str()};

    PGresult* result = PQexecParams(
        connection_,
        "SELECT id, name, age, address, work FROM persons WHERE id = $1",
        1,
        nullptr,
        params,
        nullptr,
        nullptr,
        0
    );

    checkResult(result, PGRES_TUPLES_OK);

    if (PQntuples(result) == 0) {
        PQclear(result);
        return std::nullopt;
    }

    domain::Person person = mapPerson(result, 0);

    PQclear(result);

    return person;
}

domain::Person PostgresRepository::create(domain::Person person) {
    const std::string ageValue = person.age
        ? std::to_string(*person.age)
        : "";

    const char* params[] = {
        person.name.c_str(),
        person.age ? ageValue.c_str() : nullptr,
        person.address ? person.address->c_str() : nullptr,
        person.work ? person.work->c_str() : nullptr
    };

    PGresult* result = PQexecParams(
        connection_,
        "INSERT INTO persons (name, age, address, work) "
        "VALUES ($1, $2, $3, $4) "
        "RETURNING id, name, age, address, work",
        4,
        nullptr,
        params,
        nullptr,
        nullptr,
        0
    );

    checkResult(result, PGRES_TUPLES_OK);

    domain::Person created = mapPerson(result, 0);

    PQclear(result);

    return created;
}

domain::Person PostgresRepository::update(domain::Person person) {
    const std::string idValue = std::to_string(person.id);

    const std::string ageValue = person.age
        ? std::to_string(*person.age)
        : "";

    const char* params[] = {
        idValue.c_str(),
        person.name.c_str(),
        person.age ? ageValue.c_str() : nullptr,
        person.address ? person.address->c_str() : nullptr,
        person.work ? person.work->c_str() : nullptr
    };

    PGresult* result = PQexecParams(
        connection_,
        "UPDATE persons "
        "SET name = $2, age = $3, address = $4, work = $5 "
        "WHERE id = $1 "
        "RETURNING id, name, age, address, work",
        5,
        nullptr,
        params,
        nullptr,
        nullptr,
        0
    );

    checkResult(result, PGRES_TUPLES_OK);

    if (PQntuples(result) == 0) {
        PQclear(result);
        throw std::runtime_error(
            "Cannot update person: person not found"
        );
    }

    domain::Person updated = mapPerson(result, 0);

    PQclear(result);

    return updated;
}

bool PostgresRepository::remove(int id) {
    const std::string idValue = std::to_string(id);
    const char* params[] = {idValue.c_str()};

    PGresult* result = PQexecParams(
        connection_,
        "DELETE FROM persons WHERE id = $1",
        1,
        nullptr,
        params,
        nullptr,
        nullptr,
        0
    );

    checkResult(result, PGRES_COMMAND_OK);

    const bool removed = PQcmdTuples(result)[0] != '0';

    PQclear(result);

    return removed;
}

}
