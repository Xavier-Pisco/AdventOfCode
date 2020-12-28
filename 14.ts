import * as files from 'fs';

let input = files.readFileSync('14.txt').toString().split('mask = ').map((a) => a.split(new RegExp('\nmem\\[')).map((a) => a.split(new RegExp('\\] = '))));
input.splice(0, 1);


function getNumberAtPosition(value: number, position: number) {
	return Math.floor(value / (2 ** (position))) % 2;
}

function getValueWithMask(mask: string, value: number) {
	let result = 0;
	for (let i = 0; i < mask.length; i++) {
		if (mask[i] === '1') {
			result += 2 ** (36 - i - 1);
		} else if (mask[i] === 'X') {
			result += (2 ** (36 - i - 1)) * getNumberAtPosition(value, 36 - i - 1);
		}
	}
	return result;
}

let final = 0;
let numbers: number[] = [];

for (let i = 0; i < input.length; i++) {
	const mask = input[i][0][0];
	for (let j = 1; j < input[i].length; j++) {
		const number = Number(input[i][j][1]);
		const decoded = getValueWithMask(mask, number);
		numbers[Number(input[i][j][0])] = decoded;
	}
}

for (let number of numbers) {
	if (number !== undefined)
		final += number;
}

console.log('1st Part: ' + final);

/*
let addresses = [];

function getAddresses(mask: string[], value: number) {
	let finalMask: string[] = [];
	for (let i = 0; i < mask.length; i++) {
		if (mask[i] === '0') {
			finalMask[i] = getNumberAtPosition(value, 36 - i - 1).toString();
		} else if (mask[i] === '1') {
			finalMask[i] = '1';
		} else if (mask[i] === 'X') {
			finalMask[i] = '0';
		}
	}
	return finalMask.join('');
}


for (let i = 0; i < input.length; i++) {
	const mask = input[i][0][0];
	for (let j = 1; j < input[i].length; j++) {
		const address = Number(input[i][j][0]);
		const decoded = getAddresses(mask.split(''), address);
	}
}

for (let i = 0; i < addresses.length; i++){
	addresses[i] = addresses[i].join('');
}
console.log(addresses);
*/