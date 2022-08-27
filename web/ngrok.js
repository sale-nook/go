const ngrok = require("ngrok")

ngrok
	.connect({
		proto: "http",
		addr: 3000,
		subdomain: "aws-appsync-go",
		region: "eu",
	})
	.then((url) => {
		console.log("frontend will be served from %s", url)
	})
