module;

#include <cstdint>
#include <drogon/drogon.h>

export module adapters.http.drogon_server;

import app.person_service;
import adapters.http.person_controller;

export namespace adapters::http {

void run(app::PersonService& service, std::uint16_t port) {
    registerRoutes(service);
    drogon::app().addListener("0.0.0.0", port);
    drogon::app().run();
}

}