function PostLogin(firstname, lastname, email, country, password) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter3");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    firstname: firstname,
    lastname: lastname,
    email: email,
    country: country,
    password: password,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eow3vawotvd4hkl.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));
}

function PushButton() {
  firstname = document.getElementById("firstname").value;
  lastname = document.getElementById("lastname").value;
  email = document.getElementById("email").value;
  country = document.getElementById("country").value;
  password = document.getElementById("password").value;
  PostLogin(firstname, lastname, email, country, password);
}

function GetResponse(result) {
  document.getElementById("formlogin").innerHTML = result;
}
