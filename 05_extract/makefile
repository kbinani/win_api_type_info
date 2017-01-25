all: $(ARCH)\typedef.json

$(ARCH)\typedef.json: ..\04_combine_xml\$(ARCH)\all.xml main.go doxygen.go ..\util.go
	@mkdir $(ARCH) > nul 2>&1 || true
	go run main.go doxygen.go ..\04_combine_xml\$(ARCH)\all.xml $(ARCH)

clean:
	del /Q $(ARCH)\typedef.json $(ARCH)\struct.json >nul 2>&1 || true
