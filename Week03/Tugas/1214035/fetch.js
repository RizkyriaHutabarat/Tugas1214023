function PostSignUp(firstname, lastname, email, phone, message){
    var myHeaders = new Headers();
    myHeaders.append("Login", "arykaanisap");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "firstname": firstname,
      "lastname": lastname,
      "email": email,
      "phone": phone,
      "message": message
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eosbt5wphkdkcuj.m.pipedream.net?Login=arykanisap", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }
   
    function GetResponse(result){
       document.getElementById("formsignup").innerHTML = result;
   }
   
    function PushButton(){
       firstname=document.getElementById("firstname").value;
       lastname=document.getElementById("lastname").value;
       email=document.getElementById("email").value;
       phone=document.getElementById("phone").value; //innerHTML="Enter Number";
       message=document.getElementById("message").value;
       PostSignUp(firstname,lastname,email,phone, message)
    }