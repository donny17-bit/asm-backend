function table(data) {
  console.log(data);
  const totalData = document.getElementById("total_data");
  totalData.textContent = data.total_data;

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
                                        <td>${item.Cabang}</td>
                                        <td>${item.Perwakilan}</td>
                                        <td>${item.SubPerwakilan}</td>
                                        <td>${item.Jnner}</td>
                                        <td>${item.JenisProd}</td>
                                        <td>${item.JenisPaket}</td>
                                        <td>${item.Ket}</td>
                                        <td>${item.NamaCeding}</td>
                                        <td>${item.Namaleader0}</td>
                                        <td>${item.Namaleader1}</td>
                                        <td>${item.Namaleader2}</td>
                                        <td>${item.Namaleader3}</td>
                                        <td>${item.GroupBusiness}</td>
                                        <td>${item.Business}</td>
                                        <td>${item.NoKontrak}</td>
                                        <td>${item.NoPolis}</td>
                                        <td>${item.NoCif}</td>
                                        <td>${item.ProdKe}</td>
                                        <td>${item.NamaDealer}</td>
                                        <td>${item.Tsi}</td>
                                        <td>${item.Gpw}</td>
                                        <td>${item.Disc}</td>
                                        <td>${item.Disc2}</td>
                                        <td>${item.Comm}</td>
                                        <td>${item.Oc}</td>
                                        <td>${item.Bkp}</td>
                                        <td>${item.Ngpw}</td>
                                        <td>${item.Ri}</td>
                                        <td>${item.Ricom}</td>
                                        <td>${item.Npw}</td>            
                                      </tr>`;
  });

  const table = document.getElementById("tableContainer");
  table.style.maxHeight = "1200px";
  table.style.overflowY = "visible";
}

function tableSurplus(data) {
  const totalData = document.getElementById("total_data");
  totalData.textContent = data.total_data;

  const tableBody = document.querySelector("#dataTable tbody");
  tableBody.innerHTML = "";
  data.data.forEach((item) => {
    tableBody.innerHTML += `<tr>
                                        <td>${item.Rn}</td>
                                        <td>${item.Periode}</td>
                                        <td>${item.Kanwil}</td>
                                        <td>${item.Cabang}</td>
                                        <td>${item.Perwakilan}</td>
                                        <td>${item.SubPerwakilan}</td>
                                        <td>${item.Namaleader0}</td>
                                        <td>${item.Namaleader1}</td>
                                        <td>${item.Namaleader2}</td>
                                        <td>${item.Namaleader3}</td>
                                        <td>${item.Mo}</td>
                                        <td>${item.GroupBusiness}</td>
                                        <td>${item.Business}</td>
                                        <td>${item.ClientName}</td>
                                        <td>${item.NoPolis}</td>
                                        <td>${item.NoCif}</td>
                                        <td>${item.JenisPaket}</td>
                                        <td>${item.Keterangan}</td>
                                        <td>${item.NamaCeding}</td>
                                        <td>${item.NamaDealer}</td>
                                        <td>${item.Tsi}</td>
                                        <td>${item.Gpw}</td>
                                        <td>${item.Disc}</td>
                                        <td>${item.Disc2}</td>
                                        <td>${item.Comm}</td>
                                        <td>${item.Oc}</td>
                                        <td>${item.Bkp}</td>
                                        <td>${item.Ngpw}</td>
                                        <td>${item.Ri}</td>
                                        <td>${item.Ricom}</td>
                                        <td>${item.Npw}</td>
                                        <td>${item.CadPremi}</td>
                                        <td>${item.CadPremi1}</td>
                                        <td>${item.PremiumReserve}</td>
                                        <td>${item.Npe}</td>
                                        <td>${item.AcceptedClaim}</td>			
                                        <td>${item.RejectedClaim}</td>
                                        <td>${item.OutstandingClaim}</td>
                                        <td>${item.ReversedClaim}</td>
                                        <td>${item.SurplusUw}</td>
                                      </tr>`;
  });

  const table = document.getElementById("tableContainer");
  table.style.maxHeight = "1200px";
  table.style.overflowY = "visible";
}

function tableSurplusYr(data) {
  const totalData = document.getElementById("total_data");
  totalData.textContent = data.total_data;

  const tableBody = document.querySelector("#dataTable tbody");
  tableBody.innerHTML = "";
  data.data.forEach((item) => {
    tableBody.innerHTML += `<tr>
                                        <td>${item.Rn}</td>
                                        <td>${item.Proddatetime}</td>
                                        <td>${item.Prodkey}</td>
                                        <td>${item.Kanwil}</td>
                                        <td>${item.Cabang}</td>
                                        <td>${item.Perwakilan}</td>
                                        <td>${item.SubPerwakilan}</td>
                                        <td>${item.Namaleader0}</td>
                                        <td>${item.Namaleader1}</td>
                                        <td>${item.Namaleader2}</td>
                                        <td>${item.Namaleader3}</td>
                                        <td>${item.Mo}</td>
                                        <td>${item.GroupBusiness}</td>
                                        <td>${item.Business}</td>
                                        <td>${item.ClientName}</td>
                                        <td>${item.NoPolis}</td>
                                        <td>${item.NoCif}</td>
                                        <td>${item.JenisPaket}</td>
                                        <td>${item.Keterangan}</td>
                                        <td>${item.NamaCeding}</td>
                                        <td>${item.Okupasi}</td>
                                        <td>${item.NamaDealer}</td>
                                        <td>${item.Tsi}</td>
                                        <td>${item.Gpw}</td>
                                        <td>${item.Disc}</td>
                                        <td>${item.Disc2}</td>
                                        <td>${item.Comm}</td>
                                        <td>${item.Oc}</td>
                                        <td>${item.Bkp}</td>
                                        <td>${item.Ngpw}</td>
                                        <td>${item.Ri}</td>
                                        <td>${item.Ricom}</td>
                                        <td>${item.Npw}</td>
                                        <td>${item.CadPremi}</td>
                                        <td>${item.CadPremi1}</td>
                                        <td>${item.PremiumReserve}</td>
                                        <td>${item.Npe}</td>
                                        <td>${item.AcceptedClaim}</td>			
                                        <td>${item.RejectedClaim}</td>
                                        <td>${item.OutstandingClaim}</td>
                                        <td>${item.ReversedClaim}</td>
                                        <td>${item.SurplusUw}</td>
                                      </tr>`;
  });

  const table = document.getElementById("tableContainer");
  table.style.maxHeight = "1200px";
  table.style.overflowY = "visible";
}

function tableAccept(data) {
  console.log(data);
  const totalData = document.getElementById("total_data");
  totalData.textContent = data.total_data;

  const tableBody = document.querySelector("#dataTable tbody");
  tableBody.innerHTML = "";
  data.data.forEach((item) => {
    tableBody.innerHTML += `<tr>
                                        <td>${item.Rn}</td>
                                        <td>${item.Kanwil}</td>
                                        <td>${item.Cabang}</td>
                                        <td>${item.Perwakilan}</td>
                                        <td>${item.SubPerwakilan}</td>
                                        <td>${item.Namaleader0}</td>
                                        <td>${item.Namaleader1}</td>
                                        <td>${item.Namaleader2}</td>
                                        <td>${item.Namaleader3}</td>
                                        <td>${item.GroupBusiness}</td>
                                        <td>${item.Business}</td>
                                        <td>${item.TahunPolis}</td>
                                        <td>${item.AcceptedNo}</td>
                                        <td>${item.NoKlaim}</td>
                                        <td>${item.NoPolis}</td>
                                        <td>${item.NoCif}</td>
                                        <td>${item.ClientName}</td>
                                        <td>${item.Mo}</td>
                                        <td>${item.PrepareDate}</td>
                                        <td>${item.DateOfLoss}</td>
                                        <td>${item.AcceptedDate}</td>
                                        <td>${item.BeginDate}</td>
                                        <td>${item.EndDate}</td>
                                        <td>${item.JenisPaket}</td>
                                        <td>${item.Workshop}</td>
                                        <td>${item.NamaDealer}</td>
                                        <td>${item.ColDesk}</td>
                                        <td>${item.RiskLoc}</td>
                                        <td>${item.Tsi}</td>
                                        <td>${item.AcceptedClaim}</td>
                                        <td>${item.AcceptedClaimRp}</td>
                                        <td>${item.AccKlaimGrossRp}</td>
                                        <td>${item.OwnRetention}</td>
                                        <td>${item.CoIns}</td>
                                        <td>${item.Psrspl}</td>
                                        <td>${item.Qsri}</td>
                                        <td>${item.Er1}</td>
                                        <td>${item.Surplus1}</td>
                                        <td>${item.Surplus2}</td>            
                                        <td>${item.Er2}</td>
                                        <td>${item.PsrqsRi}</td>
                                        <td>${item.PsrqsOr}</td>
                                        <td>${item.Ors}</td>
                                        <td>${item.Facultative}</td>
                                        <td>${item.Facobl}</td>
                                        <td>${item.Bppdan}</td>
                                        <td>${item.Xl}</td>
                                        <td>${item.Pss}</td>
                                        <td>${item.Prgbi}</td>
                                        <td>${item.Pfra}</td>
                                        <td>${item.Fsplnsri}</td>
                                        <td>${item.Psplnsri}</td>
                                        <td>${item.Fsplnsor}</td>
                                        <td>${item.Psplnsor}</td>
                                        <td>${item.Facobsrb}</td>
                                        <td>${item.Facobindt}</td>
                                      </tr>`;
  });

  const table = document.getElementById("tableContainer");
  table.style.maxHeight = "1200px";
  table.style.overflowY = "visible";
}
