function PostSignUp(firstName, lastName, email, password, c_password) {
  var myHeaders = new Headers();
  myHeaders.append("create", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    firstName: firstName,
    lastName: lastName,
    email: email,
    password: password,
    c_password: c_password,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eopnksu1qas87k8.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.log("error", error));
}

function PushButton() {
  firstName = document.getElementById("firstName").value;
  lastName = document.getElementById("lastName").value;
  email = document.getElementById("email").value;
  password = document.getElementById("password").value;
  c_password = document.getElementById("c_password").value;
  PostSignUp(firstName, lastName, email, password, c_password);
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}
