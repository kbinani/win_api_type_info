all: $(ARCH)\xml\index.xml

$(ARCH)\xml\index.xml: ..\02_filter\$(ARCH)\stdafx.min.c Doxyfile
	@mkdir $(ARCH) > nul 2>&1 || true
	copy ..\02_filter\$(ARCH)\stdafx.min.c $(ARCH) >nul
	cd $(ARCH)
	doxygen -s ..\Doxyfile >nul
	del /Q *.tmp >nul 2>&1
	@cd $(MAKEDIR)

clean:
	rd /S /Q $(ARCH)\xml >nul 2>&1 || true
	del /Q $(ARCH)\*.tmp >nul 2>&1
