module;

#include <cstdlib>
#include <stdexcept>

export module composition.application;

import app.person_service;
import adapters.http.drogon_server;
import adapters.postgres.postgres_repository;

export namespace composition {

void run() {
    const char* connectionString = std::getenv("DATABASE_URL");

    if (!connectionString || *connectionString == '\0') {
        throw std::runtime_error(
            "Environment variable DATABASE_URL is not set"
        );
    }

    adapters::postgres::PostgresRepository repository(connectionString);
    app::PersonService service(repository);

    adapters::http::run(service);
}

}