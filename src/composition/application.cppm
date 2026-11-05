module;

#include <charconv>
#include <cstdlib>
#include <cstdint>
#include <stdexcept>
#include <string>
#include <system_error>

export module composition.application;

import app.person_service;
import adapters.http.drogon_server;
import adapters.postgres.postgres_repository;

namespace {

std::string getRequiredEnv(const char* name) {
    const char* value = std::getenv(name);

    if (!value || *value == '\0') {
        throw std::runtime_error(
            std::string("Environment variable ") + name + " is not set"
        );
    }

    return value;
}

std::string quoteConninfoValue(const std::string& value) {
    std::string result = "'";

    for (char character : value) {
        if (character == '\\' || character == '\'') {
            result += '\\';
        }

        result += character;
    }

    result += '\'';
    return result;
}

std::uint16_t getHttpPort() {
    const char* value = std::getenv("HTTP_PORT");

    if (!value || *value == '\0') {
        return 8080;
    }

    unsigned int port = 0;
    const std::string portString(value);

    const auto [end, error] = std::from_chars(
        portString.data(),
        portString.data() + portString.size(),
        port
    );

    if (error != std::errc{} ||
        end != portString.data() + portString.size() ||
        port == 0 ||
        port > 65535) {
        throw std::runtime_error("Invalid HTTP_PORT");
    }

    return static_cast<std::uint16_t>(port);
}

std::string getDatabaseConnectionString() {
    return
        "host=" + quoteConninfoValue(getRequiredEnv("POSTGRES_HOST")) +
        " port=" + quoteConninfoValue(getRequiredEnv("POSTGRES_PORT")) +
        " dbname=" + quoteConninfoValue(getRequiredEnv("POSTGRES_DB")) +
        " user=" + quoteConninfoValue(getRequiredEnv("POSTGRES_USER")) +
        " password=" + quoteConninfoValue(getRequiredEnv("POSTGRES_PASSWORD"));
}

}

export namespace composition {

void run() {
    const std::string connectionString = getDatabaseConnectionString();
    const std::uint16_t httpPort = getHttpPort();

    adapters::postgres::PostgresRepository repository(
        connectionString.c_str()
    );

    app::PersonService service(repository);

    adapters::http::run(service, httpPort);
}

}