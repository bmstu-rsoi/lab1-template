
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

HttpResponsePtr validationError(
    const std::string& field,
    const std::string& message
) {
    Json::Value body;
    body["message"] = "Validation failed";
    body["errors"] = Json::Value(Json::objectValue);
    body["errors"][field] = message;

    return jsonResponse(body, k400BadRequest);
}

HttpResponsePtr invalidBody() {
    return validationError(
        "body",
        "Request body must be a JSON object"
    );
}

HttpResponsePtr notFound() {
    Json::Value body;
    body["message"] = "Person not found";

    return jsonResponse(body, k404NotFound);
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
        error = invalidBody();
        return false;
    }

    if (!json->isMember("name")) {
        error = validationError("name", "Field is required");
        return false;
    }

    if (!(*json)["name"].isString()) {
        error = validationError("name", "Must be a string");
        return false;
    }

    data.name = (*json)["name"].asString();

    if (json->isMember("age")) {
        if (!(*json)["age"].isInt()) {
            error = validationError("age", "Must be an integer");
            return false;
        }

        data.age = (*json)["age"].asInt();
    }

    if (json->isMember("address")) {
        if (!(*json)["address"].isString()) {
            error = validationError("address", "Must be a string");
            return false;
        }

        data.address = (*json)["address"].asString();
    }

    if (json->isMember("work")) {
        if (!(*json)["work"].isString()) {
            error = validationError("work", "Must be a string");
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
        error = invalidBody();
        return false;
    }

    if (!json->isMember("name")) {
        error = validationError("name", "Field is required");
        return false;
    }

    if (!(*json)["name"].isString()) {
        error = validationError("name", "Must be a string");
        return false;
    }

    patch.name = (*json)["name"].asString();

    if (json->isMember("age")) {
        if (!(*json)["age"].isInt()) {
            error = validationError("age", "Must be an integer");
            return false;
        }

        patch.age = (*json)["age"].asInt();
    }

    if (json->isMember("address")) {
        if (!(*json)["address"].isString()) {
            error = validationError("address", "Must be a string");
            return false;
        }

        patch.address = (*json)["address"].asString();
    }

    if (json->isMember("work")) {
        if (!(*json)["work"].isString()) {
            error = validationError("work", "Must be a string");
            return false;
        }

        patch.work = (*json)["work"].asString();
    }

    return true;
}

void handleFindAll(
    app::PersonService& service,
    const ResponseCallback& callback
) {
    Json::Value result(Json::arrayValue);

    for (const auto& person : service.findAll()) {
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
    service.remove(id);

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
