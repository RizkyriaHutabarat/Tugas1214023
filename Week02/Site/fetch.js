var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

var requestOptions = {
  method: "GET",
  headers: myHeaders,
  redirect: "follow",
};

hasil = "";

fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
<<<<<<< HEAD
  .then(response => responsez.text())
  .then(result => tampilkan(result))
  .catch(error => console.log('error', error));

  function tampilkan(result){
    console.log(result);
    hasil=JSON.parse(result);

  }
=======
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.log("error", error));

function tampilkan(result) {
  console.log(result);
  hasil = JSON.parse(result);
}
>>>>>>> a06d2e7df129140481434e789585758a0d3bc887
