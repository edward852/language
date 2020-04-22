cmake_minimum_required(VERSION 3.10)
project(alignment)

set(CMAKE_CXX_STANDARD 11)

# flags to be used when compiling 32 bit program
# g++ multilib installed
#set(CMAKE_CXX_FLAGS -m32)

add_executable(alignment main.cpp)