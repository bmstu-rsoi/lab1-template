find_package(PostgreSQL REQUIRED)

add_library(person_postgres)

target_sources(person_postgres
    PUBLIC
        FILE_SET CXX_MODULES
        FILES
            postgres_repository.cppm
)

target_link_libraries(person_postgres
    PUBLIC
        person_app
    PRIVATE
        PostgreSQL::PostgreSQL
)

target_compile_features(person_postgres PUBLIC cxx_std_20)