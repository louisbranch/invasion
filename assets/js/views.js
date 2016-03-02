var game = new Game();

Vue.config.delimiters = ['${', '}'];

var ChatView = Vue.extend({
  init: function() {
    var view = this;
    game.events.on('chat.message', function(msg) {
      view.log.push({name: 'Luiz', content: msg});
    });
  },
  methods: {
    send: function() {
      console.log('send');
      var res = new Response();
      res.createChatMessage(this.msg);
      this.log.push({name: 'luiz', content: this.msg});
      this.msg = '';
    }
  }
});

game.events.on('join.game', function() {
  var chat = new ChatView({
    el: '#chat',
    data: {
      log: [],
      msg: ''
    }
  });
});

