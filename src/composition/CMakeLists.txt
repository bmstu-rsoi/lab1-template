add_library(person_composition)

target_sources(person_composition
    PUBLIC
        FILE_SET CXX_MODULES
        FILES
            application.cppm
)

target_link_libraries(person_composition
    PUBLIC
        person_app
        person_http
        person_postgres
)

target_compile_features(person_composition PUBLIC cxx_std_20)