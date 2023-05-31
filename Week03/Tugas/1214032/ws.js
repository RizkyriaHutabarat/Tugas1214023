function PostSignUp(namadepan,namabelakang,no,email,password){
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");
 
  var raw = JSON.stringify({
    "namadepan": namadepan,
    "namabelakang": namabelakang,
    "no" : no,
    "email": email,
    "password": password
  });
 
  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };
 
  fetch("https://eomeku4ngk7ooy4.m.pipedream.net", requestOptions)
    .then(response => response.text())
    .then(result => GetResponse(result))
    .catch(error => console.log('error', error));
  }

  function PushButton(){
    namadepan=document.getElementById("namadepan").value;
    namabelakang=document.getElementById("namabelakang").value;
    no=document.getElementById("no").value;
    email=document.getElementById("email").value;
    password=document.getElementById("password").value;
    PostSignUp(namadepan,namabelakang,no,email,password);
}

function GetResponse(result){
  document.getElementById("formsignup").innerHTML = result;
}