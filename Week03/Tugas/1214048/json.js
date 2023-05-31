function PostSignUp(NamaLengkap, email, Birth, PhoneNumber, password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "adecandra12");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "NamaLengkap": NamaLengkap,
      "email": email,
      "Birth": Birth,
      "PhoneNumber": PhoneNumber,
      "password": password
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eouo679oz7aez7o.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
}

function GetResponse(result){
      document.getElementById("formsignup").innerHTML = result;
}
    
function PushButton(){
    NamaLengkap=document.getElementById("NamaLengkap").value;
    email=document.getElementById("email").value;
    Birth=document.getElementById("Birth").value;
    PhoneNumber=document.getElementById("PhoneNumber").value;
    password=document.getElementById("password").value;
    PostSignUp(NamaLengkap, email, Birth, PhoneNumber, password);
}