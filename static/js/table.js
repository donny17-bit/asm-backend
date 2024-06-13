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
}
