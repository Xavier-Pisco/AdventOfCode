import * as files from 'fs';

let input = files.readFileSync('07.txt').toString().split(new RegExp(' bags\*.\n'))
	.map((a) => a.split(new RegExp(' bags\* contain '))
		.map((b) => b.split(new RegExp(' bags\*, '))));

input.splice(input.length - 1, 1);

function checkBag1(bags: string[][][], currentBag: string): number {
	for (let i = 0; i < bags.length; i++) {
		if (bags[i][0][0] == currentBag) {
			if (bags[i][1][0] === 'no other') {
				return 0;
			} else {
				for (let j = 0; j < bags[i][1].length; j++) {
					let bag = bags[i][1][j].substr(2);
					if (bag === 'shiny gold') {
						return 1;
					}
					if (checkBag1(bags, bag) == 1)
						return 1;
				}
			}
		}
	}
	return 0;
}

function checkBag2(bags: string[][][], currentBag: string): number {
	for (let i = 0; i < bags.length; i++) {
		if (bags[i][0][0] == currentBag) {
			if (bags[i][1][0] === 'no other') {
				return 0;
			} else {
				let total = 0;
				for (let j = 0; j < bags[i][1].length; j++) {
					let bag = bags[i][1][j].substr(2);
					let bag2 = checkBag2(bags, bag);
					total += Number(bags[i][1][j][0]);
					total += bag2 * Number(bags[i][1][j][0]);
				}
				return total;
			}
		}
	}
	return 0;
}

let result = 0;

for (let i = 0; i < input.length; i++) {
	result += checkBag1(input, input[i][0][0]);
}

console.log('1st Part: ' + result);
console.log('2nd Part: ' + checkBag2(input, 'shiny gold'));
