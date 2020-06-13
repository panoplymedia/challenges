const express = require("express");
const app = express();
const parser = require("body-parser");
const passport = require("passport");
const session = require("express-session");
const fileUpload = require('express-fileupload');


const CONFIG = require("./config.js");

let fs = require('fs');
const csvParser = require('csv-parser');

const pool = require('./db');
const GitHubStrategy = require("passport-github2").Strategy;

passport.serializeUser(function (user, done) {
	done(null, user);
  });
  passport.deserializeUser(function (obj, done) {
	done(null, obj);
  });


passport.use(
new GitHubStrategy(
	{
		clientID: CONFIG.GITHUB_CLIENT_ID,
		clientSecret: CONFIG.GITHUB_CLIENT_SECRET,
		callbackURL: CONFIG.GITHUB_CALLBACK_URL,
	},
		function (accessToken, refreshToken, profile, done) {
			process.nextTick(function () {
				return done(null, profile);
			});
		}
	)
);

app.use(parser.urlencoded({ extended: true }));
app.use(parser.json());

app.use(function (req, res, next) {
	res.header("Access-Control-Allow-Methods", "POST,GET,PATCH,PUT,DELETE");
	res.header(
		"Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept"
	);
	res.header("Access-Control-Allow-Credentials", "true");
	res.header("Access-Control-Allow-Origin", CONFIG.EXPRESS_ALLOW_ORIGIN);
	

	next();
});

app.use(
	session({ secret: "acme app", resave: false, saveUninitialized: false })
  );
  app.use(passport.initialize());
  app.use(passport.session());



app.use(fileUpload({
    useTempFiles : true,
    tempFileDir : '/tmp/'
}));





app.post("/api/upload", (req, res) => {
	// console.log("request body:", req.body);

	if( inlineCheck(req) ) {
			if (!req.files || Object.keys(req.files).length === 0)
				res.json({ 'Error': 'No file uploaded' })
			else {

				//keep the object properties consistent with our db columns
				let sales = {
					customer_name: [],
					description: [],
					price: [],
					quantity: [],
					merchant_name: [],
					merchant_address: []
				}

				let csv = req.files.csv

				
				//give it a unique name (like a unix timestamp) in case we decide to keep these on the server
				let csvName = Math.floor(Date.now() / 1000) + '.csv';
				let filePath = CONFIG.FILE_PATH + `/${csvName}`;

				
				let result = [];
				
				csv.mv(filePath).then(_ => {

					fs.createReadStream(filePath)
					.pipe(csvParser())
					.on('data', (data) => result.push(data))
					.on('end', () => {
						// console.log(result);

						result.forEach(sale => {
							console.log(sale);
							pool.query(`INSERT INTO salesdata VALUES( DEFAULT, $1, $2, $3, $4, $5, $6)`, 
							[ sale['Customer Name'], sale['Item Description'], sale['Item Price'], 
							sale['Quantity'], sale['Merchant Name'], sale['Merchant Address'] ]);
						})
						res.json(result);
					});

				})
			}
	}
	else {
		res.statusCode(401);
	}
});

app.get('/api/deleteall', (req, res) => {
	if(inlineCheck(req)) {
		pool.query('DELETE FROM salesdata').then(_ => console.log('deleted'));
		res.json({ 'Success': 'deleted'});
	}
	else {
		res.statusCode(401);
	}
})

app.get('/api/salesdata', (req, res) => {
	if(inlineCheck(req)) {
		pool.query('SELECT * FROM salesdata')
		.then(result => {
			// console.log(result);
				if(result.rowCount > 0)
					res.json(result.rows);
				else 
					res.json({ 'rowCount' : 0})
		});
	}
	else {
		res.statusCode(401);
	}
})

app.get('/api/salesdata/revenue',  (req, res) => {
	if(inlineCheck(req)) {
		pool.query('SELECT * FROM salesdata')
		.then(result => {
			// console.log(result);
			if(result.rowCount > 0)
			{
				let total = 0;
				result.rows.forEach(sale => {
					total += parseFloat(sale.price * sale.quantity);
				})
			
				console.log(total);
				res.json({'revenue': total});
			}
			else 
				res.json({ 'revenue' : 0})
		})
	} 
	else {
		res.statusCode(401);
	}
})



app.get("/login", (req, res) => {

	res.redirect("/auth/github");
  });

  app.get(
	"/auth/github",
	passport.authenticate("github", { scope: ["read:user"] }),
	function (req, res) {}
  );
  
  app.get(
	"/auth/github/callback",
	passport.authenticate("github", {
		
	  failureRedirect: "/logout",

	}),
	(req, res) =>{
	//   console.log(req);
	  console.log('callback')
	  	// res.json(req.user);
	  res.redirect('http://localhost:3000');
	}
  );

  app.get("/logout", (req, res) => {
	req.logout();
	
  });

  app.get('/authcheck', (req, res) => {

		console.log(req.user);

		if(inlineCheck(req)) {
			req.user.auth = true;
			res.json(req.user);
		} 
		else res.json({ auth: false });

  })

  
app.set("port", process.env.PORT || 8080);

app.listen(app.get("port"), () => {
	console.log(`PORT: ${app.get("port")} `);
});



inlineCheck = (req) => {
	// console.log(req);
	return req.session.passport.user !== undefined ? true: false;
};