document
  .getElementById("formFilter")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size").value; //inlcude
    const begin_date = document.getElementById("begin_date").value; //include
    const end_date = document.getElementById("end_date").value; //include
    const no_polis = document.getElementById("no_polis").value; //include
    const no_cif = document.getElementById("no_cif").value;
    const client_name = document.getElementById("client_name").value; // include
    const branch = document.getElementById("branch").value;
    const business = document.getElementById("business").value; // include
    const sumbis = document.getElementById("sumbis").value;

    // Call the API (assuming a POST request)
    fetch("/api/production-longterm", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        page: "1",
        page_size: page_size,
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
  });

document
  .getElementById("next_page")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const next_page = document.getElementById("next_page").textContent;
    const page_size = document.getElementById("page_size").value;
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
        page: next_page,
        page_size: page_size,
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
  });

document
  .getElementById("previous_page")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const previous_page = document.getElementById("previous_page").textContent;
    const page_size = document.getElementById("page_size").value;
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
        page: previous_page,
        page_size: page_size,
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
  });

document
  .getElementById("first_page")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size").value;
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
        page_size: page_size,
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
  });

document
  .getElementById("last_page")
  .addEventListener("click", function (event) {
    event.preventDefault();

    const last_page = document.getElementById("last_page").textContent;
    const page_size = document.getElementById("page_size").value;
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
        page: last_page,
        page_size: page_size,
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
  });

document.getElementById("previous").addEventListener("click", function (event) {
  event.preventDefault();

  const previous_page = document.getElementById("previous_page").textContent;
  const page_size = document.getElementById("page_size").value;
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
      page: previous_page,
      page_size: page_size,
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
});

document.getElementById("next").addEventListener("click", function (event) {
  event.preventDefault();

  const next_page = document.getElementById("next_page").textContent;
  const page_size = document.getElementById("page_size").value;
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
      page: next_page,
      page_size: page_size,
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
});

document.getElementById("export").addEventListener("click", function (event) {
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
  fetch("/api/export-production-longterm", {
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
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      return response.blob();
    })
    .then((blob) => {
      // Create a link element
      const link = document.createElement("a");

      // Create a URL for the Blob
      const url = window.URL.createObjectURL(blob);
      link.href = url;
      link.download = "Detail_Produksi_Longterm.xlsx";

      // Append the link to the body
      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    })
    .catch((error) => console.error("Error:", error));
});
