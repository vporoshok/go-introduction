docs: docs/index.html

docs/index.html: *.adoc
	asciidoctor main.adoc -o docs/index.html

watch:
	@while true; do \
		make docs; \
		inotifywait -qe modify *.adoc; \
	done