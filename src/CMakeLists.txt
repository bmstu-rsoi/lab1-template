add_subdirectory(domain)
add_subdirectory(app)
add_subdirectory(adapters)
add_subdirectory(composition)

add_executable(person_service main.cpp)

target_link_libraries(person_service
    PRIVATE
        person_composition
)

target_compile_features(person_service PRIVATE cxx_std_20)