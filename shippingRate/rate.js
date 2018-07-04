$('#Entities').on('change', function() {
  var vendorName = this.firstElementChild.options[this.firstElementChild.selectedIndex].text;	//or value
  var csName = $("#CustomerName").val()

  var jsonReq = {
           vendor : vendorName,
           customer: csName
       };
  
	 $.ajax({
		type: "POST",
		url: "/getShippingRate",
        contentType: 'application/json',
		data: JSON.stringify(jsonReq),
		success: function (data) {
			if (data) {
				$("#ShippingFee").append(vendorName + " Shipping rate for " + (csName?csName:"All") + " = <strong>" + data.value + "</strong>$<BR>");
			} else {
				$("#ShippingFee").html("Shipping rate = <strong>Not avaiable</strong> $");
			}
		}

	 });
});
