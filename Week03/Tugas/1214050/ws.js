function PostSignUp(nama, email, negara, telepon, password, terms) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    nama: nama,
    email: email,
    negara: negara,
    telepon: telepon,
    password: password,
    terms: terms
  });

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };

  fetch("https://eo4gldaglxfeivy.m.pipedream.net", requestOptions)
    .then(response => response.text())
    .then(result => GetResponse(result))
    .catch(error => console.log('error', error));
}
function PushButton() {
  // alert("dani")
  let nama = document.getElementById("nama").value;
  let email = document.getElementById("email").value;
  let negara = document.getElementById("negara").value;
  let telepon = document.getElementById("telepon").value;
  let password = document.getElementById("password").value;
  let terms = document.getElementById("terms").checked

  if (nama || email || negara || telepon || password) {
    PostSignUp(nama, email, negara, telepon, password, terms);
    document.getElementById("formsignup").style.display = 'none';
    document.getElementById("formsignup").style.display = 'block';
  } else{
    document.getElementById("status").innerHTML = 'form tidak lengkap';
  }

}
function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}