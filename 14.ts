import * as files from 'fs';

let input = files.readFileSync('14.txt').toString().split(new RegExp('\n?mask = ')).map((a) => a.split(new RegExp('\nmem\\[')).map((a) => a.split(new RegExp('\\] = '))));
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

numbers = [];

function simplify(mask: string[], address: number) {
	for (let i = 0; i < mask.length; i++) {
		if (mask[i] === '0') {
			mask[i] = (getNumberAtPosition(address, 36 - i - 1)).toString();
		} else if (mask[i] === '1') {
			mask[i] = '1';
		}
	}
	return mask.join('');
}

function setAddresses(mask: string[], address: number, value: number) {
	for (let i = 0; i < mask.length; i++) {
		if (mask[i] === 'X') {
			mask[i] = '0';
			setAddresses(Array.from(mask), address, value);
			mask[i] = '1';
		}

	}

	addresses.push(parseInt(mask.join(''), 2));
	numbers[parseInt(mask.join(''), 2)] = value;
}

let addresses = [];

for (let i = 0; i < input.length; i++) {
	let mask = input[i][0][0];
	for (let j = 1; j < input[i].length; j++) {
		let address = Number(input[i][j][0]);
		mask = simplify(mask.split(''), address);
		setAddresses(mask.split(''), address, Number(input[i][j][1]));
	}
}

final = 0;

addresses.sort((a, b) => a - b);

for (let i = 0; i < addresses.length; i++) {
	if (addresses[i] === addresses[i - 1])
		continue;

	final += numbers[addresses[i]];
}

console.log('2nd Part: ' + final);
