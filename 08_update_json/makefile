all: $(ARCH)\struct.json

$(ARCH)\struct.json: ..\07_dump_struct_size\$(ARCH)\result.txt ..\05_extract\$(ARCH)\struct.json ..\05_extract\$(ARCH)\enum.json main.go ..\util.go
	@mkdir $(ARCH) > nul 2>&1 || true
	go run main.go ..\07_dump_struct_size\$(ARCH)\result.txt ..\05_extract\$(ARCH)\struct.json ..\05_extract\$(ARCH)\enum.json $(ARCH)\struct.json $(ARCH)\enum.json

clean:
	del /Q $(ARCH)\struct.json $(ARCH)\enum.json >nul 2>&1 || true
