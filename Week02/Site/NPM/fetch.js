var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AMsnp_KW3uPWTf6gN4GDNl7XAoOShdRL2.VK05aaDbN1FeG%2BScGHtOuxENv5s2ABoZZzLpqN%2FUbZs");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

hasil = ""

fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
  .then(response => response.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result) {
      iniJson = JSON.parse(result);
      console.log(iniJson);
      Length = iniJson.length;

      for(i=0; i<Length; i++)
      {
        hasil += "<tr>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+iniJson[i].text+"</td>";
        hasil += "</tr>";
      }

      document.getElementById("inidata").innerHTML = hasil;
}