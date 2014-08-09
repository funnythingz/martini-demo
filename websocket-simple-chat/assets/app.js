var App = function() {

    var webSocket = new WebSocket('ws://localhost:3000/ws');

    webSocket.onopen = function() {

        webSocket.onmessage = function(response) {
            location.href = '/';
        }

        $('#form').on('submit', function(e) {

            e.preventDefault();

            if(! _.isEmpty($('#inputMessage').val())) {
                webSocket.send(JSON.stringify({type: 'new', data: $('#inputMessage').val()}));
            }

        });

        var deleteButton = $('.delete-entry');

        deleteButton.on('click', function(e) {

            e.stopPropagation();

            var data = $(this).data('id');
            var url = '/ws/' + data;

            var promise = $.ajax({
                "url" : url,
                "data" : data,
                "type" : "DELETE"
            });

            promise.done(function() {
                webSocket.send(JSON.stringify({type: 'delete', data: data}));
            });

        });

    }

}

$(function() {
    new App();
});
