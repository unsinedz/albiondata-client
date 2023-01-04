TOP_SRCDIR = ../..
PACKAGE = albiondata-client
PACKAGE_NAME = Albion Private Prices Client
PACKAGE_BUGREPORT = https://github.com/broderickhyman/albiondata-client/issues
INSTALLER_COMPRESSION = lzma
PACKAGE_EXE = albiondata-client.exe

all: nsis

nsis:
	makensis -DPACKAGE_VERSION="1.1.0" -DPACKAGE="$(PACKAGE)" -DPACKAGE_NAME="$(PACKAGE_NAME)" -DPACKAGE_EXE="$(PACKAGE_EXE)" -DPACKAGE_BUGREPORT="$(PACKAGE_BUGREPORT)" -DTOP_SRCDIR="$(TOP_SRCDIR)" -DOUTFILE=../../albiondata-client-amd64-installer.exe -X"SetCompressor /FINAL /SOLID $(INSTALLER_COMPRESSION)" albiondata-client.nsi
