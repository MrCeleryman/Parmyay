const spawn = require('child_process').spawn;
const fs = require('fs');
const resolve = require("path").resolve.bind(null, __dirname, "../");

module.exports = {
	spawnProc: function (proc, args, config) {
		let newProc = spawn(proc, args, config)
			.on("error", err => {
				console.log(err);
				process.exit(1);
			});

		newProc.stderr.on("data", d => {
			console.log(d.toString("utf8"));
		});
		newProc.stdout.on("data", d => {
			console.log(d.toString("utf8"));
		});

		return newProc;
	},
	setupStructure: function () {
		const checkDirs = [
			resolve('dist'), 
			resolve("api", "pkg"), 
			resolve("api", "src"), 
			resolve("api", "bin")
		];

		makeIfNotExist = x => {
			if (!fs.existsSync(x)) {
				fs.mkdirSync(x);
			}
		}
		checkDirs.forEach(makeIfNotExist);
	},
	env: {
		GOPATH: resolve("api"),
		GOBIN: resolve("api", "bin"),
	}
}
