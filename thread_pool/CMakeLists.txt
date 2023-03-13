cmake_minimum_required(VERSION 3.22)
project(thread_pool)

set(CMAKE_CXX_STANDARD 17)

include_directories(.)

add_executable(thread_pool
        thread_pool.cpp
        thread_pool.h)
