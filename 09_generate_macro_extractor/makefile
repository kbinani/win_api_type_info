all: $(ARCH)\macro_extractor.c

$(ARCH)\macro_extractor.c: main.go ..\02_filter\$(ARCH)\included_files.txt
	@mkdir $(ARCH) > nul 2>&1 || true
	go run main.go ..\02_filter\$(ARCH)\included_files.txt $(ARCH)\macro_extractor.c

clean:
	del /Q $(ARCH)\macro_extractor.c >nul 2>&1 || true
