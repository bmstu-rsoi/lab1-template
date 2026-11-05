module;

#include <drogon/drogon.h>

export module adapters.http.drogon_server;

import app.person_service;
import adapters.http.person_controller;

export namespace adapters::http {

void run(app::PersonService& service) {
    registerRoutes(service);
    drogon::app().run();
}

}