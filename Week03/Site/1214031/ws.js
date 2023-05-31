function PostSignUp(fname, lname, email, password) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    namadepan: fname,
    namabelakang: lname,
    email: email,
    password: password,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eo6bnvvbam0rby.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));
}

function PushButton() {
  namadepan = document.getElementById("fname").value;
  namabelakang = document.getElementById("lname").value;
  email = document.getElementById("email").value;
  password = document.getElementById("password").value;
  PostSignUp(namadepan, namabelakang, email, password);
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}
