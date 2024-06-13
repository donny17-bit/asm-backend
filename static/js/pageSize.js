document
  .getElementById("page_size_10")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size");

    page_size.value = "10";
    page_size.textContent = "10";

    callApi(page_size);
  });

document
  .getElementById("page_size_50")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size");

    page_size.value = "50";
    page_size.textContent = "50";

    callApi(page_size);
  });

document
  .getElementById("page_size_100")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size");

    page_size.value = "100";
    page_size.textContent = "100";

    callApi(page_size);
  });

document
  .getElementById("page_size_500")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size");

    page_size.value = "500";
    page_size.textContent = "500";

    callApi(page_size);
  });

document
  .getElementById("page_size_1000")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size");

    page_size.value = "1000";
    page_size.textContent = "1000";

    callApi(page_size);
  });

function callApi(page_size) {
  const page_size_value = page_size.value;
  const begin_date = document.getElementById("begin_date").value;
  const end_date = document.getElementById("end_date").value;
  const no_polis = document.getElementById("no_polis").value;
  const no_cif = document.getElementById("no_cif").value;
  const client_name = document.getElementById("client_name").value;
  const branch = document.getElementById("branch").value;
  const business = document.getElementById("business").value;
  const sumbis = document.getElementById("sumbis").value;

  // Call the API (assuming a POST request)
  fetch("/api/production-longterm", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      page: "1",
      page_size: page_size_value,
      begin_date: begin_date,
      end_date: end_date,
      no_polis: no_polis,
      no_cif: no_cif,
      client_name: client_name,
      branch: branch,
      business: business,
      sumbis: sumbis,
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      // table
      table(data);

      // pagination script
      pagination(data);
    })
    .catch((error) => console.error("Error:", error));
}
