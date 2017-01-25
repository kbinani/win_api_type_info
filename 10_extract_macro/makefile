all: $(ARCH)\10_extract_macro.dir\Debug\macro_extractor.i

$(ARCH)\10_extract_macro.dir\Debug\macro_extractor.i: ../09_generate_macro_extractor/$(ARCH)/macro_extractor.c CMakeLists.txt
	@mkdir $(ARCH) > nul 2>&1 || true
	cd $(ARCH)
	cmake .. -G "$(VS_GENERATOR)" -DARCH=$(ARCH)
	cmake --build . > nul || true
	@cd $(MAKEDIR)

clean:
	rd /S /Q $(ARCH) >nul 2>&1 || true
