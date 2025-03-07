run:
	scripts/run.sh

run-windows:
	cmd /c .\scripts\run.cmd

fmt:
	scripts/fmt.sh

validate-fmt:
	scripts/validate-fmt.sh

build-windows:
	cmd /c .\scripts\build-windows.cmd

build-linux:
	scripts/build-linux.sh

build-darwin:
	scripts/build-darwin.sh
