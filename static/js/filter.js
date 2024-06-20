document.getElementById("clear").addEventListener("click", function (event) {
  event.preventDefault();

  const begin_date = document.getElementById("begin_date");
  const end_date = document.getElementById("end_date");
  const no_polis = document.getElementById("no_polis");
  const no_cif = document.getElementById("no_cif");
  const client_name = document.getElementById("client_name");
  const branch = document.getElementById("branch");
  const business = document.getElementById("business");
  const sumbis = document.getElementById("sumbis");

  begin_date.value = "";
  end_date.value = "";
  no_polis.value = "";
  no_cif.value = "";
  client_name.value = "";
  branch.value = "";
  business.value = "";
  sumbis.value = "";
});

const mthname = document.getElementById("mthname");
if (mthname) {
}
