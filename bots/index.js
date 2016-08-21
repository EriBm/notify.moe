global.arn = require('../lib')
global.chalk = require('chalk')
global.Promise = require('bluebird')
global.fs = Promise.promisifyAll(require('fs'))

arn.runningBackgroundJobs = true

arn.db.ready.then(Promise.coroutine(function*() {
	let files = yield fs.readdirAsync('bots')
	let filterJob = process.argv[2]

	files.forEach(file => {
		if(file === 'index.js')
			return
		
		if(filterJob && file !== filterJob + '.js')
			return

		console.log(chalk.green('[Starting bot]'), chalk.blue(file.replace('.js', '')))
		require('./' + file)
	})
}))