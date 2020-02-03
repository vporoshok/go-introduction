sources = *.adoc examples/**/*.go

docs: docs/index.html

docs/index.html: ${sources}
	asciidoctor main.adoc -o docs/index.html

watch:
	@while true; do \
		make docs; \
		inotifywait -qe modify ${sources}; \
	done