import * as files from 'fs';

let input: number[] = files.readFileSync('01.txt').toString().split('\n').map(Number);

for (let i: number = 0; i < input.length - 1; i++) {
	for (let j: number = i + 1; j < input.length; j++) {
		if (input[i] + input[j] === 2020) {
			console.log(input[i] + ' * ' + input[j] + ' = ' + input[i] * input[j]);
		}
	}
}


for (let i: number = 0; i < input.length - 2; i++) {
	for (let j: number = i + 1; j < input.length - 1; j++) {
		for (let x: number = j + 1; x < input.length; x++) {
			if (input[i] + input[j] + input[x] === 2020) {
				console.log(input[i] + ' * ' + input[j] + ' * ' + input[x] + ' = ' + input[i] * input[j] * input[x]);
			}
		}
	}
}
