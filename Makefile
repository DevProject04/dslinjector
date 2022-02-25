goc = go
source = main.go
target = dslinjector

$(target): $(source)
	$(goc) build -o $(target) $(source)

rebuild: $(source)
	rm $(target)
	$(goc) build -o $(target) $(source)

install: $(target)
	$(goc) build -o /usr/local/bin/$(target) $(source)

uninstall: $(target)
	rm /usr/local/bin/$(target)

run: $(source)
	$(goc) run $(source)

clean:
	rm $(target)