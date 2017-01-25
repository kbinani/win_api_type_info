all: ..\07_dump_struct_size\generated\$(ARCH)\main.cpp

..\07_dump_struct_size\generated\$(ARCH)\main.cpp: ..\05_extract\$(ARCH)\struct.json ..\05_extract\$(ARCH)\enum.json main.go blacklist.txt ..\util.go
	@mkdir ..\07_dump_struct_size\generated\$(ARCH) >nul 2>&1 || true
	go run main.go ..\05_extract\$(ARCH)\struct.json ..\05_extract\$(ARCH)\enum.json ..\07_dump_struct_size\generated\$(ARCH)

clean:
	@rd /S /Q ..\07_dump_struct_size\generated\$(ARCH) >nul 2>&1 || true
