#pragma once

#include <vector>

#define OFFSETBITOF(__type, __member) []() -> size_t { \
    std::vector<uint8_t> value_storage(sizeof(__type)); \
    typedef decltype(__type::__member) member_type; \
    std::vector<uint8_t> member_storage(sizeof(member_type), 0xff); \
    __type* value = (__type*)value_storage.data(); \
    member_type* member = (member_type*)member_storage.data(); \
    value->__member = *member; \
    size_t offset = 0; \
    for (size_t i = 0; i < sizeof(__type); ++i) { \
        uint8_t b = value_storage[i]; \
        if (b == 0) { \
            continue; \
        } \
        for (size_t j = 0; j < 8; ++j) { \
            if ((b & (0x1 << j)) == (0x1 << j)) { \
                return i * 8 + j; \
            } \
        } \
    } \
    return 0; \
} ()
