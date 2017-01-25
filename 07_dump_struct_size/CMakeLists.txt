project(dump_struct_size)
cmake_minimum_required(VERSION 2.8)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/cmakemodules/add_msvc_precompiled_header)
include(add_msvc_precompiled_header)

add_definitions(/DUNICODE /D_WINSOCK_DEPRECATED_NO_WARNINGS)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} /MP /EHsc")

file(GLOB_RECURSE dump_struct_size_src "generated/${ARCH}/*.cpp")
add_msvc_precompiled_header("../01_preprocess/stdafx.h" "stdafx.cpp" dump_struct_size_src)

add_executable(dump_struct_size
	${dump_struct_size_src}
	generated/${ARCH}/main.cpp
	offsetof.hpp)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../01_preprocess)
set_target_properties(dump_struct_size PROPERTIES COMPILE_FLAGS "/bigobj")
