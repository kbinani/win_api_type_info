all:
	cd 01_preprocess && nmake /NOLOGO
	cd 02_filter && nmake /NOLOGO
	cd 03_doxygen && nmake /NOLOGO
	cd 04_combine_xml && nmake /NOLOGO
	cd 05_extract && nmake /NOLOGO
	cd 06_generate_struct_size_dumper && nmake /NOLOGO
	cd 07_dump_struct_size && nmake /NOLOGO
	cd 08_update_json && nmake /NOLOGO
	cd 09_generate_macro_extractor && nmake /NOLOGO
	cd 10_extract_macro && nmake /NOLOGO
	cd 11_const_json && nmake /NOLOGO
	copy /Y 08_update_json\$(ARCH)\struct.json struct_$(ARCH).json >nul
	copy /Y 08_update_json\$(ARCH)\enum.json enum_$(ARCH).json >nul
	copy /Y 05_extract\$(ARCH)\typedef.json typedef_$(ARCH).json >nul
	copy /Y 11_const_json\$(ARCH)\const.json const_$(ARCH).json >nul

clean:
	cd 01_preprocess && nmake /NOLOGO clean
	cd 02_filter && nmake /NOLOGO clean
	cd 03_doxygen && nmake /NOLOGO clean
	cd 04_combine_xml && nmake /NOLOGO clean
	cd 05_extract && nmake /NOLOGO clean
	cd 06_generate_struct_size_dumper && nmake /NOLOGO clean
	cd 07_dump_struct_size && nmake /NOLOGO clean
	cd 08_update_json && nmake /NOLOGO clean
	cd 09_generate_macro_extractor && nmake /NOLOGO clean
	cd 10_extract_macro && nmake /NOLOGO clean
	cd 11_const_json && nmake /NOLOGO clean
	del /Q struct_$(ARCH).json typedef_$(ARCH).json enum_$(ARCH).json const_$(ARCH).json >nul 2>&1 || true
