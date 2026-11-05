
module;

#include <drogon/drogon.h>

#include <functional>
#include <optional>
#include <string>
#include <utility>

export module adapters.http.person_controller;

import app.person_service;
import domain.person;

namespace {

using namespace drogon;

using ResponseCallback =
    std::function<void(const HttpResponsePtr&)>;

HttpResponsePtr jsonResponse(
    const Json::Value& body,
    HttpStatusCode status = k200OK
) {
    auto response = HttpResponse::newHttpJsonResponse(body);
    response->setStatusCode(status);
    return response;
}

HttpResponsePtr errorResponse(
    const std::string& message,
    HttpStatusCode status = k400BadRequest
) {
    Json::Value body;
    body["error"] = message;

    return jsonResponse(body, status);
}

HttpResponsePtr notFound() {
    return errorResponse("Person not found", k404NotFound);
}

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

bool parseCreate(
    const Json::Value* json,
    app::PersonCreate& data,
    HttpResponsePtr& error
) {
    if (!json || !json->isObject()) {
        error = errorResponse("Request body must be a JSON object");
        return false;
    }

    if (!json->isMember("name") || !(*json)["name"].isString()) {
        error = errorResponse(
            "Field 'name' is required and must be a string"
        );
        return false;
    }

    data.name = (*json)["name"].asString();

    if (json->isMember("age") && !(*json)["age"].isNull()) {
        if (!(*json)["age"].isInt()) {
            error = errorResponse("Field 'age' must be an integer");
            return false;
        }

        data.age = (*json)["age"].asInt();
    }

    if (json->isMember("address") && !(*json)["address"].isNull()) {
        if (!(*json)["address"].isString()) {
            error = errorResponse("Field 'address' must be a string");
            return false;
        }

        data.address = (*json)["address"].asString();
    }

    if (json->isMember("work") && !(*json)["work"].isNull()) {
        if (!(*json)["work"].isString()) {
            error = errorResponse("Field 'work' must be a string");
            return false;
        }

        data.work = (*json)["work"].asString();
    }

    return true;
}

bool parsePatch(
    const Json::Value* json,
    app::PersonPatch& patch,
    HttpResponsePtr& error
) {
    if (!json || !json->isObject()) {
        error = errorResponse("Request body must be a JSON object");
        return false;
    }

    if (json->isMember("name")) {
        if (!(*json)["name"].isString()) {
            error = errorResponse("Field 'name' must be a string");
            return false;
        }

        patch.name = (*json)["name"].asString();
    }

    if (json->isMember("age")) {
        if ((*json)["age"].isNull()) {
            patch.age.emplace(std::nullopt);
        } else {
            if (!(*json)["age"].isInt()) {
                error = errorResponse(
                    "Field 'age' must be an integer or null"
                );
                return false;
            }

            patch.age.emplace((*json)["age"].asInt());
        }
    }

    if (json->isMember("address")) {
        if ((*json)["address"].isNull()) {
            patch.address.emplace(std::nullopt);
        } else {
            if (!(*json)["address"].isString()) {
                error = errorResponse(
                    "Field 'address' must be a string or null"
                );
                return false;
            }

            patch.address.emplace((*json)["address"].asString());
        }
    }

    if (json->isMember("work")) {
        if ((*json)["work"].isNull()) {
            patch.work.emplace(std::nullopt);
        } else {
            if (!(*json)["work"].isString()) {
                error = errorResponse(
                    "Field 'work' must be a string or null"
                );
                return false;
            }

            patch.work.emplace((*json)["work"].asString());
        }
    }

    return true;
}

void handleFindAll(
    app::PersonService& service,
    const ResponseCallback& callback
) {
    const auto persons = service.findAll();

    Json::Value result(Json::arrayValue);

    for (const auto& person : persons) {
        result.append(toJson(person));
    }

    callback(jsonResponse(result));
}

void handleCreate(
    app::PersonService& service,
    const HttpRequestPtr& request,
    const ResponseCallback& callback
) {
    app::PersonCreate data;
    HttpResponsePtr error;
    const auto json = request->getJsonObject();

    if (!parseCreate(json.get(), data, error)) {
        callback(error);
        return;
    }

    const auto person = service.create(std::move(data));

    auto response = HttpResponse::newHttpResponse();
    response->setStatusCode(k201Created);
    response->addHeader(
        "Location",
        "/api/v1/persons/" + std::to_string(person.id)
    );

    callback(response);
}

void handleFindById(
    app::PersonService& service,
    int id,
    const ResponseCallback& callback
) {
    const auto person = service.findById(id);

    if (!person) {
        callback(notFound());
        return;
    }

    callback(jsonResponse(toJson(*person)));
}

void handleDelete(
    app::PersonService& service,
    int id,
    const ResponseCallback& callback
) {
    if (!service.remove(id)) {
        callback(notFound());
        return;
    }

    auto response = HttpResponse::newHttpResponse();
    response->setStatusCode(k204NoContent);

    callback(response);
}

void handleUpdate(
    app::PersonService& service,
    const HttpRequestPtr& request,
    int id,
    const ResponseCallback& callback
) {
    app::PersonPatch patch;
    HttpResponsePtr error;
    const auto json = request->getJsonObject();

    if (!parsePatch(json.get(), patch, error)) {
        callback(error);
        return;
    }

    const auto person = service.update(id, patch);

    if (!person) {
        callback(notFound());
        return;
    }

    callback(jsonResponse(toJson(*person)));
}

void registerCollectionRoutes(app::PersonService& service) {
    drogon::app().registerHandler(
        "/api/v1/persons",
        [&service](
            const HttpRequestPtr&,
            ResponseCallback&& callback
        ) {
            handleFindAll(service, callback);
        },
        {Get}
    );

    drogon::app().registerHandler(
        "/api/v1/persons",
        [&service](
            const HttpRequestPtr& request,
            ResponseCallback&& callback
        ) {
            handleCreate(service, request, callback);
        },
        {Post}
    );
}

void registerItemRoutes(app::PersonService& service) {
    drogon::app().registerHandler(
        "/api/v1/persons/{1}",
        [&service](
            const HttpRequestPtr&,
            ResponseCallback&& callback,
            int id
        ) {
            handleFindById(service, id, callback);
        },
        {Get}
    );

    drogon::app().registerHandler(
        "/api/v1/persons/{1}",
        [&service](
            const HttpRequestPtr&,
            ResponseCallback&& callback,
            int id
        ) {
            handleDelete(service, id, callback);
        },
        {Delete}
    );

    drogon::app().registerHandler(
        "/api/v1/persons/{1}",
        [&service](
            const HttpRequestPtr& request,
            ResponseCallback&& callback,
            int id
        ) {
            handleUpdate(service, request, id, callback);
        },
        {Patch}
    );
}

}

export namespace adapters::http {

void registerRoutes(app::PersonService& service) {
    registerCollectionRoutes(service);
    registerItemRoutes(service);
}

}
