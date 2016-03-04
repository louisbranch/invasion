js_dir = assets/js
js_lib_dir = bower_components

js_files = $(js_lib_dir)/flatbuffers/index.js\
					 $(js_lib_dir)/vue/dist/vue.js\
					 $(js_lib_dir)/EventEmitter/index.js\
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
