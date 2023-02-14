go build -ldflags "-s -w -X main.version=1.1.0" -o albiondata-client.exe -v -x albiondata-client.go

start thirdparty/rcedit/rcedit.exe albiondata-client.exe --set-icon icon/albiondata-client.ico

cd pkg/nsis
make nsis
cd ../..
