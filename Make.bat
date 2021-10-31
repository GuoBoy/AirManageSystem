chdir "./frontend/air-manage-system"
call
yarn run build
call
chdir "../../"
@REM go build -o AirManageSystem.exe main.go
go build -o AirManageSystem.exe -ldflags '-w -s' main.go
md AirManageSystem

copy "AirManageSystem.exe" "./AirManageSystem/AirManageSystem.exe"
xcopy "./frontend/air-manage-system/dist" "./AirManageSystem/dist"
xcopy "./assets" "./AirManageSystem/assets"