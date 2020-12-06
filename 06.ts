import { group } from 'console';
import * as files from 'fs';

let input: string[] = files.readFileSync('06.txt').toString().split('\n\n');

let differentAnswers = 0;
for (let i = 0; i < input.length; i++) {
	let temp = input[i].split('\n').join('');

	let groupAnswer = temp.split('');

	let answers: string[] = [];
	for (let j = 0; j < groupAnswer.length; j++) {
		if (!answers.includes(groupAnswer[j])) {
			answers.push(groupAnswer[j]);
		}
	}
	differentAnswers += answers.length;
}


console.log('1st Part: ' + differentAnswers);

differentAnswers = 0;
for (let i = 0; i < input.length; i++) {
	let groupAnswers = input[i].split('\n');

	for (let j = 0; j < groupAnswers[0].length; j++) {
		let yes = 0;
		for (let x = 0; x < groupAnswers.length; x++) {
			if (groupAnswers[x].includes(groupAnswers[0][j])) {
				yes += 1;
			}
		}

		if (yes === groupAnswers.length)
			differentAnswers += 1;

	}
}

console.log('2nd Part: ' + differentAnswers);
