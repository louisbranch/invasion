function Game() {
  this.events = new EventEmitter();
}

Game.prototype.connectWS = function(url) {
  var game = this;
  var ws = new WebSocket(url);

  ws.addEventListener("open", function() {
    game.events.trigger("connection.open");
  });

  ws.addEventListener("close", function() {
    game.events.trigger("connection.close");
  });

  ws.addEventListener("message",function(event) {
    var fileReader = new FileReader();

    fileReader.onload = function() {
      var buffer = this.result;
      var data  = new Uint8Array(buffer);
      var builder = new flatbuffers.ByteBuffer(data);
      var msg = server.Message.getRootAsMessage(builder);

      var unionType = msg.responseType();

      if (msg.response(unionType)) {
        switch (unionType) {
          case server.Response.ChatMessage:
            var c = msg.response(new server.ChatMessage());
            game.events.trigger('chat.message', c.message());
            break;
          default:
            throw("Unknown respose type", unionType);
        }
      }
    }

    fileReader.readAsArrayBuffer(event.data);
  });
}
