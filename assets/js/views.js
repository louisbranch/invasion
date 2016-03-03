/* globals Game, Vue, Response */
var game = new Game();

Vue.config.delimiters = ['${', '}'];

var ChatView = Vue.extend({
  template: '#template-chat',
  replace: false,
  data: function() {
    return { log: [], msg: '' };
  },
  init: function() {
    var view = this;
    game.events.on('chat.message', function(msg) {
      view.log.push({name: 'Luiz', content: msg});
    });
  },
  methods: {
    send: function() {
      var msg = this.msg.trim();
      if (msg.length === 0) return;
      var res = new Response();
      res.createChatMessage(msg);
      this.log.push({name: 'luiz', content: msg});
      this.msg = '';
    }
  }
});

game.events.on('join.game', function() {
  var chat = new ChatView();
  chat.$mount('#chat');
});
