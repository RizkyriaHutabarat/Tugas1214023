 function PostSignUp(namadepan,namabelakang,alamat, phone, message){
    var myHeaders = new Headers();
    myHeaders.append("Login", "pedirasyaila");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "alamat": alamat,
      "phone": phone,
      "message": message
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
    };
    console.log(raw)
   
    fetch("https://eo6vcal8x1jx6rm.m.pipedream.net", requestOptions)
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
        alamat=document.getElementById("alamat").value;
        phone=document.getElementById("phone").value;
        message=document.getElementById("message").value;
        PostSignUp(namadepan,namabelakang, alamat, phone, message);
  }
