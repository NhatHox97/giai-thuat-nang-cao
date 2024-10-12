cmake_minimum_required(VERSION 3.10)
project(TSP)

set(CMAKE_CXX_STANDARD 17)

# Include header files
include_directories(include)

# Add the executable
add_executable(TSP main.cpp src/Graph.cpp src/TSP.cpp)