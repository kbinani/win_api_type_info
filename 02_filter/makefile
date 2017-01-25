all: $(ARCH)\stdafx.min.c

$(ARCH)\stdafx.min.c: main.go ..\01_preprocess\$(ARCH)\01_preprocess.dir\Debug\stdafx.i
	@mkdir $(ARCH) > nul 2>&1 || true
	go run main.go ..\01_preprocess\$(ARCH)\01_preprocess.dir\Debug\stdafx.i $(ARCH)\stdafx.min.c $(ARCH)\included_files.txt

clean:
	del /Q $(ARCH)\stdafx.min.c >nul 2>&1 || true
