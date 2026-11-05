add_library(person_domain)

target_sources(person_domain
    PUBLIC
        FILE_SET CXX_MODULES
        FILES
            person.cppm
)

target_compile_features(person_domain PUBLIC cxx_std_20)
