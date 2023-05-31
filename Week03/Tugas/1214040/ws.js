function PostSignUp(username,email,namadepan,namabelakang,alamat,kota,negara,kodepos,tentangsaya){
    var myHeaders = new Headers();
    myHeaders.append("Login", "marlinalubis");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "username" : username,
      "email" : email,
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "alamat" : alamat,
      "kota" : kota,
      "negara" : negara,
      "kodpos" : kodepos,
      "tentangsaya" : tentangsaya
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("https://eok0n499950n04x.m.pipedream.net", requestOptions)
      .then(response => response.text())
      .then(result => GetResponse(result))
      .catch(error => console.log('error', error));
    }
  
  function GetResponse(result){
      document.getElementById("formaccountcard").innerHTML = result;
  }
    function PushButton(){
        username=document.getElementById("username").value;
        email=document.getElementById("email").value;
        namadepan=document.getElementById("namadepan").value;
        namabelakang=document.getElementById("namabelakang").value;
        alamat=document.getElementById("alamat").value;
        kota=document.getElementById("kota").value;
        negara=document.getElementById("negara").value;
        kodepos=document.getElementById("kodepos").value;
        tentangsaya=document.getElementById("tentangsaya").value;
        PostSignUp(username,email,namadepan,namabelakang,alamat,kota,negara,kodepos,tentangsaya);
    }
    