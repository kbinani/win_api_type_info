all: $(ARCH)\const.json

$(ARCH)\const.json: main.go ..\10_extract_macro\$(ARCH)\10_extract_macro.dir\Debug\macro_extractor.i
	@mkdir $(ARCH) > nul 2>&1 || true
	go run main.go ..\10_extract_macro\$(ARCH)\10_extract_macro.dir\Debug\macro_extractor.i $(ARCH)\const.json

clean:
	rd /S /Q $(ARCH) >nul 2>&1 || true
