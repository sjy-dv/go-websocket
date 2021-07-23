var chat = {
  user_id: "",

  ws: null,

  init: function () {
    if (!window.WebSocket) {
      alert("No WebSocket!");
      return;
    }

    this.connect();
    this.user_id = prompt("user id?");
  },
  bind: function () {
    $("#send_message").click(() => {
      if (this.ws.readyState === this.ws.OPEN) {
        this.ws.send(
          JSON.stringify({
            type: "message",
            user: this.user_id,
            data: $("#message").val(),
          })
        );
      }
      $("#message").val("");
    });
  },

  connect: function () {
    this.ws = new WebSocket("ws://" + window.location.host + "/chat");

    this.ws.onopen = function (e) {
      console.log("onopen", arguments);
    };

    this.ws.onclose = function (e) {
      console.log("onclose", arguments);
    };

    this.ws.onmessage = function (e) {
      console.info(e.data);
      $("#room").prepend(
        "<div><span>" + JSON.parse(e.data).data + "</span></div>"
      );
    };
  },
};
