function Response() {
  this._builder = new flatbuffers.Builder(0);
};

Response.prototype.createAccount = function(name, email) {
  var b = this._builder;
  var name = b.createString(name);
  var email = b.createString(email);
  client.CreateAccount.startCreateAccount(b);
  client.CreateAccount.addName(b, name);
  client.CreateAccount.addEmail(b, email);
  var req = client.CreateAccount.endCreateAccount(b);
  return this._finish(req, client.Request.CreateAccount);
}

Response.prototype.createGame = function() {
  var b = this._builder;
  client.CreateGame.startCreateGame(b);
  var req = client.CreateGame.endCreateGame(b);
  return this._finish(req, client.Request.CreateGame);
};

Response.prototype.joinGame = function() {

};

Response.prototype.createChatMessage = function(content) {
  var b = this._builder;
  var txt = b.createString(content);
  client.ChatMessage.startChatMessage(b);
  client.ChatMessage.addMessage(b, txt);
  var req = client.ChatMessage.endChatMessage(b);
  return this._finish(req, client.Request.ChatMessage);
};

Response.prototype._finish = function (req, type) {
  var b = this._builder;
  client.Message.startMessage(b);
  client.Message.addRequestType(b, type);
  client.Message.addRequest(b, req);
  var msg = client.Message.endMessage(b);
  b.finish(msg);
  return b.asUint8Array();
}
