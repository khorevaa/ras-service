@echo off
rem run this script as admin

if not exist .\bin\go_build_github_com_khorevaa_ras_service.exe (
    echo Build the example before installing by running "go build"
    goto :exit
)

sc create ras-service2 binpath= "%CD%\bin\go_build_github_com_khorevaa_ras_service.exe svc" start= auto DisplayName= "RAS Control Service"
sc description ras-service2 "ras-service"
net start ras-service2
sc query ras-service2

echo Check example.log

:exit