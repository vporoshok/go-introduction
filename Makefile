sources = *.adoc examples/**/*.go

docs: docs/index.html

docs/index.html: ${sources}
	asciidoctor -r asciidoctor-diagram main.adoc -o docs/index.html

watch:
	@while true; do \
		make docs; \
		inotifywait -qe modify ${sources}; \
	done

pdf: main.pdf

main.pdf: ${sources}
	asciidoctor-pdf -r asciidoctor-diagram main.adoc