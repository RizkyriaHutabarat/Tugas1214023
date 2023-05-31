function PostSignUp(namadepan, namabelakang, no_hp, email, password, confirmpass) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "fatwaff");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    namadepan: namadepan,
    namabelakang: namabelakang,
    no_hp: no_hp,
    email: email,
    password: password,
    confirmpass: confirmpass,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eobivrw4p8r7ag5.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));
}

function PushButton() {
  namadepan = document.getElementById("namadepan").value;
  namabelakang = document.getElementById("namabelakang").value;
  no_hp = document.getElementById("no_hp").value;
  email = document.getElementById("email").value;
  password = document.getElementById("password").value;
  confirmpass = document.getElementById("confirmpass").value;
  if (!namadepan || !namabelakang || !no_hp || !email || !password) {
    document.getElementById("validpass").innerHTML = "Mohon untuk melengkapi data";
  } else {
    if (confirmpass != password) {
      document.getElementById("validpass").innerHTML = "Konfirmasi password salah!";
    } else {
      PostSignUp(namadepan, namabelakang, no_hp, email, password, confirmpass);
    }
  }
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}
