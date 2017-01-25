all: $(ARCH)\result.txt

$(ARCH)\result.txt: $(ARCH)\Debug\dump_struct_size.exe
	$(ARCH)\Debug\dump_struct_size.exe > $(ARCH)\result.txt

$(ARCH)\dump_struct_size.vcxproj: CMakeLists.txt generated\$(ARCH)\main.cpp
	@mkdir $(ARCH) > nul 2>&1 || true
	cd $(ARCH)
	cmake .. -G "$(VS_GENERATOR)" -DARCH=$(ARCH) >nul
	copy /b dump_struct_size.vcxproj+ > nul
	@cd $(MAKEDIR)

$(ARCH)\Debug\dump_struct_size.exe: $(ARCH)\dump_struct_size.vcxproj generated/$(ARCH)/main.cpp ..\01_preprocess\stdafx.h
	@mkdir $(ARCH) > nul 2>&1 || true
	cd $(ARCH)
	cmake --build . >..\..\log\7_dump_struct_size_$(ARCH)_stdout.txt 2>..\..\log\7_dump_struct_size_$(ARCH)_stderr.txt
	cd Debug && copy /b dump_struct_size.exe+ > nul
	@cd $(MAKEDIR)

clean:
	del /Q $(ARCH)\result.txt >nul 2>&1 || true
