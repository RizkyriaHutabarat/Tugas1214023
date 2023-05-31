function PostSignUp(namadepan, namabelakang, email, password) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    namadepan: namadepan,
    namabelakang: namabelakang,
    email: email,
    password: password,
    tanggal: tanggal,
    // "no": no
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eoxwoj3x0lej4zr.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}

function PushButton() {
  namadepan = document.getElementById("namadepan").value;
  namabelakang = document.getElementById("namabelakang").value;
  email = document.getElementById("email").value;
  password = document.getElementById("password").value;
  tanggal = document.getElementById("tanggal").value;
  PostSignUp(namadepan, namabelakang, email, password, tanggal);
}
