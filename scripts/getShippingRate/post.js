logToFile("post org:",origData);
logToFile("post rep:",curReply);

if (origData.vendor == 'Dominos Ritza') {
	if (origData.customer == 'Guy') {
		curReply.value = 0;
	} else {
		var dt = new Date();
		var dy = dt.getDay();
		if (dy == 3) {
			curReply.value /= 2;
		}
	}
};

logToFile("post final rep:",curReply);

ret(curReply);


