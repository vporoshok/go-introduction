sources = *.adoc examples/**/*.go

docs: docs/index.html

docs/index.html: ${sources}
	asciidoctor -r asciidoctor-diagram main.adoc -o docs/index.html

watch: docs
	fswatch -o ${sources} | xargs -n1 -I{} make docs

pdf: main.pdf

main.pdf: ${sources}
	asciidoctor-pdf -r asciidoctor-diagram main.adoc