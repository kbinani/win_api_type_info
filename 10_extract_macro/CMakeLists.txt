project(10_extract_macro)
cmake_minimum_required(VERSION 2.8)

add_definitions(/DUNICODE /D_WINSOCK_DEPRECATED_NO_WARNINGS)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} /P /EP /MP /EHsc")

add_executable(10_extract_macro ../09_generate_macro_extractor/${ARCH}/macro_extractor.c)
include_directories(../01_preprocess)
