.PHONY: build
build: rsrc.syso
	GOOS=windows go build .

build-gui: rsrc.syso
	GOOS=windows go build -ldflags "-H windowsgui -s -w" .

icon.ico: icon.png
	convert -resize 256x256 icon.png icon.ico

rsrc.syso: icon.ico
	rsrc -ico icon.ico -o rsrc.syso

.PHONY: clean
clean:
	- rm icon.ico
	- rm rsrc.syso

