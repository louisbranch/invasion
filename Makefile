encode:
	flatc --go -o internal/protocol protocol/*.fbs
	flatc --js -o assets/js/protocol protocol/*.fbs
