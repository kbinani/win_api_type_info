all: $(ARCH)\01_preprocess.dir\Debug\stdafx.i

$(ARCH)\01_preprocess.dir\Debug\stdafx.i: CMakeLists.txt stdafx.c stdafx.h
	@mkdir $(ARCH) > nul 2>&1 || true
	cd $(ARCH)
	cmake .. -G "$(VS_GENERATOR)" >nul
	cmake --build . >nul || true
	@cd $(MAKEDIR)

clean:
	rd /S /Q $(ARCH) >nul 2>&1 || true
