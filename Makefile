docs: docs/index.html

docs/index.html: *.adoc
	asciidoctor main.adoc -o docs/index.html
