document
  .getElementById("formFilter")
  .addEventListener("submit", function (event) {
    event.preventDefault();

    const page_size = document.getElementById("page_size").value; //inlcude
    const begin_date = document.getElementById("begin_date").value; //include
    const end_date = document.getElementById("end_date").value; //include
    const no_polis = document.getElementById("no_polis").value; //include
    const no_cif = document.getElementById("no_cif").value; // include
    const client_name = document.getElementById("client_name").value; // include
    const branch = document.getElementById("branch").value; // include
    const business = document.getElementById("business").value; // include
    const sumbis = document.getElementById("sumbis").value; // include

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
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
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        table(data);

        // pagination script
        pagination(data);
      })
      .catch((error) => {
        return console.error("Error:", error);
      });
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

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
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

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
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

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
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

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
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

  const url = window.location.href;
  const path = url.split("/").pop();

  // Call the API (assuming a POST request)
  fetch(`/api/${path}`, {
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

  const url = window.location.href;
  const path = url.split("/").pop();

  // Call the API (assuming a POST request)
  fetch(`/api/${path}`, {
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

  const url = window.location.href;
  const path = url.split("/").pop();

  // Call the API (assuming a POST request)
  fetch(`/api/export-${path}`, {
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

      const contentDisposition = response.headers.get("Content-Disposition");
      let filename = "";

      if (
        contentDisposition &&
        contentDisposition.indexOf("attachment") !== -1
      ) {
        const filenameRegex = /filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/;
        const matches = filenameRegex.exec(contentDisposition);
        if (matches != null && matches[1]) {
          filename = matches[1].replace(/['"]/g, "");
        }
      }

      return response.blob().then((blob) => ({ blob, filename }));
    })
    .then(({ blob, filename }) => {
      // Create a link element
      const link = document.createElement("a");

      // Create a URL for the Blob
      const url = window.URL.createObjectURL(blob);
      link.href = url;
      link.download = filename;

      // Append the link to the body
      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    })
    .catch((error) => console.error("Error:", error));
});

document.getElementById("logout").addEventListener("click", function (event) {
  event.preventDefault();

  fetch("/api/logout", {
    method: "GET",
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      return response.json();
    })
    .then((json) => {
      if (json.status != 200) {
        return console.log(json.message);
      }

      return (window.location.href = "/login");
    })
    .catch((error) => console.error("Error:", error));
});
