install: supported
	go install -p 6 . ./cursors ./focus \
		./frame ./heads ./hook ./layout ./logger ./misc ./render \
		./stack ./wingo-cmd ./wini ./wm ./workspace ./xclient

gofmt:
	gofmt -w *.go cursors/*.go focus/*.go frame/*.go \
		heads/*.go hook/*.go layout/*.go logger/*.go misc/*.go \
		render/*.go stack/*.go wingo-cmd/*.go wini/*.go wm/*.go \
		workspace/*.go xclient/*.go
	colcheck -c 80 *.go */*.go

cmd:
	go install github.com/BurntSushi/wingo/wingo-cmd

supported:
	scripts/generate-supported | gofmt > ewmh_supported.go

loc:
	find ./ -name '*.go' \
		-and -not -wholename './tests*' -print \
		| sort | xargs wc -l

tags:
	find ./ \( -name '*.go' \
					   -and -not -wholename './tests/*' \) \
			 -print0 \
	| xargs -0 gotags > TAGS

push:
	git push origin master
	git push github master
