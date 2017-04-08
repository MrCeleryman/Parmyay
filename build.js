const spawn = require('child_process').spawn;
const fs = require('fs');
const path = require("path");

process.env.GOPATH = path.resolve(__dirname, "src");
process.env.GOBIN = path.resolve(__dirname, "src", "bin");

// Path to the directory Go is installed in
process.env.GOROOT = process.env.GOROOT;

const checkDirs = ['dist', 'src/pkg', 'src/src', 'src/bin']

makeIfNotExist = x => {
	if (!fs.existsSync(x)) {
		fs.mkdirSync(x);
	}
}
checkDirs.forEach(makeIfNotExist);

function spawnProc(proc, args, config) {
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
}

spawnProc("go", ["get"], {
	cwd: path.resolve(__dirname, "src/"),
	shell: true
})
.on("close", exitCode => {
	if (exitCode != 0) { 
		process.exit(exitCode);
	}

	spawnProc("go", ["build"], {
		cwd: path.resolve(__dirname, "src")
	}).on("close", exitCode => {
		let sourcePath = path.resolve(__dirname, "src", "src.exe");
		let source = fs.createReadStream(sourcePath);
		let dest = fs.createWriteStream(path.resolve(__dirname, "dist", "api.exe"));

		source.pipe(dest);
		source.on("end", () => {
			fs.unlinkSync(sourcePath)
		}).on("error", (err) => {
			console.log(err);
		})
	});
})


