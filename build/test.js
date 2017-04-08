const resolve = require("path").resolve.bind(null, __dirname, "../");
const util = require("./index");

process.env = Object.assign({}, process.env, util.env);
util.setupStructure();

util.spawnProc("go", ["test"], {
	cwd: resolve("src"),
	shell: true
})
.on("close", exitCode => {
	if (exitCode != 0) { 
		process.exit(exitCode);
	}
})
