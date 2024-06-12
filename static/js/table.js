document
  .getElementById("formFilter")
  .addEventListener("submit", function (event) {
    event.preventDefault();

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
        const tableBody = document.querySelector("#dataTable tbody");
        tableBody.innerHTML = "";
        data.data.forEach((item) => {
          tableBody.innerHTML += `<tr>
                                        <td>${item.Rn}</td>
                                        <td>${item.ProdDate}</td>
                                        <td>${item.BeginDate}</td>
                                        <td>${item.EndDate}</td>
                                        <td>${item.Mo}</td>
                                        <td>${item.ClientName}</td>
                                        <td>${item.Kanwil}</td>
                                    </tr>`;
        });
      })
      .catch((error) => console.error("Error:", error));
  });
