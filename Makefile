js_dir = assets/js

js_files = $(js_dir)/lib/flatbuffers.js\
					 $(js_dir)/lib/vue.js\
					 $(js_dir)/lib/EventEmitter.js\
					 $(js_dir)/protocol/client_generated.js\
					 $(js_dir)/protocol/server_generated.js\
					 $(js_dir)/game.js\
					 $(js_dir)/response.js\
					 $(js_dir)/views.js

encode:
	rm -rf internal/protocol/* assets/js/protocol/*
	flatc --go -o internal/protocol protocol/*.fbs
	flatc --js -o assets/js/protocol protocol/*.fbs

js: $(js_files)
	cat > assets/dist/main.js $^

watch:
	find $(js_dir) -type f -name "*.js" | entr make js
