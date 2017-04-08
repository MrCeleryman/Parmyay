const fs = require('fs');
const resolve = require("path").resolve.bind(null, __dirname, "../");
const util = require("./index");

process.env = Object.assign({}, process.env, util.env);
util.setupStructure();

util.spawnProc("go", ["get"], {
	cwd: resolve("src"),
	shell: true
})
.on("close", exitCode => {
	if (exitCode != 0) { 
		process.exit(exitCode);
	}

	util.spawnProc("go", ["build"], {
		cwd: resolve("src")
	}).on("close", exitCode => {
		let sourcePath = resolve("src", "src.exe");
		let source = fs.createReadStream(sourcePath);
		let dest = fs.createWriteStream(resolve("dist", "api.exe"));

		source.pipe(dest);
		source.on("end", () => {
			fs.unlinkSync(sourcePath)
		}).on("error", (err) => {
			console.log(err);
		})
	});
})
