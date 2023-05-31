const { get } = require("http");

function PostSignUp(namadepan,namabelakang,email,notelepon, password, confirmpassword){
  var myHeaders = new Headers();

var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({
  "namadepan": namadepan,
  "namabelakang": namabelakang,
  "email": email,
  "notelepon": notelepon,
  "password": password,
  "confirmpassword": confirmpassword
});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("https://eoc7j7toh3j5ef.m.pipedream.net", requestOptions)
  .then(response => response.text())
  .then(result => GetResponse(result))
  .catch(error => console.log('error', error));
}

function GetResponse(result){
  document.getElementById("formsignup").innerHTML = result;
}

function PushButton(){
  namadepan=document.getElementById("namadepan").value;
  namabelakang=document.getElementById("namabelakang").value;
  email=document.getElementById("email").value;
  notelepon=document.getElementById("notelepon").value;
  password=document.getElementById("password").value;


  confirmpassword=document.getElementById("confirmpassword").value;

  if(confirmpassword != password) {
    alert("Konfirmasi Password Salah");
  } else {
    PostSignUp(namadepan,namabelakang,email,notelepon, password, confirmpassword);
  }

  
}