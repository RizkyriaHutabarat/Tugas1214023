function PostResv(fName, lName, guest, date, time, kehadiran) {
  var myHeaders = new Headers();
  myHeaders.append("Login", "chapter03");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    namadepan: fName,
    namabelakang: lName,
    tamu: guest,
    tanggal: date,
    waktu: time,
    hadir: kehadiran,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eo6bnvvbam0rby.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => console.log("error", error));
}

function PushButton() {
  namadepan = document.getElementById("fName").value;
  namabelakang = document.getElementById("lName").value;
  tamu = document.getElementById("guest").value;
  tanggal = document.getElementById("date").value;
  waktu = document.getElementById("time").value;
  //radioStatHadir
  // var y = document.getElementById("r1").checked;
  // hadir = y ? "Iya" : "Tidak";
  if (namadepan || namabelakang || tamu || tanggal || waktu) {
    PostResv(namadepan, namabelakang, tamu, tanggal, waktu, kehadiran());
  } else {
    document.getElementById("status").innerHTML = "Lengkasi Semua Kolom!";
  }
}

function GetResponse(result) {
  document.getElementById("resvform").innerHTML = result;
}

function kehadiran() {
  var y = document.getElementById("r1").checked;
  return y ? "Iya" : "Tidak";
}
