@echo off

setlocal

call "%VS140COMNTOOLS%\vsvars32.bat"

set VS_VERSION=14 2015

set ARCH=win32
set VS_GENERATOR=Visual Studio %VS_VERSION%
nmake /NOLOGO clean

set ARCH=x64
set VS_GENERATOR=Visual Studio %VS_VERSION% Win64
nmake /NOLOGO clean
