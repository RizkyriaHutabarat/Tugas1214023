function PostSignUp(displayname, username, password, alamat, email) {
  var myHeaders = new Headers();
  myHeaders.append('Login', 'chapter03');
  myHeaders.append('Content-Type', 'application/json');

  var raw = JSON.stringify({
    displayname: displayname,
    username: username,
    password: password,
    alamat: alamat,
    email: email,
  });

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow',
  };

  fetch('https://eoy4cghvqqmlrnr.m.pipedream.net', requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log('error', error));
}

function PushButton() {
  displayname = document.getElementById('displayname').value;
  username = document.getElementById('username').value;
  password = document.getElementById('password').value;
  alamat = document.getElementById('alamat').value;
  email = document.getElementById('email').value;
  PostSignUp(displayname, username, password, alamat, email);
}

function GetResponse(result) {
  document.getElementById('formsignup').innerHTML = result;
}
