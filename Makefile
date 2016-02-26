encode:
	rm -rf internal/protocol/* assets/js/protocol/*
	flatc --go -o internal/protocol protocol/*.fbs
	flatc --js -o assets/js/protocol protocol/*.fbs
