
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

    drogon::app().registerHandler(
    "/api/v1/persons",
    [&service](const drogon::HttpRequestPtr& request,
               std::function<void(const drogon::HttpResponsePtr&)>&& callback) {

        const auto& json = request->getJsonObject();

        if (!json || !json->isMember("name") || !(*json)["name"].isString()) {
            Json::Value error;
            error["error"] = "Field 'name' is required and must be a string";

            callback(jsonResponse(error, drogon::k400BadRequest));
            return;
        }

        app::PersonCreate data;
        data.name = (*json)["name"].asString();

        if (json->isMember("age") && !(*json)["age"].isNull()) {
            if (!(*json)["age"].isInt()) {
                Json::Value error;
                error["error"] = "Field 'age' must be an integer";

                callback(jsonResponse(error, drogon::k400BadRequest));
                return;
            }

            data.age = (*json)["age"].asInt();
        }

        if (json->isMember("address") && !(*json)["address"].isNull()) {
            if (!(*json)["address"].isString()) {
                Json::Value error;
                error["error"] = "Field 'address' must be a string";

                callback(jsonResponse(error, drogon::k400BadRequest));
                return;
            }

            data.address = (*json)["address"].asString();
        }

        if (json->isMember("work") && !(*json)["work"].isNull()) {
            if (!(*json)["work"].isString()) {
                Json::Value error;
                error["error"] = "Field 'work' must be a string";

                callback(jsonResponse(error, drogon::k400BadRequest));
                return;
            }

            data.work = (*json)["work"].asString();
        }

        const auto person = service.create(std::move(data));

        auto response = drogon::HttpResponse::newHttpResponse();
        response->setStatusCode(drogon::k201Created);
        response->addHeader(
            "Location",
            "/api/v1/persons/" + std::to_string(person.id)
        );

        callback(response);
    },
    {drogon::Post}
);

drogon::app().registerHandler(
    "/api/v1/persons/{1}",
    [&service](const drogon::HttpRequestPtr& request,
               std::function<void(const drogon::HttpResponsePtr&)>&& callback,
               int id) {

        const auto& json = request->getJsonObject();

        if (!json || !json->isObject()) {
            Json::Value error;
            error["error"] = "Request body must be a JSON object";

            callback(jsonResponse(error, drogon::k400BadRequest));
            return;
        }

        app::PersonPatch patch;

        if (json->isMember("name")) {
            if (!(*json)["name"].isString()) {
                Json::Value error;
                error["error"] = "Field 'name' must be a string";

                callback(jsonResponse(error, drogon::k400BadRequest));
                return;
            }

            patch.name = (*json)["name"].asString();
        }

        if (json->isMember("age")) {
            if ((*json)["age"].isNull()) {
                patch.age = std::nullopt;
            } else {
                if (!(*json)["age"].isInt()) {
                    Json::Value error;
                    error["error"] = "Field 'age' must be an integer or null";

                    callback(jsonResponse(error, drogon::k400BadRequest));
                    return;
                }

                patch.age = (*json)["age"].asInt();
            }
        }

        if (json->isMember("address")) {
            if ((*json)["address"].isNull()) {
                patch.address = std::nullopt;
            } else {
                if (!(*json)["address"].isString()) {
                    Json::Value error;
                    error["error"] = "Field 'address' must be a string or null";

                    callback(jsonResponse(error, drogon::k400BadRequest));
                    return;
                }

                patch.address = (*json)["address"].asString();
            }
        }

        if (json->isMember("work")) {
            if ((*json)["work"].isNull()) {
                patch.work = std::nullopt;
            } else {
                if (!(*json)["work"].isString()) {
                    Json::Value error;
                    error["error"] = "Field 'work' must be a string or null";

                    callback(jsonResponse(error, drogon::k400BadRequest));
                    return;
                }

                patch.work = (*json)["work"].asString();
            }
        }

        const auto person = service.update(id, patch);

        if (!person) {
            callback(notFound());
            return;
        }

        callback(jsonResponse(toJson(*person)));
    },
    {drogon::Patch}
);

}

}
