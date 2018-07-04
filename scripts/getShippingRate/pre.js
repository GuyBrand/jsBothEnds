logToFile("pre", origData);


if (origData.vendor == 'LA-Beard' && origData.customer == 'Guy') {
		origData.vendor = 'Not-Ella';
};
logToFile("pre hooked ", origData);

ret(origData);
