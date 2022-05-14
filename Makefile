goc    = go
source = main.go
target = dslinjector
dir    = /usr/local/bin

$(target): $(source)
	$(goc) build -o $(target) $(source)

install: $(target)
	$(goc) build -o $(dir)/$(target) $(source)

uninstall: $(target)
	rm /usr/local/bin/$(target)

run: $(source)
	$(goc) run $(source)

clean:
	rm $(target)