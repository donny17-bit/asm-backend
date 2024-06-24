const formFilter = document.getElementById("formFilter");
const nextPage = document.getElementById("next_page");
const previousPage = document.getElementById("previous_page");
const firstPage = document.getElementById("first_page");
const lastPage = document.getElementById("last_page");
const previous = document.getElementById("previous");
const next = document.getElementById("next");
const exportBtn = document.getElementById("export");
const logout = document.getElementById("logout");
let dataReq = {};

const last_page = document.getElementById("last_page");
const previous_page = document.getElementById("previous_page");
const next_page = document.getElementById("next_page");
const page_size = document.getElementById("page_size");
const periode = document.getElementById("periode");
const begin_date = document.getElementById("begin_date");
const end_date = document.getElementById("end_date");
const no_polis = document.getElementById("no_polis");
const no_cif = document.getElementById("no_cif");
const client_name = document.getElementById("client_name");
const branch = document.getElementById("branch");
const business = document.getElementById("business");
const sumbis = document.getElementById("sumbis");

if (formFilter) {
  formFilter.addEventListener("submit", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: "1",
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: "1",
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => {
        return console.error("Error: ", error);
      });
  });
}

if (nextPage) {
  nextPage.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: next_page.textContent,
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: next_page.textContent,
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (previousPage) {
  previousPage.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: previous_page.textContent,
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: previous_page.textContent,
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        // table
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (firstPage) {
  firstPage.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: "1",
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: "1",
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        // table
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (lastPage) {
  lastPage.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: last_page.textContent,
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: last_page.textContent,
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    console.log("data req : ", dataReq);

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        // table
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (previous) {
  previous.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: previous_page.textContent,
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: previous_page.textContent,
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        // table
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (next) {
  next.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        page: next_page.textContent,
        page_size: page_size.value,
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        page: next_page.textContent,
        page_size: page_size.value,
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
    })
      .then((response) => response.json())
      .then((data) => {
        // table
        if (data.status === 400) {
          alert(data.message);
          return;
        }

        // table
        if (path.includes("surplus")) {
          if (path.includes("yearly")) {
            tableSurplusYr(data);
          } else {
            tableSurplus(data);
          }
        } else {
          table(data);
        }

        // pagination script
        pagination(data);
      })
      .catch((error) => console.error("Error:", error));
  });
}

if (exportBtn) {
  exportBtn.addEventListener("click", function (event) {
    event.preventDefault();

    if (periode) {
      dataReq = {
        periode: periode.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    } else {
      dataReq = {
        begin_date: begin_date.value,
        end_date: end_date.value,
        no_polis: no_polis.value,
        no_cif: no_cif.value,
        client_name: client_name.value,
        branch: branch.value,
        business: business.value,
        sumbis: sumbis.value,
      };
    }

    const url = window.location.href;
    const path = url.split("/").pop();

    // Call the API (assuming a POST request)
    fetch(`/api/export-${path}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(dataReq),
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
}

if (logout) {
  logout.addEventListener("click", function (event) {
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
}
