function pagination(data) {
  const page = document.querySelector("#pagination");
  const firstPage = document.getElementById("first_page");
  const lastPage = document.getElementById("last_page");
  const next = document.getElementById("next");
  const previous = document.getElementById("previous");
  const currentPage = document.getElementById("current_page");
  const nextPage = document.getElementById("next_page");
  const previousPage = document.getElementById("previous_page");
  const morePrevious = document.getElementById("more_previous");
  const moreNext = document.getElementById("more_next");

  //  current page button
  currentPage.textContent = data.current_page;
  currentPage.classList.remove("disabled");
  currentPage.classList.add("active");

  //  previous button
  if (data.current_page === 1) {
    previous.disabled = true;
  } else {
    previous.disabled = false;
    previous.classList.remove("disabled");
  }

  //  next button
  if (data.current_page === data.max_page) {
    next.disabled = true;
  } else {
    next.disabled = false;
    next.classList.remove("disabled");
  }

  //  previous page (number) button
  if (data.previous_page !== 0) {
    previousPage.classList.remove("d-none");
    previousPage.textContent = data.previous_page;
  } else {
    previousPage.classList.add("d-none");
  }

  //  next page (number) button
  if (data.next_page !== data.max_page) {
    nextPage.classList.remove("d-none");
    nextPage.textContent = data.next_page;
  } else {
    nextPage.classList.add("d-none");
  }

  // more previous button & first page button
  if (data.current_page >= 3) {
    morePrevious.classList.remove("d-none");
    firstPage.classList.remove("d-none");
  } else {
    morePrevious.classList.add("d-none");
    firstPage.classList.add("d-none");
  }

  // last page button
  if (data.current_page <= data.max_page - 1) {
    lastPage.classList.remove("d-none");
    lastPage.textContent = data.max_page;
    nextPage.textContent = data.next_page;
  } else {
    lastPage.classList.add("d-none");
    nextPage.textContent = data.max_page;
  }

  // more next button
  if (data.current_page <= data.max_page - 2) {
    moreNext.classList.remove("d-none");
  } else {
    moreNext.classList.add("d-none");
  }
}
