function PostSignUp(firstname,lastname,email,password,passwordconfirm){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "firstname": firstname,
      "lastname": lastname,
      "email": email,
      "password": password,
      "passwordconfirm": passwordconfirm
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eoik64oynys7sve.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }
  
    function PushButton(){
      firstname=document.getElementById("firstname").value;
      lastname=document.getElementById("lastname").value;
      email=document.getElementById("email").value;
      password=document.getElementById("password").value;
      passwordconfirm=document.getElementById("passwordconfirm").value;
      PostSignUp(firstname,lastname,email,password,passwordconfirm);
  }
  
  function GetResponse(result){
    document.getElementById("formsignup").innerHTML = result;
  }