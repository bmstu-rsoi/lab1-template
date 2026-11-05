
module;

#include <drogon/drogon.h>

#include <string>
#include <string_view>

export module adapters.http.person_controller;

import app.person_service;
import domain.person;

namespace {

using namespace drogon;

Json::Value toJson(const domain::Person& person) {
    Json::Value json;

    json["id"] = person.id;
    json["name"] = person.name;

    if (person.age) {
        json["age"] = *person.age;
    }

    if (person.address) {
        json["address"] = *person.address;
    }

    if (person.work) {
        json["work"] = *person.work;
    }

    return json;
}

HttpResponsePtr jsonResponse(
    const Json::Value& body,
    HttpStatusCode status = k200OK
) {
    auto response = HttpResponse::newHttpJsonResponse(body);
    response->setStatusCode(status);
    return response;
}

HttpResponsePtr notFound() {
    Json::Value body;
    body["error"] = "Person not found";

    return jsonResponse(body, k404NotFound);
}

} // namespace

export namespace adapters::http {

void registerRoutes(app::PersonService& service) {

    drogon::app().registerHandler(
        "/api/v1/persons",
        [&service](const drogon::HttpRequestPtr&,
                   std::function<void(const drogon::HttpResponsePtr&)>&& callback) {

            const auto persons = service.findAll();

            Json::Value result(Json::arrayValue);

            for (const auto& person : persons) {
                result.append(toJson(person));
            }

            callback(jsonResponse(result));
        },
        {drogon::Get}
    );

    drogon::app().registerHandler(
        "/api/v1/persons/{1}",
        [&service](const drogon::HttpRequestPtr& request,
                   std::function<void(const drogon::HttpResponsePtr&)>&& callback,
                   int id) {

            if (request->method() == drogon::Get) {
                const auto person = service.findById(id);

                if (!person) {
                    callback(notFound());
                    return;
                }

                callback(jsonResponse(toJson(*person)));
                return;
            }

            if (request->method() == drogon::Delete) {
                if (!service.remove(id)) {
                    callback(notFound());
                    return;
                }

                auto response = drogon::HttpResponse::newHttpResponse();
                response->setStatusCode(drogon::k204NoContent);
                callback(response);
            }
        },
        {drogon::Get, drogon::Delete}
    );

}

}
