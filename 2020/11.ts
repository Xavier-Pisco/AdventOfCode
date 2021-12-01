import * as files from 'fs';

let input = files.readFileSync('11.txt').toString().split('\n');

function countAdjacent(input: string[], position: number) {
	let x = position % input[0].length;
	let y = Math.floor(position / input[0].length);


	let count = 0;
	if (x > 0 && input[y][x - 1] === '#')
		count += 1;
	if (x < input[0].length - 1 && input[y][x + 1] === '#')
		count += 1;
	if (y > 0 && x > 0 && input[y - 1][x - 1] === '#')
		count += 1;
	if (y > 0 && x < input[0].length - 1 && input[y - 1][x + 1] === '#')
		count += 1;
	if (y < input.length - 1 && x > 0 && input[y + 1][x - 1] === '#')
		count += 1;
	if (y < input.length - 1 && x < input[0].length - 1 && input[y + 1][x + 1] === '#')
		count += 1;
	if (y > 0 && input[y - 1][x] === '#')
		count += 1;
	if (y < input.length - 1 && input[y + 1][x] === '#')
		count += 1;

	return count;
}


function firstPart(input: string[]) {
	let previous: string[] = Array.from(input);
	while (true) {
		let temp: string[] = [];
		for (let i = 0; i < previous.length; i++) {
			temp.push("");
			for (let j = 0; j < previous[i].length; j++) {
				if (previous[i][j] === 'L' && countAdjacent(previous, i * previous[i].length + j) === 0)
					temp[i] += '#'
				else if (previous[i][j] === '#' && countAdjacent(previous, i * previous[i].length + j) >= 4)
					temp[i] += 'L';
				else
					temp[i] += previous[i][j];
			}
		}
		let final = true;
		for (let i = 0; i < previous.length; i++) {
			for (let j = 0; j < previous[i].length; j++) {
				if (previous[i][j] !== temp[i][j]) {
					final = false;
					break;
				}
			}
		}
		if (final) return temp;
		previous = Array.from(temp);
	}
}

let result = firstPart(input);

let count = 0;
for (let i = 0; i < result.length; i++) {
	for (let j = 0; j < result[i].length; j++) {
		if (result[i][j] === '#')
			count += 1;
	}
}
console.log('1st Part: ', count);


function checkUp(input: string[], x: number, y: number) {
	if (y === 0)
		return false;
	if (input[y - 1][x] === 'L')
		return false;
	if (input[y - 1][x] === '#')
		return true;
	return checkUp(input, x, y - 1);
}

function checkDown(input: string[], x: number, y: number) {
	if (y === input.length - 1)
		return false;
	if (input[y + 1][x] === 'L')
		return false;
	if (input[y + 1][x] === '#')
		return true;
	return checkDown(input, x, y + 1);
}

function checkLeft(input: string[], x: number, y: number) {
	if (x === 0)
		return false;
	if (input[y][x - 1] === 'L')
		return false;
	if (input[y][x - 1] === '#')
		return true;
	return checkLeft(input, x - 1, y);
}

function checkRight(input: string[], x: number, y: number) {
	if (x === input[y].length - 1)
		return false;
	if (input[y][x + 1] === 'L')
		return false;
	if (input[y][x + 1] === '#')
		return true;
	return checkRight(input, x + 1, y);
}

function checkUpLeft(input: string[], x: number, y: number) {
	if (x === 0 || y === 0)
		return false;
	if (input[y - 1][x - 1] === 'L')
		return false;
	if (input[y - 1][x - 1] === '#')
		return true;
	return checkUpLeft(input, x - 1, y - 1);
}

function checkUpRight(input: string[], x: number, y: number) {
	if (x === input[y].length - 1 || y === 0)
		return false;
	if (input[y - 1][x + 1] === 'L')
		return false;
	if (input[y - 1][x + 1] === '#')
		return true;
	return checkUpRight(input, x + 1, y - 1);
}

function checkDownLeft(input: string[], x: number, y: number) {
	if (x === 0 || y === input.length - 1)
		return false;
	if (input[y + 1][x - 1] === 'L')
		return false;
	if (input[y + 1][x - 1] === '#')
		return true;
	return checkDownLeft(input, x - 1, y + 1);
}

function checkDownRight(input: string[], x: number, y: number) {
	if (x === input[y].length || y === input.length - 1)
		return false;
	if (input[y + 1][x + 1] === 'L')
		return false;
	if (input[y + 1][x + 1] === '#')
		return true;
	return checkDownRight(input, x + 1, y + 1);
}

function countVisible(input, x, y) {
	let count = 0;
	if (checkLeft(input, x, y))
		count += 1;
	if (checkRight(input, x, y))
		count += 1;
	if (checkUp(input, x, y))
		count += 1;
	if (checkDown(input, x, y))
		count += 1;
	if (checkUpRight(input, x, y))
		count += 1;
	if (checkUpLeft(input, x, y))
		count += 1;
	if (checkDownLeft(input, x, y))
		count += 1;
	if (checkDownRight(input, x, y))
		count += 1;

	return count;
}

function secondPart(input: string[]) {
	let previous: string[] = Array.from(input);
	while (true) {
		let temp: string[] = [];
		for (let i = 0; i < previous.length; i++) {
			temp.push("");
			for (let j = 0; j < previous[i].length; j++) {
				if (previous[i][j] === 'L' && countVisible(previous, j, i) === 0)
					temp[i] += '#'
				else if (previous[i][j] === '#' && countVisible(previous, j, i) >= 5)
					temp[i] += 'L';
				else
					temp[i] += previous[i][j];
			}
		}
		let final = true;
		for (let i = 0; i < previous.length; i++) {
			for (let j = 0; j < previous[i].length; j++) {
				if (previous[i][j] !== temp[i][j]) {
					final = false;
					break;
				}
			}
		}

		if (final) return temp;
		previous = Array.from(temp);
	}
}


result = secondPart(input);

count = 0;
for (let i = 0; i < result.length; i++) {
	for (let j = 0; j < result[i].length; j++) {
		if (result[i][j] === '#')
			count += 1;
	}
}

console.log('2nd Part: ', count);
