import * as files from 'fs';

let input = files.readFileSync('08.txt').toString().split('\n').map((a) => a.split(' '));

let fail = false;

function calculateAcc(input: string[][]) {
	let accumulator = 0;
	let visited: boolean[] = [];

	for (let i = 0; i < input.length; i++) {
		visited[i] = false;
	}

	for (let i = 0; i < input.length; i++) {
		if (visited[i] === true) {
			fail = true;
			break;
		}
		visited[i] = true;
		if (input[i][0] === "nop") {
			continue;
		} else if (input[i][0] === "acc") {
			let value = input[i][1].substr(1, input[i][1].length);
			if (input[i][1][0] === '+') {
				accumulator += Number(value);
			} else {
				accumulator -= Number(value);
			}
		} else if (input[i][0] === "jmp") {
			let value = input[i][1].substr(1, input[i][1].length);
			if (input[i][1][0] === '+') {
				i += Number(value) - 1;
			} else {
				i -= Number(value) + 1;
			}
		}
	}
	return accumulator;
}

console.log('1st Part: ' + calculateAcc(input));

for (let i = 0; i < input.length; i++) {
	fail = false;
	if (input[i][0] === "nop") {
		input[i][0] = "jmp";
		let accumulator = calculateAcc(input);
		if (!fail) {
			console.log('2nd Part: ' + accumulator);
			break;
		}
		input[i][0] = 'nop';
	} else if (input[i][0] === "jmp") {
		input[i][0] = "nop";
		let accumulator = calculateAcc(input);
		if (!fail) {
			console.log('2nd Part: ' + accumulator);
			break;
		}
		input[i][0] = 'jmp';
	}
}
