$(document).ready(function () {
  listUsers()
});

function listUsers() {
  $.getJSON("/api/v1/users", (data) => {
    var users = ''
    console.log(data)
    for (var i = 0; i < data.user.length; i++) {
      users += '<li class="list-group-item">' + data.user[i].title + '</li>'
    }
    $('#users').html('')
    $('#users').append(users)
  })
}

$('#add_user').on('click', (e) => {
  var user = $('#user').val()
  $.post("/api/v1/users", "user=" + user, (data) => {
    console.log(data)
    $('#users').prepend('<li class="list-group-item">' + data.name + '</li>')
  })
})