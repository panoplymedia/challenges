const express = require("express");
const app = express();
const parser = require("body-parser");
const fetch = require("node-fetch");
const shell = require("shelljs");
const fileUpload = require('express-fileupload');
const FormData = require('form-data');

const CONFIG = require("./config.js");

let fs = require('fs'),
// http = require('http'),
// https = require('https');


// shell.exec('ffmpeg -formats');

// let options = {
// 	key: fs.readFileSync('/etc/letsencrypt/live/danieledminster.com/privkey.pem'),
// 	cert: fs.readFileSync('/etc/letsencrypt/live/danieledminster.com/cert.pem')
// };

// let server = https.createServer(options, app).listen(8080, () => 'listening on 8080');



// let redis = require("redis");
// let session = require("express-session");
// let redisStore = require("connect-redis")(session);
// let client = redis.createClient();


app.use(parser.urlencoded({ extended: true }));
app.use(parser.json());

app.use(function (req, res, next) {
	res.header("Access-Control-Allow-Methods", "POST,GET,PATCH,PUT,DELETE");
	res.header(
		"Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept"
	);
	// res.header("Access-Control-Allow-Credentials", "true");
	res.header("Access-Control-Allow-Origin", CONFIG.EXPRESS_ALLOW_ORIGIN);
	

	next();
});

app.use(
	// session({
	// 	secret: "gameofbandsdev",
	// 	// cookieName: 'gobsession',
	// 	// create new redis store.
	// 	name: "__gameofbandsdev",
	// 	store: new redisStore({
	// 		host: "localhost",
	// 		port: 6379,
	// 		client: client,
	// 		ttl: 260,
	// 	}),
	// 	saveUninitialized: false,
	// 	resave: false,
	// 	cookie: {
	// 		secure: false,
	// 		maxAge: 31536000000,
	// 		httpOnly: false,
	// 	},
	// })
);

app.use(fileUpload({
    useTempFiles : true,
    tempFileDir : '/tmp/'
}));


function setUserSession(sessionVars, req, expressResponse) {
	let baseURL = "https://oauth.reddit.com/api/v1/me";
	// let rm;
	const f = fetch(baseURL, {
		headers: {
			Authorization: `Bearer ${sessionVars.access_token}`,
			scope: "identity",
		},
	})
		.then((res) => res.json())
		.then((res) => {
			// req.session.uid = res.name;
			req.session.key = res.name;
			req.session["tokens"] = sessionVars;
			console.log("session key:", req.session.key);
			console.log(res.name, res.link_karma, res.comment_karma);

			// globalsess = req.session;

			// expressResponse.json(req.session);
			req.session.save();

			expressResponse.redirect(CONFIG.ORIGIN_HOME);
			// response.redirect('https://google.com');
		});
	// return rm;
}

app.get("/api/all", (req, res) => {
	// console.log("get session id: ", req.session.id);
	// res.json(req.session);
});






app.post("/api/upload", (req, res) => {
	console.log("request body:", req.body);


	if (!req.files || Object.keys(req.files).length === 0)
		res.json({ 'Error': 'No file uploaded' })
	else {

		//keep the object properties consistent with our db columns
		let sale = {
			customer_name: '',
			description: '',
			price: '',
			quantity: '',
			merchant_name: '',
			merchant_address: ''
		}


		// let userSong = req.files.song;

		let csv = req.files.csv

		
		//give it a unique name (like a unix timestamp) in case we decide to keep these on the server
		let csvName = Math.floor(Date.now() / 1000) + '.csv';
		let filePath = CONFIG.FILE_PATH + `/${csvName}`;
		
		csv.mv(filePath);

		let fh = readFileSync(filePath);



	}
});

app.delete("/delete/:id", (req, res) => {
	console.log(req);
	Song.findById(req.params.id, (err, songUpdate) => {
		if (err) {
			res.status(500).send(err);
		} else {
			songUpdate.remove((err) => {
				if (err) res.status(500).send(err);
				else {
					let success = {
						response: "Removed from Database",
					};
					res.json(success);
				}
			});
		}
	});
});

// app.listen(4000, () => console.log("listening on localhost:4000/"));


app.set("port", process.env.PORT || 8080);


app.listen(app.get("port"), () => {
	console.log(`✅ PORT: ${app.get("port")} 🌟`);
});

