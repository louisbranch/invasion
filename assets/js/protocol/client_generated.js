// automatically generated by the FlatBuffers compiler, do not modify

/**
 * @const
*/
var client = client || {};

/**
 * @enum
 */
client.Request = {
  NONE: 0,
  CreateAccount: 1,
  CreateGame: 2,
  JoinGame: 3,
  LeaveGame: 4,
  ChatMessage: 5
};

/**
 * @constructor
 */
client.Message = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.Message}
 */
client.Message.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.Message=} obj
 * @returns {client.Message}
 */
client.Message.getRootAsMessage = function(bb, obj) {
  return (obj || new client.Message).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.Message.prototype.token = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @returns {client.Request}
 */
client.Message.prototype.requestType = function() {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? /** @type {client.Request} */ (this.bb.readUint8(this.bb_pos + offset)) : client.Request.NONE;
};

/**
 * @param {flatbuffers.Table} obj
 * @returns {?flatbuffers.Table}
 */
client.Message.prototype.request = function(obj) {
  var offset = this.bb.__offset(this.bb_pos, 8);
  return offset ? this.bb.__union(obj, this.bb_pos + offset) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.Message.startMessage = function(builder) {
  builder.startObject(3);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} tokenOffset
 */
client.Message.addToken = function(builder, tokenOffset) {
  builder.addFieldOffset(0, tokenOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {client.Request} requestType
 */
client.Message.addRequestType = function(builder, requestType) {
  builder.addFieldInt8(1, requestType, client.Request.NONE);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} requestOffset
 */
client.Message.addRequest = function(builder, requestOffset) {
  builder.addFieldOffset(2, requestOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.Message.endMessage = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} offset
 */
client.Message.finishMessageBuffer = function(builder, offset) {
  builder.finish(offset);
};

/**
 * @constructor
 */
client.CreateAccount = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.CreateAccount}
 */
client.CreateAccount.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.CreateAccount=} obj
 * @returns {client.CreateAccount}
 */
client.CreateAccount.getRootAsCreateAccount = function(bb, obj) {
  return (obj || new client.CreateAccount).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.CreateAccount.prototype.name = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.CreateAccount.prototype.email = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.CreateAccount.startCreateAccount = function(builder) {
  builder.startObject(2);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} nameOffset
 */
client.CreateAccount.addName = function(builder, nameOffset) {
  builder.addFieldOffset(0, nameOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} emailOffset
 */
client.CreateAccount.addEmail = function(builder, emailOffset) {
  builder.addFieldOffset(1, emailOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.CreateAccount.endCreateAccount = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
client.CreateGame = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.CreateGame}
 */
client.CreateGame.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.CreateGame=} obj
 * @returns {client.CreateGame}
 */
client.CreateGame.getRootAsCreateGame = function(bb, obj) {
  return (obj || new client.CreateGame).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.CreateGame.startCreateGame = function(builder) {
  builder.startObject(0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.CreateGame.endCreateGame = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
client.JoinGame = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.JoinGame}
 */
client.JoinGame.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.JoinGame=} obj
 * @returns {client.JoinGame}
 */
client.JoinGame.getRootAsJoinGame = function(bb, obj) {
  return (obj || new client.JoinGame).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.JoinGame.prototype.gameId = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.JoinGame.startJoinGame = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} gameIdOffset
 */
client.JoinGame.addGameId = function(builder, gameIdOffset) {
  builder.addFieldOffset(0, gameIdOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.JoinGame.endJoinGame = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
client.LeaveGame = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.LeaveGame}
 */
client.LeaveGame.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.LeaveGame=} obj
 * @returns {client.LeaveGame}
 */
client.LeaveGame.getRootAsLeaveGame = function(bb, obj) {
  return (obj || new client.LeaveGame).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.LeaveGame.prototype.gameId = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.LeaveGame.startLeaveGame = function(builder) {
  builder.startObject(1);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} gameIdOffset
 */
client.LeaveGame.addGameId = function(builder, gameIdOffset) {
  builder.addFieldOffset(0, gameIdOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.LeaveGame.endLeaveGame = function(builder) {
  var offset = builder.endObject();
  return offset;
};

/**
 * @constructor
 */
client.ChatMessage = function() {
  /**
   * @type {flatbuffers.ByteBuffer}
   */
  this.bb = null;

  /**
   * @type {number}
   */
  this.bb_pos = 0;
};

/**
 * @param {number} i
 * @param {flatbuffers.ByteBuffer} bb
 * @returns {client.ChatMessage}
 */
client.ChatMessage.prototype.__init = function(i, bb) {
  this.bb_pos = i;
  this.bb = bb;
  return this;
};

/**
 * @param {flatbuffers.ByteBuffer} bb
 * @param {client.ChatMessage=} obj
 * @returns {client.ChatMessage}
 */
client.ChatMessage.getRootAsChatMessage = function(bb, obj) {
  return (obj || new client.ChatMessage).__init(bb.readInt32(bb.position()) + bb.position(), bb);
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.ChatMessage.prototype.gameId = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 4);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.ChatMessage.prototype.message = function(optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 6);
  return offset ? this.bb.__string(this.bb_pos + offset, optionalEncoding) : null;
};

/**
 * @param {number} index
 * @param {flatbuffers.Encoding=} optionalEncoding
 * @returns {string|Uint8Array}
 */
client.ChatMessage.prototype.recipients = function(index, optionalEncoding) {
  var offset = this.bb.__offset(this.bb_pos, 8);
  return offset ? this.bb.__string(this.bb.__vector(this.bb_pos + offset) + index * 4, optionalEncoding) : null;
};

/**
 * @returns {number}
 */
client.ChatMessage.prototype.recipientsLength = function() {
  var offset = this.bb.__offset(this.bb_pos, 8);
  return offset ? this.bb.__vector_len(this.bb_pos + offset) : 0;
};

/**
 * @param {flatbuffers.Builder} builder
 */
client.ChatMessage.startChatMessage = function(builder) {
  builder.startObject(3);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} gameIdOffset
 */
client.ChatMessage.addGameId = function(builder, gameIdOffset) {
  builder.addFieldOffset(0, gameIdOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} messageOffset
 */
client.ChatMessage.addMessage = function(builder, messageOffset) {
  builder.addFieldOffset(1, messageOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {flatbuffers.Offset} recipientsOffset
 */
client.ChatMessage.addRecipients = function(builder, recipientsOffset) {
  builder.addFieldOffset(2, recipientsOffset, 0);
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {Array.<flatbuffers.Offset>} data
 * @returns {flatbuffers.Offset}
 */
client.ChatMessage.createRecipientsVector = function(builder, data) {
  builder.startVector(4, data.length, 4);
  for (var i = data.length - 1; i >= 0; i--) {
    builder.addOffset(data[i]);
  }
  return builder.endVector();
};

/**
 * @param {flatbuffers.Builder} builder
 * @param {number} numElems
 */
client.ChatMessage.startRecipientsVector = function(builder, numElems) {
  builder.startVector(4, numElems, 4);
};

/**
 * @param {flatbuffers.Builder} builder
 * @returns {flatbuffers.Offset}
 */
client.ChatMessage.endChatMessage = function(builder) {
  var offset = builder.endObject();
  return offset;
};

// Exports for Node.js and RequireJS
this.client = client;