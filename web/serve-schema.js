const fs = require("fs")
const	http = require("http")

if (process.env.NODE_ENV === "production") {
	console.log("Production enabled, not serving schema.")
	return
} else {
	console.log("Production disabled, serving schema on http://localhost:8000/schema.graphql")
}

http.createServer(function (req, res) {
	fs.readFile("./src/graphql/schema.graphql", function (err, data) {
		if (err) {
			res.writeHead(404)
			res.end(JSON.stringify(err))
			return
		}
		res.writeHead(200)
		res.end(data)
	})
}).listen(8000)
