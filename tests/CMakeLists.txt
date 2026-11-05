find_package(GTest CONFIG REQUIRED)

add_executable(person_service_tests
    app/person_service_test.cpp
)

target_link_libraries(person_service_tests
    PRIVATE
        GTest::gtest_main
        person_app
)

target_compile_features(person_service_tests PRIVATE cxx_std_20)

include(GoogleTest)
gtest_discover_tests(person_service_tests)