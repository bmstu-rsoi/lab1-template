cmake_minimum_required(VERSION 3.28)

set(CMAKE_EXPERIMENTAL_CXX_MODULE_CMAKE_API
    "0e5b6991-d74f-4b3d-a41c-cf096e0b2508"
)

project(person_service LANGUAGES CXX)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_CXX_SCAN_FOR_MODULES ON)

add_subdirectory(src)

include(CTest)

if(BUILD_TESTING)
    add_subdirectory(tests)
endif()