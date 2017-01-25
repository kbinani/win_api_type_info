project(01_preprocess)
cmake_minimum_required(VERSION 2.8)

add_definitions(/DUNICODE /D_WINSOCK_DEPRECATED_NO_WARNINGS)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} /P /MP /EHsc")

add_executable(01_preprocess stdafx.c)
include_directories(.)
