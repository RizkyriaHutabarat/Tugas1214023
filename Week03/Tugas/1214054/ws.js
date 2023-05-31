function PostSignUp(fullname,email,country,phone_number,password,confirmpassword){
    var myHeaders = new Headers();
    myHeaders.append("Login", "dimas123");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "fullname": fullname,
      "email": email,
      "country": country,
      "phone_number": phone_number,
      "password": password,
      "confirmpassword" : confirmpassword
    });

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("https://eomrg36udm00vdw.m.pipedream.net?Login=dimas123", requestOptions)
  .then(response => response.text())
  .then(result => GetResponse(result))
  .catch(error => console.log('error', error));
}

function PushButton(){
    fullname=document.getElementById("fullname").value;
    email=document.getElementById("email").value;
    country=document.getElementById("country").value;
    phone_number=document.getElementById("phone_number").value;
    password=document.getElementById("password").value;
    confirmpassword=document.getElementById("confirmpassword").value;
    if(!fullname,!email,!country,!phone_number,!password,!confirmpassword){
      alert("Field tidak boleh kosong")
    }else if(confirmpassword != password){
      alert("Password salah")
    }else{
      PostSignUp(fullname,email,country,phone_number,password,confirmpassword);
    }
  }
  
function GetResponse(){ 
    document.getElementById("pesan").innerHTML = "<p class='py-3 bg-green-500 text-white w-full rounded text-center'>Sign Up Success !!!</p>";
    document.getElementById("fullname").value = "";
    document.getElementById("email").value = "";
    document.getElementById("country").value = "";
    document.getElementById("phone_number").value = "";
    document.getElementById("password").value = "";
    document.getElementById("confirmpassword").value = "";
}