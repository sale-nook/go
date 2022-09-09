const ngrok = require("ngrok")

ngrok
	.connect({
		proto: "http",
		addr: 3000,
		subdomain: "cdk-appsync-go",
		region: "eu",
		configPath: "/home/dave/.config/ngrok/ngrok.yml",
	})
	.then((url) => {
		console.log("frontend will be served from %s", url)
	})
