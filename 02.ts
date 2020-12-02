import * as files from 'fs';

let input: string[] = files.readFileSync('02.txt').toString().split('\n');
let total: number = 0;

for (let i = 0; i < input.length; i++) {
	let parts: string[] = input[i].split(' ');
	let values: number[] = parts[0].split('-').map(Number);
	let char: string = parts[1][0];
	let charNumber: number = 0;
	for (let j = 0; j < parts[2].length; j++) {
		if (parts[2][j] === char) {
			charNumber += 1;
		}
	}
	if (charNumber <= values[1] && charNumber >= values[0])
		total += 1;
}
console.log('1st Part: ' + total);


total = 0;

for (let i = 0; i < input.length; i++) {
	let parts: string[] = input[i].split(' ');
	let values: number[] = parts[0].split('-').map(Number).map((value) => value - 1);
	let char: string = parts[1][0];

	if (parts[2][values[0]] !== parts[2][values[1]] && (parts[2][values[0]] === char || parts[2][values[1]] === char))
		total += 1;
}
console.log('2nd Part: ' + total);

