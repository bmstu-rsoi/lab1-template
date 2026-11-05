if(POLICY CMP0153)
    cmake_policy(SET CMP0153 OLD)
endif()
find_package(Drogon CONFIG REQUIRED)

add_library(person_http)

target_sources(person_http
    PUBLIC
        FILE_SET CXX_MODULES
        FILES
            drogon_server.cppm
            person_controller.cppm
)

target_link_libraries(person_http
    PRIVATE
        Drogon::Drogon
    PUBLIC
        person_app
)

target_compile_features(person_http PUBLIC cxx_std_20)