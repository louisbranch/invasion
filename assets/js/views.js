var game = new Game();

Vue.config.delimiters = ['${', '}'];

new Vue({
  el: '#chat',
  init: function() {
    var view = this;
    game.events.on('chat.message', function(msg) {
      view.log.push({name: 'Luiz', content: msg});
    });
  },
  data: {
    log: [],
    msg: ''
  },
  methods: {
    send: function() {
      var res = new Response();
      res.createChatMessage(this.msg);
      this.log.push({name: 'luiz', content: this.msg});
      this.msg = '';
    }
  }
});
