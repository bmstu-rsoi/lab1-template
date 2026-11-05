add_library(person_app)

target_sources(person_app
    PUBLIC
        FILE_SET CXX_MODULES
        FILES
            person_repository.cppm
            person_service.cppm
)

target_link_libraries(person_app PUBLIC person_domain)
target_compile_features(person_app PUBLIC cxx_std_20)