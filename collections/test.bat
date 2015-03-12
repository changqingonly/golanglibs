
@echo off

rem cd %~dp0
set GOPATH=%CD%;%GOPATH%

go test -v .

