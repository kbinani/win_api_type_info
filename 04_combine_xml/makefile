libxsltdir=libxslt-1.1.26.win32
libxml2dir=libxml2-2.7.8.win32
libiconvdir=iconv-1.9.2.win32
zlibdir=zlib-1.2.5

all: $(ARCH)\all.xml

$(ARCH)\all.xml: bin\xsltproc.exe ..\03_doxygen\$(ARCH)\xml\index.xml
	bin\xsltproc.exe --output $(ARCH)\all.xml ..\03_doxygen\$(ARCH)\xml\combine.xslt ..\03_doxygen\$(ARCH)\xml\index.xml

bin\xsltproc.exe: $(libxsltdir)\bin\xsltproc.exe $(libxml2dir)\bin\libxml2.dll $(libiconvdir)\bin\iconv.dll $(zlibdir)\bin\zlib1.dll
	@mkdir bin > nul 2>&1 || true
	xcopy /Y $(libxsltdir)\bin\xsltproc.exe .\bin
	xcopy /Y $(libxsltdir)\bin\libexslt.dll .\bin
	xcopy /Y $(libxsltdir)\bin\libxslt.dll .\bin
	xcopy /Y $(libxml2dir)\bin\libxml2.dll .\bin
	xcopy /Y $(libiconvdir)\bin\iconv.dll .\bin
	xcopy /Y $(zlibdir)\bin\zlib1.dll .\bin
	cd bin
	copy /b xsltproc.exe+ > nul
	@cd $(MAKEDIR)

$(libxsltdir)\bin\xsltproc.exe: $(libxsltdir).zip
	@rd /S /Q $(libxsltdir) > nul 2>&1 || true
	unzip $(libxsltdir).zip
	cd $(libxsltdir)\bin
	copy /b xsltproc.exe+ > nul
	@cd $(MAKEDIR)

$(libxml2dir)\bin\libxml2.dll: $(libxml2dir).zip
	@rd /S /Q $(libxml2dir) > nul 2>&1 || true
	unzip $(libxml2dir).zip
	cd $(libxml2dir)\bin
	copy /b libxml2.dll+ > nul
	@cd $(MAKEDIR)

$(libiconvdir)\bin\iconv.dll: $(libiconvdir).zip
	@rd /S /Q $(libiconvdir) > nul 2>&1 || true
	unzip $(libiconvdir).zip
	cd $(libiconvdir)\bin
	copy /b iconv.dll+ > nul
	@cd $(MAKEDIR)

$(zlibdir)\bin\zlib1.dll: $(zlibdir).win32.zip
	@rd /S /Q $(zlibdir) > nul 2>&1 || true
	unzip $(zlibdir).win32.zip
	cd $(zlibdir)\bin
	copy /b zlib1.dll+ > nul
	@cd $(MAKEDIR)

$(libxsltdir).zip:
	wget "ftp://ftp.zlatkovic.com/libxml/$(libxsltdir).zip"

$(libxml2dir).zip:
	wget "ftp://ftp.zlatkovic.com/libxml/$(libxml2dir).zip"

$(libiconvdir).zip:
	wget "ftp://ftp.zlatkovic.com/libxml/$(libiconvdir).zip"

$(zlibdir).win32.zip:
	wget "ftp://ftp.zlatkovic.com/libxml/$(zlibdir).win32.zip"

clean:
	del $(ARCH)\all.xml >nul 2>&1
