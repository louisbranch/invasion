(function(window, document) {

  var Response = new window.EventEmitter();

  var Vue = window.Vue;
  var ws = new WebSocket("ws://localhost:8080/ws");

  Vue.config.delimiters = ['${', '}'];

  ws.addEventListener("open", function() {
    var buf = createAccount("luiz", "me@luizbranco.com");
    ws.send(buf);
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
            Response.trigger('chat.message', c.message());
            break;
          default:
            console.log("Unknown respose type", unionType);
        }
      }
    }

    fileReader.readAsArrayBuffer(event.data);
  });

  function createAccount(name, email) {
    var builder = new flatbuffers.Builder(0);
    var name = builder.createString(name);
    var email = builder.createString(email);
    client.CreateAccount.startCreateAccount(builder);
    client.CreateAccount.addName(builder, name);
    client.CreateAccount.addEmail(builder, email);
    var req = client.CreateAccount.endCreateAccount(builder);
    return finishMessage(builder, req, client.Request.CreateAccount);
  }

  function createGame() {
    var builder = new flatbuffers.Builder(0);
    client.CreateGame.startCreateGame(builder);
    var req = client.CreateGame.endCreateGame(builder);
    return finishMessage(builder, req, client.Request.CreateGame);
  }

  function createChatMessage(content) {
    var builder = new flatbuffers.Builder(0);
    var txt = builder.createString(content);
    client.ChatMessage.startChatMessage(builder);
    client.ChatMessage.addMessage(builder, txt);
    var req = client.ChatMessage.endChatMessage(builder);
    return finishMessage(builder, req, client.Request.ChatMessage);
  }

  function finishMessage(builder, req, type) {
    client.Message.startMessage(builder);
    client.Message.addRequestType(builder, type);
    client.Message.addRequest(builder, req);
    var msg = client.Message.endMessage(builder);
    builder.finish(msg);
    return builder.asUint8Array();
  }


  var chatLog = [];

  new Vue({
    el: '#chat',
    init: function() {
      var view = this;
      Response.on('chat.message', function(msg) {
        view.log.push({name: 'Luiz', content: msg});
      });
    },
    data: {
      log: [],
      msg: ''
    },
    methods: {
      send: function() {
        createChatMessage(this.msg);
        this.msg = '';
      }
    }
  });

}(window, document));
